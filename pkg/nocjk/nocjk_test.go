package nocjk_test

import (
	"testing"

	"github.com/aethiopicuschan/nocjk/pkg/nocjk"
	"github.com/stretchr/testify/assert"
)

func TestFindChineseLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{
			name:     "No Chinese",
			input:    "Hello\nこんにちは\n안녕하세요",
			expected: []int{},
		},
		{
			name:     "Only Han characters",
			input:    "你好\nこんにちは\nEnglish",
			expected: []int{1},
		},
		{
			name:     "Mixed with Japanese (should be excluded)",
			input:    "漢字とひらがな\n漢字とカタカナ\n汉字 only",
			expected: []int{3},
		},
		{
			name:     "Multiple Chinese lines",
			input:    "早上好\n晚安\nThis is English.",
			expected: []int{1, 2},
		},
		{
			name:     "Empty lines",
			input:    "\n\n漢字\n",
			expected: []int{3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := nocjk.FindChineseLines(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFindJapaneseLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{
			name:     "No Japanese",
			input:    "This is English.\nAnother line.",
			expected: []int{},
		},
		{
			name:     "One Japanese line",
			input:    "これは日本語の行です。\nThis is English.",
			expected: []int{1},
		},
		{
			name:     "Multiple Japanese lines",
			input:    "こんにちは\n世界\nHello",
			expected: []int{1, 2},
		},
		{
			name:     "Mixed characters in one line",
			input:    "Line1\nEnglish and カタカナ\n123",
			expected: []int{2},
		},
		{
			name:     "Empty lines",
			input:    "\n\n漢字がある行\n",
			expected: []int{3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := nocjk.FindJapaneseLines(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFindKoreanLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{
			name:     "No Korean",
			input:    "Hello\nこんにちは\n你好",
			expected: []int{},
		},
		{
			name:     "One Korean line",
			input:    "안녕하세요\nThis is English.",
			expected: []int{1},
		},
		{
			name:     "Multiple Korean lines",
			input:    "안녕\n하세요\nHello",
			expected: []int{1, 2},
		},
		{
			name:     "Mixed content with Hangul",
			input:    "Line1\nEnglish and 한글\n123",
			expected: []int{2},
		},
		{
			name:     "Empty lines",
			input:    "\n\n한글이 있는 줄\n",
			expected: []int{3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := nocjk.FindKoreanLines(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFindCJKLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string][]int
	}{
		{
			name:  "All scripts present",
			input: "你好\nこんにちは\n안녕하세요\nHello",
			expected: map[string][]int{
				"chinese":  []int{1},
				"japanese": []int{1, 2},
				"korean":   []int{3},
			},
		},
		{
			name:  "Only Japanese and Korean",
			input: "こんにちは\n안녕하세요\nEnglish",
			expected: map[string][]int{
				"chinese":  []int{},
				"japanese": []int{1},
				"korean":   []int{2},
			},
		},
		{
			name:  "Only Chinese with Japanese exclusion",
			input: "汉字\nカタカナと漢字\nEnglish",
			expected: map[string][]int{
				"chinese":  []int{1},
				"japanese": []int{1, 2},
				"korean":   []int{},
			},
		},
		{
			name:  "Empty input",
			input: "",
			expected: map[string][]int{
				"chinese":  []int{},
				"japanese": []int{},
				"korean":   []int{},
			},
		},
		{
			name:  "Multiple matches per category",
			input: "你好\n早上好\nこんにちは\n世界\n안녕\n하세요",
			expected: map[string][]int{
				"chinese":  []int{1, 2, 4},
				"japanese": []int{1, 2, 3, 4},
				"korean":   []int{5, 6},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := nocjk.FindCJKLines(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
