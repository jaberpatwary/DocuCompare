package router

import (
	"app/src/controller"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProjectRoutes(app *fiber.App, db *gorm.DB) {
	pc := controller.NewProjectController(db)
	
	api := app.Group("/api")
	
	// Public Routes
	api.Get("/projects", pc.GetProjects)
	api.Get("/projects/:id", pc.GetProjectByID)
	
	// Admin Routes (should be protected with middleware in production)
	api.Post("/projects", pc.AddProject)
	api.Put("/projects/:id", pc.UpdateProject)
	api.Delete("/projects/:id", pc.DeleteProject)
	api.Patch("/projects/:id/featured", pc.ToggleFeatured)
	
	// Legacy endpoint for compatibility
	api.Get("/articles", pc.GetProjects)
}
