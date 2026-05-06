package router

import (
	"app/src/controller"
	"app/src/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PortfolioRoutes(app *fiber.App, db *gorm.DB) {
	pc := controller.NewPortfolioController(db)
	
	api := app.Group("/api")

	// Public Routes (for frontend)
	api.Get("/profile", pc.GetProfile)
	api.Get("/experience", pc.GetExperiences)
	api.Get("/achievements", pc.GetAchievements)
	api.Get("/photos", pc.GetPhotos)
	api.Get("/videos", pc.GetVideos)

	// Protected Routes (for admin)
	// We use the JwtConfig middleware to protect these routes
	protected := api.Group("/", middleware.JwtConfig())
	
	protected.Post("/profile", pc.UpdateProfile)
	protected.Post("/experience", pc.AddExperience)
	protected.Put("/experience/:id", pc.UpdateExperience)
	protected.Delete("/experience/:id", pc.DeleteExperience)

	// Achievement Routes
	protected.Post("/achievements", pc.AddAchievement)
	protected.Put("/achievements/:id", pc.UpdateAchievement)
	protected.Delete("/achievements/:id", pc.DeleteAchievement)

	// Photo Routes
	protected.Post("/photos", pc.AddPhoto)
	protected.Delete("/photos/:id", pc.DeletePhoto)

	// Video Routes
	protected.Post("/videos", pc.AddVideo)
	protected.Put("/videos/:id", pc.UpdateVideo)
	protected.Delete("/videos/:id", pc.DeleteVideo)
	
	// Image Upload (if needed, adding it to protected group)
	// protected.Post("/upload-image", pc.UploadImage)
}

