package router

import (
	"app/src/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Routes(app *fiber.App, db *gorm.DB) {
	// Initialize services
	UserService := service.NewUserService(db)

	v1 := app.Group("/v1")
	api := app.Group("/api")

	UserRoutes(v1, UserService)
	UserRoutes(api, UserService)

    
    // Register Portfolio Routes
    PortfolioRoutes(app, db)
	NavRoutes(app, db)
	ProjectRoutes(app, db)
	CompareRoutes(app, db)
	
	// Health check endpoint
	v1.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "Server is running",
		})
	})
}

