package model

import (
	"time"
)

type User struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	Email        string    `gorm:"type:varchar(150);unique;not null" json:"email"`
	Phone        string    `gorm:"type:varchar(20)" json:"phone"`
	PasswordHash string    `gorm:"type:text;not null" json:"-"` // never expose in JSON
	AvatarURL    string    `gorm:"type:text" json:"avatar_url"`
	Status       string    `gorm:"type:varchar(20);not null;default:'active';check:status IN ('active','inactive','banned')" json:"status"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
