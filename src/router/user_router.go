package router

import (
	"app/src/controller"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(v1 fiber.Router, u service.UserService) {
	// Initialize the UserController with the UserService
	userController := controller.NewUserController(u)
	// Define user-related routes
	v1.Post("/auth/login", userController.Login)
	userGroup := v1.Group("/users")
	userGroup.Post("/", userController.CreateUser)
	userGroup.Get("/", userController.GetAll)
	userGroup.Get("/:userId", userController.GetByUserId)
	userGroup.Put("/:userId", userController.UpdateUser)
	userGroup.Delete("/:userId", userController.DeleteUser)
}
