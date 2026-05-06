package controller

import (
	"app/src/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProjectController struct {
	DB *gorm.DB
}

func NewProjectController(db *gorm.DB) *ProjectController {
	return &ProjectController{DB: db}
}

// GetProjects returns all projects
func (pc *ProjectController) GetProjects(c *fiber.Ctx) error {
	var projects []model.Project
	
	// Get limit from query params (default 10)
	limit := c.QueryInt("limit", 10)
	
	// Order by featured first, then by order
	pc.DB.Order("featured DESC, \"order\" ASC, id DESC").Limit(limit).Find(&projects)
	
	return c.JSON(projects)
}

// GetProjectByID returns a single project
func (pc *ProjectController) GetProjectByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var project model.Project
	
	if err := pc.DB.First(&project, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
	}
	
	return c.JSON(project)
}

// AddProject creates a new project
func (pc *ProjectController) AddProject(c *fiber.Ctx) error {
	var project model.Project
	
	if err := c.BodyParser(&project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	if result := pc.DB.Create(&project); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create project"})
	}
	
	return c.JSON(project)
}

// UpdateProject updates an existing project
func (pc *ProjectController) UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	var project model.Project
	
	if err := pc.DB.First(&project, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
	}
	
	if err := c.BodyParser(&project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	
	pc.DB.Save(&project)
	return c.JSON(project)
}

// DeleteProject deletes a project
func (pc *ProjectController) DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	
	if err := pc.DB.Delete(&model.Project{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete project"})
	}
	
	return c.JSON(fiber.Map{"message": "Project deleted successfully"})
}

// ToggleFeatured toggles the featured status of a project
func (pc *ProjectController) ToggleFeatured(c *fiber.Ctx) error {
	id := c.Params("id")
	var project model.Project
	
	if err := pc.DB.First(&project, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
	}
	
	project.Featured = !project.Featured
	pc.DB.Save(&project)
	
	return c.JSON(project)
}
