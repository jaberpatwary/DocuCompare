package model

import "gorm.io/gorm"

// NavItem represents a navigation menu item
type NavItem struct {
	gorm.Model
	Label  string `json:"label"`
	Link   string `json:"link"`   // URL or anchor (#about)
	Order  int    `json:"order"`  // Display order
}
