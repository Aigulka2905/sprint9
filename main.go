package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Int()
	}
	return numbers
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, num := range data {
		if num > max {
			max = num
		}
	}
	return max
}

func splitIntoChunks(data []int) [][]int {
	chunks := make([][]int, CHUNKS)
	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if start >= len(data) {
			chunks[i] = []int{}
			continue
		}

		if end > len(data) {
			end = len(data)
		}

		chunks[i] = data[start:end]
	}

	return chunks
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup
	maxResults := make([]int, CHUNKS)
	used := make([]bool, CHUNKS)

	chunks := splitIntoChunks(data)

	for i, chunk := range chunks {
		wg.Add(1)
		go func(chunk []int, index int) {
			defer wg.Done()

			if len(chunk) == 0 {
				return
			}

			currentMax := maximum(chunk)
			maxResults[index] = currentMax
			used[index] = true
		}(chunk, i)
	}

	wg.Wait()

	globalMax := 0
	hasValidResult := false
	for i, max := range maxResults {
		if used[i] {
			if !hasValidResult || max > globalMax {
				globalMax = max
				hasValidResult = true
			}
		}
	}

	if !hasValidResult {
		return 0
	}
	return globalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	generateSlice := generateRandomElements(SIZE)
	if generateSlice == nil {
		fmt.Println("Ошибка генерации данных: получен nil-слайс")
		return
	}
	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	oneStreamMax := maximum(generateSlice)

	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", oneStreamMax, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max := maxChunks(generateSlice)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
