package utils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// NormalizeText performs production-level cleanup for Bengali and English text
func NormalizeText(text string) string {
	// 1. Unicode Normalization (NFC)
	text = norm.NFC.String(text)

	// 2. Lowercase (for English)
	text = strings.ToLower(text)

	// 3. Remove punctuation (optional, but good for similarity)
	// We keep Bengali punctuation like । (Dari) if needed, but for comparison we often strip them
	reg := regexp.MustCompile(`[!"#$%&'()*+,\-./:;<=>?@[\\\]^_{|}~।]`)
	text = reg.ReplaceAllString(text, " ")

	// 4. Remove special non-printable characters
	text = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return ' '
	}, text)

	// 5. Replace multiple spaces/newlines with a single space
	spaceReg := regexp.MustCompile(`\s+`)
	text = spaceReg.ReplaceAllString(text, " ")

	return strings.TrimSpace(text)
}
