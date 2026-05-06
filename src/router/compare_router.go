package router

import (
	"app/src/controller"
	"app/src/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CompareRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize Services
	ocrService := service.NewOCRService()
	compareService := service.NewCompareService(db, ocrService)
	
	// Initialize the Controller
	compareController := controller.NewCompareController(compareService)

	// API Group
	api := app.Group("/api/v1/compare")
	
	// Define compare routes
	api.Post("/process", compareController.Compare)
	api.Get("/history", compareController.GetHistory)
	api.Get("/history/:id", compareController.GetHistoryByID)
	api.Delete("/history/:id", compareController.DeleteHistory)
}
