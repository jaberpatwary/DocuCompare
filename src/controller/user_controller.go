package controller

import (
	"app/src/response"
	"app/src/service"
	"app/src/validation"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	_UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{

		_UserService: userService,
	}
}

// Create
func (u *UserController) CreateUser(c *fiber.Ctx) error {

	user, err := u._UserService.CreateUser(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

// Get All
func (u *UserController) GetAll(c *fiber.Ctx) error {
	query := &validation.QueryUser{

		Page:   c.QueryInt("Page", 1),
		Limit:  c.QueryInt("Limit", 20),
		Search: c.Query("Search", ""),
	}

	users, err := u._UserService.GetAll(c, query)

	if err != nil {

		return err
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

// Get By Id
func (u *UserController) GetByUserId(c *fiber.Ctx) error {

	UserId := c.Params("userId")

	user, err := u._UserService.GetByUserId(c, UserId)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)

}

func (u *UserController) UpdateUser(c *fiber.Ctx) error {
	req := new(validation.UpdateUser2)
	userID := c.Params("userId")

	fmt.Printf("Raw body: %s\n", string(c.Body()))

	if err := c.BodyParser(req); err != nil {
		fmt.Printf("BodyParser error: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	fmt.Printf("Parsed request: %+v\n", req)

	user, err := u._UserService.UpdateUser(c, req, userID)
	if err != nil {

		if fiberErr, ok := err.(*fiber.Error); ok {
			return c.Status(fiberErr.Code).JSON(fiber.Map{
				"error": fiberErr.Message,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// delete
func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("userId")

	if err := u._UserService.DeleteUser(c, userID); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(response.Common{
			Code:    fiber.StatusOK,
			Status:  "success",
			Message: "Delete user successfully",
		})

}

func (u *UserController) Login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	token, err := u._UserService.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HTTPOnly: false, // Set to false so JS can still use it for Bearer if needed
		SameSite: "Lax",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
