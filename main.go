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

func generateRandomElements(size int) ([]int, error) {
	if size < 0 {
		return nil, fmt.Errorf("длина слайса должна быть неотрицательной")
	}
	if size == 0 {
		return nil, fmt.Errorf("длина слайса должна быть больше 0")
	}

	numbers := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		numbers[i] = r.Intn(100)
	}
	return numbers, nil
}

func maximum(data []int) (int, error) {
	// ваш код здесь
	if data == nil || len(data) == 0 {
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
	if len(data) == 0 {
		return nil
	}
	var result [][]int
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
		result = append(result, data[start:end])
	}
	return result
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	subSlices := splitIntoChunks(data)
	results := make(chan int, CHUNKS)
	var wg sync.WaitGroup

	for i, chunk := range subSlices {
		wg.Add(1)
		go func(idx int, chunk []int) {
			defer wg.Done()
			if len(chunk) == 0 {
				results <- 0
				return
			}

			currentMax := chunk[0]
			for _, num := range chunk {
				if num > currentMax {
					currentMax = num
				}
			}
			results <- currentMax
		}(i+1, chunk)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	maxResult := 0
	for num := range results {
		if num > maxResult {
			maxResult = num
		}
	}
	return maxResult
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	generateSlice, err := generateRandomElements(SIZE)
	if err != nil {
		fmt.Printf("Ошибка генерации данных: %v\n", err)
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
