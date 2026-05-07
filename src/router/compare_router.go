package router

import (
	"app/src/controller"
	"app/src/middleware"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CompareRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize Services
	ocrService := service.NewOCRService()
	documentService := service.NewDocumentService(ocrService)
	compareService := service.NewCompareService(db, ocrService, documentService)

	// Initialize Controllers
	compareController := controller.NewCompareController(compareService)

	// Compare API Group (with file validation middleware)
	api := app.Group("/api/v1/compare")
	api.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Compare API is working", "status": "ok"})
	})
	api.Post("/process", middleware.FileValidation(), compareController.Compare)
	api.Post("/extract", middleware.FileValidation(), compareController.Extract)
	api.Get("/history", compareController.GetHistory)
	api.Get("/history/:id", compareController.GetHistoryByID)
	api.Delete("/history/:id", compareController.DeleteHistory)
}
