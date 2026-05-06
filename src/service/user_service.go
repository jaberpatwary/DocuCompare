package service

import (
	"app/src/config"
	models "app/src/model"
	"app/src/utils"
	"app/src/validation"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(c *fiber.Ctx) (*models.User, error)
	GetAll(c *fiber.Ctx, params *validation.QueryUser) ([]models.User, error)
	GetByUserId(c *fiber.Ctx, id string) (*models.User, error)
	//GetByPhoneNumber(c *fiber.Ctx)
	UpdateUser(c *fiber.Ctx, req *validation.UpdateUser2, id string) (*models.User, error)
	DeleteUser(c *fiber.Ctx, id string) error
	Login(email, password string) (string, error)
}

// Define methods for expert service

type userService struct {
	DB *gorm.DB
}

// DB servie init
func NewUserService(db *gorm.DB) UserService {
	return &userService{DB: db}
}

// Create
func (s *userService) CreateUser(c *fiber.Ctx) (*models.User, error) {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return nil, err
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Get All
func (s *userService) GetAll(c *fiber.Ctx, params *validation.QueryUser) ([]models.User, error) {

	var users []models.User

	offset := (params.Page - 1) * params.Limit

	query := s.DB.WithContext(c.Context()).Order("created_at asc")

	if search := params.Search; search != "" {
		query = query.Where("name LIKE? or phone LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	result := query.Find(&users).Offset(offset)

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	result = query.Limit(params.Limit).Offset(offset).Find(&users)
	if result.Error != nil {

		return nil, result.Error
	}

	return users, result.Error

}

// GetbyId

func (s *userService) GetByUserId(c *fiber.Ctx, id string) (*models.User, error) {
	user := new(models.User)

	result := s.DB.WithContext(c.Context()).First(&user, "id = ?", id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Update User
func (s *userService) UpdateUser(c *fiber.Ctx, req *validation.UpdateUser2, id string) (*models.User, error) {

	if req.Name == "" && req.Email == "" && req.Phone == "" &&
		req.AvatarURL == "" && req.Status == "" && req.Password == "" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid Request: No fields to update")
	}

	updateBody := make(map[string]interface{})

	if req.Name != "" {
		updateBody["name"] = req.Name
	}
	if req.Email != "" {
		updateBody["email"] = req.Email
	}
	if req.Phone != "" {
		updateBody["phone"] = req.Phone
	}
	if req.AvatarURL != "" {
		updateBody["avatar_url"] = req.AvatarURL
	}
	if req.Status != "" {
		updateBody["status"] = req.Status
	}
	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
		}
		updateBody["password_hash"] = hashedPassword
	}

	result := s.DB.WithContext(c.Context()).Model(&models.User{}).Where("id = ?", id).Updates(updateBody)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate") {
			return nil, fiber.NewError(fiber.StatusConflict, "Email or Phone already exists")
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		user, err := s.GetByUserId(c, id)
		if err != nil || user == nil {
			return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
		}

		return user, nil
	}

	user, err := s.GetByUserId(c, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Delete
func (s *userService) DeleteUser(c *fiber.Ctx, id string) error {
	user := new(models.User)

	result := s.DB.WithContext(c.Context()).Delete(user, "id = ?", id)

	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	if result.Error != nil {

		//s.Log.Errorf("Failed to delete user: %+v", result.Error)
	}

	return result.Error

}

// Login implementation
func (s *userService) Login(email, password string) (string, error) {
	// Special case for development admin
	if email == "admin" && password == "admin" {
		return utils.GenerateToken("1", "access", config.JWTAccessExp, config.JWTSecret)
	}

	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	return utils.GenerateToken(fmt.Sprintf("%d", user.ID), "access", config.JWTAccessExp, config.JWTSecret)
}
