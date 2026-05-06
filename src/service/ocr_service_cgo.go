//go:build cgo
// +build cgo

package service

import (
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/otiai10/gosseract/v2"
)

type OCRService interface {
	ExtractTextFromImage(filePath string, lang string) (string, error)
	PreprocessImage(filePath string) (string, error)
}

type ocrService struct{}

func NewOCRService() OCRService {
	return &ocrService{}
}

func (s *ocrService) PreprocessImage(filePath string) (string, error) {
	// Load image
	src, err := imaging.Open(filePath)
	if err != nil {
		return "", err
	}

	// 1. Grayscale
	img := imaging.Grayscale(src)

	// 2. Adjust Contrast
	img = imaging.AdjustContrast(img, 20)

	// 3. Sharpen
	img = imaging.Sharpen(img, 0.5)

	// Save preprocessed image to a temporary file
	processedPath := strings.TrimSuffix(filePath, ".png") + "_processed.png"
	err = imaging.Save(img, processedPath)
	if err != nil {
		return "", err
	}

	return processedPath, nil
}

func (s *ocrService) ExtractTextFromImage(filePath string, lang string) (string, error) {
	// Preprocess first for production-level accuracy
	processedPath, err := s.PreprocessImage(filePath)
	if err == nil {
		defer os.Remove(processedPath) // Cleanup temp processed image
		filePath = processedPath
	}

	client := gosseract.NewClient()
	defer client.Close()
    // ... rest of implementation stays the same

	// Set language based on request
	// gosseract requires Tesseract language packs installed on the system (e.g. tesseract-ocr-ben, tesseract-ocr-eng)
	if lang == "bn" {
		client.SetLanguage("ben", "eng") // Bangla and English together for mixed documents
	} else {
		client.SetLanguage("eng")
	}

	err := client.SetImage(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to set image for OCR: %w", err)
	}

	text, err := client.Text()
	if err != nil {
		return "", fmt.Errorf("failed to extract text via OCR: %w", err)
	}

	return strings.TrimSpace(text), nil
}
