package router

import (
	"app/src/controller"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NavRoutes(app *fiber.App, db *gorm.DB) {
	nc := controller.NewNavController(db)

	api := app.Group("/api")

	api.Get("/nav-items", nc.GetNavItems)
	api.Post("/nav-items", nc.AddNavItem)
	api.Put("/nav-items/:id", nc.UpdateNavItem)
	api.Delete("/nav-items/:id", nc.DeleteNavItem)
}
