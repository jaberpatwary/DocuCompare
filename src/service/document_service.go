package service

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/nguyenthenguyen/docx"
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
	// 1. Try using the reliable poppler-utils `pdftotext` CLI first (if installed)
	cmd := exec.Command("pdftotext", path, "-")
	out, err := cmd.Output()
	
	text := ""
	if err == nil {
		text = strings.TrimSpace(string(out))
	}

	// 2. If it's empty, it means the PDF is an image-based scanned PDF
	// Use the OCR service to read it
	if text == "" {
		ocrText, ocrErr := s.ocr.ExtractTextFromImage(path, "bn")
		if ocrErr == nil {
			text = ocrText
		} else if err != nil {
			return "", err
		}
	}

	return text, nil
}

func (s *documentService) ExtractTextFromDocx(path string) (string, error) {
	r, err := docx.ReadDocxFile(path)
	if err != nil {
		return "", err
	}
	defer r.Close()

	docxDoc := r.Editable()
	return docxDoc.GetContent(), nil
}

func GetFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}
