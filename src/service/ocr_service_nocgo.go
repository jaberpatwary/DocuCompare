//go:build !cgo
// +build !cgo

package service

import (
	"errors"
	"fmt"
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
	return filePath, nil
}

func (s *ocrService) ExtractTextFromImage(filePath string, lang string) (string, error) {
	// When CGO is disabled (like on standard Windows environments without MSYS2/MinGW),
	// gosseract cannot be compiled because it depends on C++ Tesseract headers.
	// This fallback prevents the whole Fiber app from crashing during "go run".
	// To use real OCR, you must build the app with CGO_ENABLED=1 (e.g. via Docker).
	
	msg := fmt.Sprintf("⚠️ Real OCR skipped for %s. CGO is disabled locally on Windows. Please run via Docker with Tesseract installed to use real gosseract OCR.", filePath)
	fmt.Println(msg)
	
	return "", errors.New("CGO disabled: gosseract OCR unavailable")
}
