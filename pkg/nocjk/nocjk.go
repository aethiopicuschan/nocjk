package nocjk

import (
	"strings"
	"unicode"
)

// FindChineseLines returns a list of line indices that likely contain Chinese characters (Han only, no Hiragana or Katakana).
func FindChineseLines(text string) []int {
	lines := strings.Split(text, "\n")
	result := make([]int, 0)

	for i, line := range lines {
		hasHan := false
		hasJapanese := false

		for _, r := range line {
			switch {
			case unicode.In(r, unicode.Han):
				hasHan = true
			case unicode.In(r, unicode.Hiragana, unicode.Katakana):
				hasJapanese = true
			}
		}

		if hasHan && !hasJapanese {
			result = append(result, i+1)
		}
	}

	return result
}

// FindJapaneseLines returns a list of line indices that contain Japanese characters.
func FindJapaneseLines(text string) []int {
	lines := strings.Split(text, "\n")
	result := make([]int, 0)

	for i, line := range lines {
		for _, r := range line {
			if unicode.In(r, unicode.Hiragana, unicode.Katakana, unicode.Han) {
				result = append(result, i+1)
				break
			}
		}
	}

	return result
}

// FindKoreanLines returns a list of line indices that contain Korean characters (Hangul).
func FindKoreanLines(text string) []int {
	lines := strings.Split(text, "\n")
	result := make([]int, 0)

	for i, line := range lines {
		for _, r := range line {
			if unicode.In(r, unicode.Hangul) {
				result = append(result, i+1)
				break
			}
		}
	}

	return result
}

// FindCJKLines returns a map of detected line indices for Chinese, Japanese and Korean text.
func FindCJKLines(text string) map[string][]int {
	return map[string][]int{
		"chinese":  FindChineseLines(text),
		"japanese": FindJapaneseLines(text),
		"korean":   FindKoreanLines(text),
	}
}
