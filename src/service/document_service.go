package service

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ledongthuc/pdf"
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
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open PDF: %w", err)
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", fmt.Errorf("failed to get PDF text: %w", err)
	}
	
	_, _ = buf.ReadFrom(b)
	text := buf.String()

	// Fallback to OCR if text is suspicious/empty (scanned PDF)
	if len(strings.TrimSpace(text)) < 10 {
		// In a real production system, you'd convert PDF pages to images here
		// For now, we return a clear message or trigger OCR if the file allows
		return "", fmt.Errorf("PDF appears to be scanned or empty, please upload as image for OCR")
	}

	return text, nil
}

func (s *documentService) ExtractTextFromDocx(path string) (string, error) {
	// Open the docx file
	r, err := docx.ReadDocxFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to open DOCX: %w", err)
	}
	defer r.Close()

	// Extract text
	docxObj := r.Editable()
	rawXML := docxObj.GetContent()

	// Parse plain text out of the raw XML using a simple regex to match <w:t>...</w:t>
	// or we can strip all XML tags. A robust way is removing all <...> tags.
	var buf bytes.Buffer
	inTag := false
	for _, char := range rawXML {
		if char == '<' {
			inTag = true
			continue
		}
		if char == '>' {
			inTag = false
			continue
		}
		if !inTag {
			buf.WriteRune(char)
		}
	}

	text := strings.TrimSpace(buf.String())
	return text, nil
}

func GetFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}
