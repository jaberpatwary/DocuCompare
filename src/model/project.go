package model

import "gorm.io/gorm"

// Project represents a portfolio project/article
type Project struct {
	gorm.Model
	Title       string `json:"title"`
	Content     string `json:"content"`     // Full description
	Summary     string `json:"summary"`     // Short description
	Image       string `json:"image"`       // Project image URL
	Category    string `json:"category"`    // Comma-separated tags/technologies
	GithubLink  string `json:"github_link"` // GitHub repository URL
	LiveLink    string `json:"live_link"`   // Live demo URL
	Featured    bool   `json:"featured"`    // Is this a featured project?
	Order       int    `json:"order"`       // Display order
}
