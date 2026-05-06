package controller

import (
	"app/src/model"


	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PortfolioController struct {
	DB *gorm.DB
}

func NewPortfolioController(db *gorm.DB) *PortfolioController {
	return &PortfolioController{DB: db}
}

// --- Profile Handlers ---

func (pc *PortfolioController) GetProfile(c *fiber.Ctx) error {
	var profile model.Profile
	// Always get the first record, if not exists, return empty or default
	if result := pc.DB.First(&profile); result.Error != nil {
		// Return empty struct if no profile found, or seed a default one
		return c.JSON(model.Profile{Name: "Your Name", Headline: "Your Headline"})
	}
	return c.JSON(profile)
}

func (pc *PortfolioController) UpdateProfile(c *fiber.Ctx) error {
	var payload model.Profile
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	var profile model.Profile
	if result := pc.DB.First(&profile); result.Error != nil {
		// New Profile
		pc.DB.Create(&payload)
	} else {
		// Update existing
		profile.Name = payload.Name
		profile.Headline = payload.Headline
		profile.Bio = payload.Bio
		profile.About = payload.About
		profile.Image = payload.Image
		profile.Email = payload.Email
		profile.Github = payload.Github
		profile.Linkedin = payload.Linkedin
		profile.Resume = payload.Resume
		profile.Instagram = payload.Instagram
		profile.Twitter = payload.Twitter
		profile.Facebook = payload.Facebook
		profile.Youtube = payload.Youtube
		profile.Skills = payload.Skills
		pc.DB.Save(&profile)
	}

	return c.JSON(fiber.Map{"message": "Profile updated successfully"})
}

// --- Experience Handlers ---

func (pc *PortfolioController) GetExperiences(c *fiber.Ctx) error {
	var experiences []model.Experience
	// Order by start_date desc usually, but string date might sort weirdly. 
	// For now just fetch all.
	pc.DB.Order("id desc").Find(&experiences)
	return c.JSON(experiences)
}

func (pc *PortfolioController) AddExperience(c *fiber.Ctx) error {
	var experience model.Experience
	if err := c.BodyParser(&experience); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if result := pc.DB.Create(&experience); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not save experience"})
	}

	return c.JSON(experience)
}

func (pc *PortfolioController) UpdateExperience(c *fiber.Ctx) error {
	id := c.Params("id")
	var experience model.Experience
	
	if err := pc.DB.First(&experience, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Experience not found"})
	}

	if err := c.BodyParser(&experience); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	pc.DB.Save(&experience)
	return c.JSON(experience)
}

func (pc *PortfolioController) DeleteExperience(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := pc.DB.Delete(&model.Experience{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete experience"})
	}
	return c.JSON(fiber.Map{"message": "Experience deleted"})
}

// --- Achievement Handlers ---

func (pc *PortfolioController) GetAchievements(c *fiber.Ctx) error {
	var achievements []model.Achievement
	pc.DB.Order("id desc").Find(&achievements)
	return c.JSON(achievements)
}

func (pc *PortfolioController) AddAchievement(c *fiber.Ctx) error {
	var achievement model.Achievement
	if err := c.BodyParser(&achievement); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	pc.DB.Create(&achievement)
	return c.JSON(achievement)
}

func (pc *PortfolioController) UpdateAchievement(c *fiber.Ctx) error {
	id := c.Params("id")
	var achievement model.Achievement
	if err := pc.DB.First(&achievement, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Achievement not found"})
	}
	if err := c.BodyParser(&achievement); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	pc.DB.Save(&achievement)
	return c.JSON(achievement)
}

func (pc *PortfolioController) DeleteAchievement(c *fiber.Ctx) error {
	id := c.Params("id")
	pc.DB.Delete(&model.Achievement{}, id)
	return c.JSON(fiber.Map{"message": "Achievement deleted"})
}

// --- Photo Handlers ---

func (pc *PortfolioController) GetPhotos(c *fiber.Ctx) error {
	var photos []model.Photo
	pc.DB.Order("id desc").Find(&photos)
	return c.JSON(photos)
}

func (pc *PortfolioController) AddPhoto(c *fiber.Ctx) error {
	var photo model.Photo
	if err := c.BodyParser(&photo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	pc.DB.Create(&photo)
	return c.JSON(photo)
}

func (pc *PortfolioController) DeletePhoto(c *fiber.Ctx) error {
	id := c.Params("id")
	pc.DB.Delete(&model.Photo{}, id)
	return c.JSON(fiber.Map{"message": "Photo deleted"})
}

// --- Video Handlers ---

func (pc *PortfolioController) GetVideos(c *fiber.Ctx) error {
	var videos []model.Video
	pc.DB.Order("id desc").Find(&videos)
	return c.JSON(videos)
}

func (pc *PortfolioController) AddVideo(c *fiber.Ctx) error {
	var video model.Video
	if err := c.BodyParser(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	pc.DB.Create(&video)
	return c.JSON(video)
}

func (pc *PortfolioController) UpdateVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	var video model.Video
	if err := pc.DB.First(&video, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Video not found"})
	}
	if err := c.BodyParser(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	pc.DB.Save(&video)
	return c.JSON(video)
}

func (pc *PortfolioController) DeleteVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	pc.DB.Delete(&model.Video{}, id)
	return c.JSON(fiber.Map{"message": "Video deleted"})
}
