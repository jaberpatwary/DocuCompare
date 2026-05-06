//go:build cgo
// +build cgo

package service

import (
	"fmt"
	"strings"

	"github.com/otiai10/gosseract/v2"
)

type OCRService interface {
	ExtractTextFromImage(filePath string, lang string) (string, error)
}

type ocrService struct{}

func NewOCRService() OCRService {
	return &ocrService{}
}

func (s *ocrService) ExtractTextFromImage(filePath string, lang string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

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
