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
	if size < 0 {
		return nil
	}
	if size == 0 {
		return nil
	}

	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Int()
	}
	return numbers
}

func maximum(data []int) (int, error) {
	// ваш код здесь
	if len(data) == 0 {
		return 0, fmt.Errorf("слайс пустой или равен nil")
	}

	max := data[0]
	for _, num := range data {
		if num > max {
			max = num
		}
	}
	return max, nil
}

func splitIntoChunks(data []int) [][]int {
	chunks := make([][]int, CHUNKS) // Всегда создаем 8 чанков

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
	var mutex sync.Mutex
	maxResult := 0
	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		if start >= len(data) {
			break
		}

		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}

		wg.Add(1)
		go func(chunk []int) {
			defer wg.Done()

			currentMax, err := maximum(chunk)
			if err != nil {
				return
			}

			mutex.Lock()
			if currentMax > maxResult {
				maxResult = currentMax
			}
			mutex.Unlock()
		}(data[start:end])
	}

	wg.Wait()
	return maxResult
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
	oneStreamMax, err := maximum(generateSlice)
	if err != nil {
		fmt.Printf("Ошибка поиска максимума: %v\n", err)
		return
	}
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", oneStreamMax, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max := maxChunks(generateSlice)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
