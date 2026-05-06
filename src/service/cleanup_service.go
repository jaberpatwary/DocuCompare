package service

import (
	"app/src/utils"
	"os"
	"path/filepath"
	"time"
)

type CleanupService interface {
	CleanOldUploads(maxAgeDays int) error
}

type cleanupService struct{}

func NewCleanupService() CleanupService {
	return &cleanupService{}
}

func (s *cleanupService) CleanOldUploads(maxAgeDays int) error {
	uploadDir := "./frontend/uploads"
	cutoff := time.Now().AddDate(0, 0, -maxAgeDays)
	cleaned := 0

	err := filepath.Walk(uploadDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip errors
		}
		if info.IsDir() {
			return nil
		}
		// Skip .gitkeep
		if info.Name() == ".gitkeep" {
			return nil
		}
		// Remove files older than cutoff
		if info.ModTime().Before(cutoff) {
			if removeErr := os.Remove(path); removeErr == nil {
				cleaned++
			}
		}
		return nil
	})

	utils.Log.Infof("Cleanup completed: removed %d old files", cleaned)
	return err
}

// StartCleanupScheduler runs cleanup every 24 hours in the background
func StartCleanupScheduler(maxAgeDays int) {
	svc := NewCleanupService()
	go func() {
		for {
			time.Sleep(24 * time.Hour)
			if err := svc.CleanOldUploads(maxAgeDays); err != nil {
				utils.Log.Errorf("Cleanup error: %v", err)
			}
		}
	}()
	utils.Log.Info("File cleanup scheduler started (every 24h)")
}
