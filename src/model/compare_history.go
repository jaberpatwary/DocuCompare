package model

import (
	"time"
)

type CompareHistory struct {
	ID                 int       `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstDocumentName  string    `gorm:"type:varchar(255);not null" json:"first_document_name"`
	FirstDocumentURL   string    `gorm:"type:text;not null" json:"first_document_url"`
	FirstDocumentText  string    `gorm:"type:text" json:"first_document_text"`
	SecondDocumentName string    `gorm:"type:varchar(255);not null" json:"second_document_name"`
	SecondDocumentURL  string    `gorm:"type:text;not null" json:"second_document_url"`
	SecondDocumentText string    `gorm:"type:text" json:"second_document_text"`
	Language           string    `gorm:"type:varchar(50);default:'en'" json:"language"`
	SimilarityScore    float64   `gorm:"type:decimal(5,2)" json:"similarity_score"`
	MismatchedWords    int       `gorm:"default:0" json:"mismatched_words"`
	MissingWords       int       `gorm:"default:0" json:"missing_words"`
	ExtraWords         int       `gorm:"default:0" json:"extra_words"`
	TotalWordsCompared int       `gorm:"default:0" json:"total_words_compared"`
	CompareResult      string    `gorm:"type:jsonb" json:"compare_result"`
	ProcessingTimeMs   int       `gorm:"default:0" json:"processing_time_ms"`
	Status             string    `gorm:"type:varchar(20);default:'pending'" json:"status"`
	ErrorMessage       string    `gorm:"type:text" json:"error_message"`
	UserID             int       `gorm:"index" json:"user_id"`
	CreatedAt          time.Time `gorm:"autoCreateTime;index" json:"created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
