package router

import (
	"app/src/controller"
	"app/src/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PortfolioRoutes(app *fiber.App, db *gorm.DB) {
	pc := controller.NewPortfolioController(db)

	// Public Routes (for frontend) - no JWT required
	app.Get("/api/profile", pc.GetProfile)
	app.Get("/api/experience", pc.GetExperiences)
	app.Get("/api/achievements", pc.GetAchievements)
	app.Get("/api/photos", pc.GetPhotos)
	app.Get("/api/videos", pc.GetVideos)

	// Protected Routes (for admin) - each route gets JWT middleware individually
	app.Post("/api/profile", middleware.JwtConfig(), pc.UpdateProfile)
	app.Post("/api/experience", middleware.JwtConfig(), pc.AddExperience)
	app.Put("/api/experience/:id", middleware.JwtConfig(), pc.UpdateExperience)
	app.Delete("/api/experience/:id", middleware.JwtConfig(), pc.DeleteExperience)

	// Achievement Routes
	app.Post("/api/achievements", middleware.JwtConfig(), pc.AddAchievement)
	app.Put("/api/achievements/:id", middleware.JwtConfig(), pc.UpdateAchievement)
	app.Delete("/api/achievements/:id", middleware.JwtConfig(), pc.DeleteAchievement)

	// Photo Routes
	app.Post("/api/photos", middleware.JwtConfig(), pc.AddPhoto)
	app.Delete("/api/photos/:id", middleware.JwtConfig(), pc.DeletePhoto)

	// Video Routes
	app.Post("/api/videos", middleware.JwtConfig(), pc.AddVideo)
	app.Put("/api/videos/:id", middleware.JwtConfig(), pc.UpdateVideo)
	app.Delete("/api/videos/:id", middleware.JwtConfig(), pc.DeleteVideo)
}
