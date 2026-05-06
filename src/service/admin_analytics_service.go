package service

import (
	"app/src/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AdminAnalyticsService interface {
	GetDashboardStats(c *fiber.Ctx) (*DashboardStats, error)
	GetDailyUploads(c *fiber.Ctx) ([]DailyUploadStat, error)
	GetRecentActivity(c *fiber.Ctx) ([]model.CompareHistory, error)
}

type DashboardStats struct {
	TotalComparisons    int64   `json:"total_comparisons"`
	TotalToday          int64   `json:"total_today"`
	AvgSimilarity       float64 `json:"avg_similarity"`
	TotalPDFProcessed   int64   `json:"total_pdf_processed"`
	TotalDocxProcessed  int64   `json:"total_docx_processed"`
	TotalImageProcessed int64   `json:"total_image_processed"`
	AvgProcessingMs     float64 `json:"avg_processing_ms"`
}

type DailyUploadStat struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

type adminAnalyticsService struct {
	db *gorm.DB
}

func NewAdminAnalyticsService(db *gorm.DB) AdminAnalyticsService {
	return &adminAnalyticsService{db: db}
}

func (s *adminAnalyticsService) GetDashboardStats(c *fiber.Ctx) (*DashboardStats, error) {
	var stats DashboardStats

	s.db.Model(&model.CompareHistory{}).Count(&stats.TotalComparisons)

	today := time.Now().Format("2006-01-02")
	s.db.Model(&model.CompareHistory{}).Where("DATE(created_at) = ?", today).Count(&stats.TotalToday)

	s.db.Model(&model.CompareHistory{}).Select("AVG(similarity_score)").Scan(&stats.AvgSimilarity)
	s.db.Model(&model.CompareHistory{}).Select("AVG(processing_time_ms)").Scan(&stats.AvgProcessingMs)

	s.db.Model(&model.CompareHistory{}).Where("file_type1 = '.pdf' OR file_type2 = '.pdf'").Count(&stats.TotalPDFProcessed)
	s.db.Model(&model.CompareHistory{}).Where("file_type1 = '.docx' OR file_type2 = '.docx'").Count(&stats.TotalDocxProcessed)
	s.db.Model(&model.CompareHistory{}).Where("file_type1 IN ('.jpg','.jpeg','.png') OR file_type2 IN ('.jpg','.jpeg','.png')").Count(&stats.TotalImageProcessed)

	return &stats, nil
}

func (s *adminAnalyticsService) GetDailyUploads(c *fiber.Ctx) ([]DailyUploadStat, error) {
	var results []DailyUploadStat
	s.db.Model(&model.CompareHistory{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Group("DATE(created_at)").
		Order("date desc").
		Limit(30).
		Scan(&results)
	return results, nil
}

func (s *adminAnalyticsService) GetRecentActivity(c *fiber.Ctx) ([]model.CompareHistory, error) {
	var history []model.CompareHistory
	if err := s.db.Order("created_at desc").Limit(20).Find(&history).Error; err != nil {
		return nil, err
	}
	return history, nil
}
