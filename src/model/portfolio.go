package model

import "gorm.io/gorm"

// Profile stores the main user information
type Profile struct {
	gorm.Model
	Name     string `json:"name"`
	Headline string `json:"headline"`
	Bio      string `json:"bio"`
	About    string `json:"about"` // Detailed about section
	Image    string `json:"image"`
	Email    string `json:"email"`
	Github   string `json:"github"`
	Linkedin string `json:"linkedin"`
	Resume   string `json:"resume"` // URL to resume
	Instagram string `json:"instagram"`
	Twitter   string `json:"twitter"`
	Facebook  string `json:"facebook"`
	Youtube   string `json:"youtube"`
	Skills    string `json:"skills"` // Comma separated list of skills
}

// Experience stores work history
type Experience struct {
	gorm.Model
	Company     string   `json:"company"`
	Role        string   `json:"role"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Description string   `json:"description"`
	Technologies string  `json:"technologies"` // Comma separated
	IsCurrent   bool     `json:"is_current"`
}

// Achievement stores awards and recognitions
type Achievement struct {
	gorm.Model
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

// Photo stores gallery images
type Photo struct {
	gorm.Model
	Caption string `json:"caption"`
	Image   string `json:"image"`
}

// Video stores youtube videos
type Video struct {
	gorm.Model
	Title       string `json:"title"`
	YoutubeLink string `json:"youtube_link"`
	Description string `json:"description"`
}
