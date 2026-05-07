package service

import (
	"path/filepath"
	"strings"
)

type DocumentService interface {
	ExtractTextFromPDF(path string) (string, error)
	ExtractTextFromDocx(path string) (string, error)
}

type documentService struct {
	ocr OCRService
}

func NewDocumentService(ocr OCRService) DocumentService {
	return &documentService{ocr: ocr}
}

func (s *documentService) ExtractTextFromPDF(path string) (string, error) {
	// For testing, return dummy text
	return "Dummy PDF text", nil
}
func (s *documentService) ExtractTextFromDocx(path string) (string, error) {
	// For testing, return dummy text
	return "Dummy DOCX text", nil
}

func GetFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}
