package main

// Пишите тесты в этом файле
import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Valid positive size",
			size:    10,
			wantErr: false,
		},
		{
			name:    "Zero size",
			size:    0,
			wantErr: true,
			errMsg:  "длина слайса должна быть больше 0",
		},
		{
			name:    "Negative size",
			size:    -1,
			wantErr: true,
			errMsg:  "длина слайса должна быть неотрицательной",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := generateRandomElements(tt.size)

			if (err != nil) != tt.wantErr {
				t.Errorf("For size %d: error = %v, wantErr %v", tt.size, err, tt.wantErr)
				return
			}

			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("For size %d: error message = %q, want %q",
					tt.size, err.Error(), tt.errMsg)
			}
		})
	}
}

func TestFindMax(t *testing.T) {
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
		{
			name:     "Negative numbers",
			numbers:  []int{-10, -5, -20, -1},
			expected: -1,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maximum(tt.numbers)

			// Проверка ошибок
			if (err != nil) != tt.wantErr {
				t.Errorf("maximum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("maximum() error message = %v, want %v", err.Error(), tt.errMsg)
			}

			// Проверка результата
			if !tt.wantErr && got != tt.expected {
				t.Errorf("maximum() = %v, want %v", got, tt.expected)
			}
		})
	}
}
func TestSplitIntoChunks(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output [][]int
	}{
		{
			name:   "Разделение на 8 частей",
			input:  []int{1, 2, 3, 4, 5, 6, 7, 5, 9, 4, 8, 3, 2, 5, 4, 3},
			output: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 5}, {9, 4}, {8, 3}, {2, 5}, {4, 3}},
		},
		{
			name:   "Разделение с остатком",
			input:  []int{1, 2, 3, 3, 4, 5, 6, 7, 5, 9, 4, 8, 3, 2, 5, 4, 3},
			output: [][]int{{1, 2, 3}, {3, 4, 5}, {6, 7, 5}, {9, 4, 8}, {3, 2, 5}, {4, 3}},
		},
		{
			name:   "Пустой слайс",
			input:  []int{},
			output: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chunks := splitIntoChunks(tt.input)
			if !equalSlices(chunks, tt.output) {
				t.Errorf("splitIntoChunks() = %v, ожидается %v", chunks, tt.output)
			}
		})
	}
}

// Тест для SplitSlice (основная функция)
func TestSplitSlice(t *testing.T) {
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
			if result != tt.expected {
				t.Errorf("SplitSlice() = %d, ожидается %d", result, tt.expected)
			}
		})
	}
}

// Вспомогательная функция для сравнения слайсов
func equalSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
