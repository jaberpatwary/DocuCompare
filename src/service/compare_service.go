package service

import (
	"app/src/model"
	"app/src/utils"
	"bytes"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sergi/go-diff/diffmatchpatch"
	"gorm.io/gorm"
)

type CompareService interface {
	CompareDocuments(c *fiber.Ctx) (*model.CompareHistory, error)
	ExtractDocument(c *fiber.Ctx) (string, error)
	GetHistory(c *fiber.Ctx) ([]model.CompareHistory, error)
	GetHistoryByID(c *fiber.Ctx, id string) (*model.CompareHistory, error)
	DeleteHistory(c *fiber.Ctx, id string) error
}

type compareService struct {
	db       *gorm.DB
	ocr      OCRService
	document DocumentService
}

func NewCompareService(db *gorm.DB, ocr OCRService, document DocumentService) CompareService {
	return &compareService{
		db:       db,
		ocr:      ocr,
		document: document,
	}
}

// saveUploadedFile copies a multipart file to the destination path
func saveUploadedFile(file *multipart.FileHeader, dest string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// isImageFile checks if the filename has an image extension
func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

// isPDFFile checks if the filename is a PDF
func isPDFFile(filename string) bool {
	return strings.ToLower(filepath.Ext(filename)) == ".pdf"
}

// isDocxFile checks if the filename is a DOCX
func isDocxFile(filename string) bool {
	return strings.ToLower(filepath.Ext(filename)) == ".docx"
}

// extractTextByType routes to the correct extraction method based on file type
func (s *compareService) extractTextByType(filePath string, filename string, lang string) (string, error) {
	switch {
	case isImageFile(filename):
		utils.Log.Infof("Running OCR on image: %s", filename)
		text, err := s.ocr.ExtractTextFromImage(filePath, lang)
		if err != nil {
			return "", fmt.Errorf("OCR failed for %s: %w", filename, err)
		}
		return text, nil

	case isPDFFile(filename):
		utils.Log.Infof("Extracting text from PDF: %s", filename)
		text, err := s.document.ExtractTextFromPDF(filePath)
		if err != nil {
			return "", fmt.Errorf("PDF extraction failed for %s: %w", filename, err)
		}
		return text, nil

	case isDocxFile(filename):
		utils.Log.Infof("Extracting text from DOCX: %s", filename)
		text, err := s.document.ExtractTextFromDocx(filePath)
		if err != nil {
			return "", fmt.Errorf("DOCX extraction failed for %s: %w", filename, err)
		}
		return text, nil

	default:
		return "", fmt.Errorf("unsupported file type: %s", filepath.Ext(filename))
	}
}

// cosineSimilarity computes similarity between two word frequency maps
func cosineSimilarity(words1, words2 []string) float64 {
	freq1 := make(map[string]float64)
	freq2 := make(map[string]float64)

	for _, w := range words1 {
		freq1[w]++
	}
	for _, w := range words2 {
		freq2[w]++
	}

	// Dot product
	var dot float64
	for word, count := range freq1 {
		dot += count * freq2[word]
	}

	// Magnitudes
	mag1, mag2 := 0.0, 0.0
	for _, count := range freq1 {
		mag1 += count * count
	}
	for _, count := range freq2 {
		mag2 += count * count
	}

	if mag1 == 0 || mag2 == 0 {
		return 0
	}

	return dot / (math.Sqrt(mag1) * math.Sqrt(mag2))
}

func (s *compareService) CompareDocuments(c *fiber.Ctx) (*model.CompareHistory, error) {
	startTime := time.Now()

	file1, err := c.FormFile("file1")
	if err != nil {
		return nil, fmt.Errorf("file1 is required")
	}
	file2, err := c.FormFile("file2")
	if err != nil {
		return nil, fmt.Errorf("file2 is required")
	}
	language := c.FormValue("language", "bn")

	// Day-wise organized upload directory
	today := time.Now().Format("2006-01-02")
	uploadDir := filepath.Join("./frontend/uploads", today)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Secure filenames: timestamp + original name
	path1 := filepath.Join(uploadDir, fmt.Sprintf("%d_%s", time.Now().UnixNano(), file1.Filename))
	path2 := filepath.Join(uploadDir, fmt.Sprintf("%d_%s", time.Now().UnixNano(), file2.Filename))

	if err := saveUploadedFile(file1, path1); err != nil {
		return nil, fmt.Errorf("failed to save file1: %w", err)
	}
	defer os.Remove(path1) // Cleanup after processing

	if err := saveUploadedFile(file2, path2); err != nil {
		return nil, fmt.Errorf("failed to save file2: %w", err)
	}
	defer os.Remove(path2)

	// Extract text from both documents using real extraction
	text1, err := s.extractTextByType(path1, file1.Filename, language)
	if err != nil {
		utils.Log.Errorf("Extraction failed for file1: %v", err)
		return nil, err
	}

	text2, err := s.extractTextByType(path2, file2.Filename, language)
	if err != nil {
		utils.Log.Errorf("Extraction failed for file2: %v", err)
		return nil, err
	}

	// Normalize both texts (Bengali-friendly)
	text1 = utils.NormalizeText(text1)
	text2 = utils.NormalizeText(text2)

	if strings.TrimSpace(text1) == "" || strings.TrimSpace(text2) == "" {
		return nil, fmt.Errorf("one or both documents produced no extractable text. Check file quality or format")
	}

	// Word-level diff using diffmatchpatch
	dmp := diffmatchpatch.New()
	words1 := strings.Fields(text1)
	words2 := strings.Fields(text2)

	text1Lines := strings.Join(words1, "\n")
	text2Lines := strings.Join(words2, "\n")

	diffs := dmp.DiffMain(text1Lines, text2Lines, false)
	diffs = dmp.DiffCleanupSemantic(diffs)

	var missing, extra int
	var resultJSON bytes.Buffer
	resultJSON.WriteString("[")
	first := true

	for _, d := range diffs {
		wordStr := strings.ReplaceAll(strings.ReplaceAll(d.Text, "\n", " "), `"`, `\"`)
		wordStr = strings.TrimSpace(wordStr)
		if wordStr == "" {
			continue
		}
		if !first {
			resultJSON.WriteString(",")
		}
		first = false

		switch d.Type {
		case diffmatchpatch.DiffEqual:
			resultJSON.WriteString(fmt.Sprintf(`{"type":"equal","text":"%s"}`, wordStr))
		case diffmatchpatch.DiffInsert:
			extra += len(strings.Fields(d.Text))
			resultJSON.WriteString(fmt.Sprintf(`{"type":"insert","text":"%s"}`, wordStr))
		case diffmatchpatch.DiffDelete:
			missing += len(strings.Fields(d.Text))
			resultJSON.WriteString(fmt.Sprintf(`{"type":"delete","text":"%s"}`, wordStr))
		}
	}
	resultJSON.WriteString("]")

	mismatched := min(missing, extra)
	missing -= mismatched
	extra -= mismatched

	totalWords := len(words1)
	if len(words2) > totalWords {
		totalWords = len(words2)
	}

	// Cosine similarity for more accurate percentage
	cosineSim := cosineSimilarity(words1, words2) * 100.0
	processingMs := int(time.Since(startTime).Milliseconds())

	history := &model.CompareHistory{
		FirstDocumentName:  file1.Filename,
		FirstDocumentURL:   path1,
		FirstDocumentText:  text1,
		SecondDocumentName: file2.Filename,
		SecondDocumentURL:  path2,
		SecondDocumentText: text2,
		Language:           language,
		SimilarityScore:    math.Round(cosineSim*100) / 100,
		MismatchedWords:    mismatched,
		MissingWords:       missing,
		ExtraWords:         extra,
		TotalWordsCompared: totalWords,
		CompareResult:      resultJSON.String(),
		ProcessingTimeMs:   processingMs,
		FileSize1:          file1.Size,
		FileSize2:          file2.Size,
		FileType1:          strings.ToLower(filepath.Ext(file1.Filename)),
		FileType2:          strings.ToLower(filepath.Ext(file2.Filename)),
		Status:             "completed",
	}

	if err := s.db.Create(history).Error; err != nil {
		utils.Log.Errorf("Failed to save comparison history: %v", err)
	}

	utils.Log.Infof("Comparison complete: %.2f%% similarity, %dms processing", cosineSim, processingMs)
	return history, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (s *compareService) GetHistory(c *fiber.Ctx) ([]model.CompareHistory, error) {
	var history []model.CompareHistory
	query := s.db.Order("created_at desc")

	// Filter by date if provided
	if date := c.Query("date"); date != "" {
		query = query.Where("DATE(created_at) = ?", date)
	}
	// Filter by language
	if lang := c.Query("language"); lang != "" {
		query = query.Where("language = ?", lang)
	}
	// Limit
	query = query.Limit(100)

	if err := query.Find(&history).Error; err != nil {
		return nil, err
	}
	return history, nil
}

func (s *compareService) GetHistoryByID(c *fiber.Ctx, id string) (*model.CompareHistory, error) {
	var history model.CompareHistory
	if err := s.db.First(&history, id).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func (s *compareService) DeleteHistory(c *fiber.Ctx, id string) error {
	return s.db.Delete(&model.CompareHistory{}, id).Error
}

func (s *compareService) ExtractDocument(c *fiber.Ctx) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", fmt.Errorf("file is required")
	}
	language := c.FormValue("language", "bn")

	today := time.Now().Format("2006-01-02")
	uploadDir := filepath.Join("./frontend/uploads", today)
	os.MkdirAll(uploadDir, os.ModePerm)

	path := filepath.Join(uploadDir, fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename))
	if err := saveUploadedFile(file, path); err != nil {
		return "", err
	}
	defer os.Remove(path)

	text, err := s.extractTextByType(path, file.Filename, language)
	if err != nil {
		return "", err
	}

	return text, nil
}
