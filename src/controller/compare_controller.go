package controller

import (
	"app/src/response"
	"app/src/service"
	"github.com/gofiber/fiber/v2"
)

type CompareController struct {
	_CompareService service.CompareService
}

func NewCompareController(compareService service.CompareService) *CompareController {
	return &CompareController{
		_CompareService: compareService,
	}
}

// Compare Documents Endpoint
func (cc *CompareController) Compare(c *fiber.Ctx) error {
	result, err := cc._CompareService.CompareDocuments(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// Get All History Endpoint
func (cc *CompareController) GetHistory(c *fiber.Ctx) error {
	history, err := cc._CompareService.GetHistory(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(history)
}

// Get History By ID Endpoint
func (cc *CompareController) GetHistoryByID(c *fiber.Ctx) error {
	id := c.Params("id")
	history, err := cc._CompareService.GetHistoryByID(c, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Record not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(history)
}

// Delete History Endpoint
func (cc *CompareController) DeleteHistory(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := cc._CompareService.DeleteHistory(c, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(response.Common{
		Code:    fiber.StatusOK,
		Status:  "success",
		Message: "History deleted successfully",
	})
}

// Extract Single Document Endpoint
func (cc *CompareController) Extract(c *fiber.Ctx) error {
	text, err := cc._CompareService.ExtractDocument(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"text": text,
	})
}
