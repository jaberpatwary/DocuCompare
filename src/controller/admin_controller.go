package controller

import (
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	analyticsService service.AdminAnalyticsService
}

func NewAdminController(analyticsService service.AdminAnalyticsService) *AdminController {
	return &AdminController{analyticsService: analyticsService}
}

func (ctrl *AdminController) GetDashboardStats(c *fiber.Ctx) error {
	stats, err := ctrl.analyticsService.GetDashboardStats(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   stats,
	})
}

func (ctrl *AdminController) GetDailyUploads(c *fiber.Ctx) error {
	uploads, err := ctrl.analyticsService.GetDailyUploads(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   uploads,
	})
}

func (ctrl *AdminController) GetRecentActivity(c *fiber.Ctx) error {
	activity, err := ctrl.analyticsService.GetRecentActivity(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   activity,
	})
}
