package router

import "github.com/gofiber/fiber/v2"

func FrontendRoutes(app *fiber.App) {
	// Serve the entire frontend folder at root so img/, css/, js/ all work directly
	app.Static("/", "./frontend")
	app.Static("/uploads", "./frontend/uploads")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./frontend/index.html")
	})

	app.Get("/admin.html", func(c *fiber.Ctx) error {
		return c.SendFile("./frontend/admin.html")
	})

	app.Get("/admin", func(c *fiber.Ctx) error {
		return c.SendFile("./frontend/admin.html")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendFile("./frontend/login.html")
	})

	app.Get("/portfolio", func(c *fiber.Ctx) error {
		return c.SendFile("./frontend/index.html")
	})

	app.Get("/docucompare", func(c *fiber.Ctx) error {
		return c.SendFile("./frontend/docucompare.html")
	})
}
