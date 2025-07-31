package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantNil bool
		wantLen int
	}{
		{
			name:    "Valid positive size",
			size:    10,
			wantNil: false,
			wantLen: 10,
		},
		{
			name:    "Zero size",
			size:    0,
			wantNil: true,
			wantLen: 0,
		},
		{
			name:    "Negative size",
			size:    -1,
			wantNil: true,
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if tt.wantNil {
				assert.Nil(t, got, "Expected nil slice")
			} else {
				require.NotNil(t, got, "Slice should not be nil")
				assert.Equal(t, tt.wantLen, len(got), "Unexpected slice length")
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "Valid case: multiple elements",
			numbers:  []int{81, 87, 47, 59, 81},
			expected: 87,
			wantErr:  false,
		},
		{
			name:     "Valid case: single element",
			numbers:  []int{42},
			expected: 42,
			wantErr:  false,
		},
		{
			name:     "Empty slice",
			numbers:  []int{},
			expected: 0,
			wantErr:  true,
			errMsg:   "слайс пустой или равен nil",
		},
		{
			name:     "Nil slice",
			numbers:  nil,
			expected: 0,
			wantErr:  true,
			errMsg:   "слайс пустой или равен nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maximum(tt.numbers)

			if tt.wantErr {
				require.Error(t, err)
				assert.EqualError(t, err, tt.errMsg)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, got)
			}
		})
	}
}

func TestSplitIntoChunks(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		expectedCount int // Всегда 8
		expectedSizes []int
	}{
		{
			name:          "16 элементов (по 2 в каждом)",
			input:         make([]int, 16),
			expectedCount: 8,
			expectedSizes: []int{2, 2, 2, 2, 2, 2, 2, 2},
		},
		{
			name:          "17 элементов (3x5 + 2 + 1 пустой)",
			input:         make([]int, 17),
			expectedCount: 8,
			expectedSizes: []int{3, 3, 3, 3, 3, 2, 0, 0},
		},
		{
			name:          "5 элементов (1x5 + 3 пустых)",
			input:         make([]int, 5),
			expectedCount: 8,
			expectedSizes: []int{1, 1, 1, 1, 1, 0, 0, 0},
		},
		{
			name:          "Пустой слайс",
			input:         []int{},
			expectedCount: 8,
			expectedSizes: []int{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chunks := splitIntoChunks(tt.input)

			// Проверка количества чанков (всегда 8)
			assert.Len(t, chunks, CHUNKS, "Должно быть 8 чанков")

			// Проверка размеров
			var sizes []int
			for _, chunk := range chunks {
				sizes = append(sizes, len(chunk))
			}
			assert.Equal(t, tt.expectedSizes, sizes, "Неверные размеры чанков")

			// Проверка сохранения данных
			total := 0
			for _, chunk := range chunks {
				total += len(chunk)
			}
			assert.Equal(t, len(tt.input), total, "Потерялись элементы")
		})
	}
}
func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Максимум в последней части",
			input:    []int{1, 2, 3, 10, 5, 6, 8, 9, 4, 5, 9},
			expected: 10,
		},
		{
			name:     "Все элементы одинаковые",
			input:    []int{5, 5, 5, 5, 5, 5, 5, 5},
			expected: 5,
		},
		{
			name:     "Пустой слайс",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
