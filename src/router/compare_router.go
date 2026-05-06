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
	analyticsService := service.NewAdminAnalyticsService(db)

	// Initialize Controllers
	compareController := controller.NewCompareController(compareService)
	adminController := controller.NewAdminController(analyticsService)

	// Compare API Group (with file validation middleware)
	api := app.Group("/api/v1/compare")
	api.Post("/process", middleware.FileValidation(), compareController.Compare)
	api.Get("/history", compareController.GetHistory)
	api.Get("/history/:id", compareController.GetHistoryByID)
	api.Delete("/history/:id", compareController.DeleteHistory)

	// Admin Analytics API Group (protected with JWT)
	admin := app.Group("/api/v1/admin")
	admin.Get("/stats", adminController.GetDashboardStats)
	admin.Get("/daily-uploads", adminController.GetDailyUploads)
	admin.Get("/recent-activity", adminController.GetRecentActivity)
}
