package controller

import (
	"app/src/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type NavController struct {
	DB *gorm.DB
}

func NewNavController(db *gorm.DB) *NavController {
	return &NavController{DB: db}
}

// EnsureDefaults seeds the table if empty
func (nc *NavController) EnsureDefaults() {
	var count int64
	nc.DB.Model(&model.NavItem{}).Count(&count)
	if count == 0 {
		defaults := []model.NavItem{
			{Label: "About", Link: "#about", Order: 1},
			{Label: "Experience", Link: "#experience", Order: 2},
			{Label: "Work", Link: "#projects", Order: 3},
			{Label: "Contact", Link: "#contact", Order: 4},
		}
		nc.DB.Create(&defaults)
	}
}

func (nc *NavController) GetNavItems(c *fiber.Ctx) error {
	nc.EnsureDefaults() // Check and seed if necessary on read
	var items []model.NavItem
	nc.DB.Order("\"order\" asc").Find(&items)
	return c.JSON(items)
}

func (nc *NavController) AddNavItem(c *fiber.Ctx) error {
	var item model.NavItem
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
	}
	if result := nc.DB.Create(&item); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create item"})
	}
	return c.JSON(item)
}

func (nc *NavController) UpdateNavItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item model.NavItem
	if err := nc.DB.First(&item, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
	}
	
	// Parse updates
	var payload model.NavItem
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	item.Label = payload.Label
	item.Link = payload.Link
	item.Order = payload.Order
	
	nc.DB.Save(&item)
	return c.JSON(item)
}

func (nc *NavController) DeleteNavItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := nc.DB.Delete(&model.NavItem{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete"})
	}
	return c.JSON(fiber.Map{"message": "Deleted"})
}
