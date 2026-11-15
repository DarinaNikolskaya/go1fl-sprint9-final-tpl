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

// generateRandomElements creates a slice of random positive integers.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]int, size)
	for i := range data {
		data[i] = rng.Intn(1_000_000) + 1
	}

	return data
}

// maximum returns the largest element in data. It returns 0 for an empty slice.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, value := range data[1:] {
		if value > max {
			max = value
		}
	}

	return max
}

// maxChunks finds the maximum value by processing up to eight chunks concurrently.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkMaxima := make([]int, CHUNKS)
	var wg sync.WaitGroup

	baseChunkSize := len(data) / CHUNKS
	remainder := len(data) % CHUNKS
	start := 0

	for i := 0; i < CHUNKS; i++ {
		chunkSize := baseChunkSize
		if i < remainder {
			chunkSize++
		}
		end := start + chunkSize

		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			chunkMaxima[index] = maximum(chunk)
		}(i, data[start:end])

		start = end
	}

	wg.Wait()
	return maximum(chunkMaxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
