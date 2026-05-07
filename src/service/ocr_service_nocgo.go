//go:build !cgo
// +build !cgo

package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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
	msg := fmt.Sprintf("Using Remote OCR.space API for %s because local Tesseract is disabled.", filePath)
	fmt.Println(msg)
	return CallOCRSpaceRemote(filePath, lang)
}

// Global helper for OCR API extraction
func CallOCRSpaceRemote(filePath string, lang string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	writer.WriteField("apikey", "helloworld")
	
	// Map lang 'bn' to Bengali or English if unsupported
	ocrLang := "eng"
	if lang == "bn" { ocrLang = "eng" } // Free tier supports basic languages, 'eng' works well
	
	writer.WriteField("language", ocrLang) 
	writer.Close()

	req, err := http.NewRequest("POST", "https://api.ocr.space/parse/image", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		ParsedResults []struct {
			ParsedText string `json:"ParsedText"`
		} `json:"ParsedResults"`
		ErrorMessage []string `json:"ErrorMessage"`
		IsErroredOnProcessing bool `json:"IsErroredOnProcessing"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.IsErroredOnProcessing {
		return "", fmt.Errorf("OCR API Error: %v", result.ErrorMessage)
	}

	if len(result.ParsedResults) > 0 {
		return result.ParsedResults[0].ParsedText, nil
	}
	return "", fmt.Errorf("no text found")
}
