package middleware

import (
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".pdf":  true,
	".docx": true,
}

var allowedMimeTypes = map[string]bool{
	"image/jpeg":                                                               true,
	"image/png":                                                                true,
	"application/pdf":                                                          true,
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document":   true,
}

const maxFileSize = 20 * 1024 * 1024 // 20MB

// FileValidation validates uploaded files for security
func FileValidation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.Next() // Not a multipart request, skip
		}

		for _, files := range form.File {
			for _, file := range files {
				// Check file size
				if file.Size > maxFileSize {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "File too large. Maximum size is 20MB.",
					})
				}

				// Check extension
				ext := strings.ToLower(filepath.Ext(file.Filename))
				if !allowedExtensions[ext] {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Invalid file type: " + ext + ". Allowed: jpg, jpeg, png, pdf, docx",
					})
				}

				// Check MIME type
				contentType := file.Header.Get("Content-Type")
				if contentType != "" && !allowedMimeTypes[contentType] {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Invalid MIME type: " + contentType,
					})
				}
			}
		}

		return c.Next()
	}
}
