package main

import (
	"fmt"
	"math"
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
		return []int{}
	}

	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Intn(math.MaxInt) + 1
	}

	return elements
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxElement := data[0]
	for _, element := range data {
		if element > maxElement {
			maxElement = element
		}
	}

	return maxElement
}

func maxChunks(data []int) int {
	chunkSize := len(data) / CHUNKS
	chunkMaxima := make([]int, CHUNKS)

	var wg sync.WaitGroup
	for i := range CHUNKS {
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize

		if i == CHUNKS-1 {
			endIndex = len(data)
		}

		wg.Add(1)
		go func(chunk []int, index int) {
			defer wg.Done()
			chunkMaxima[index] = maximum(chunk)
		}(data[startIndex:endIndex], i)
	}

	wg.Wait()

	return maximum(chunkMaxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startTime := time.Now()
	maxVal := maximum(data)
	elapsed := time.Since(startTime).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxVal, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	startTime = time.Now()
	maxVal = maxChunks(data)
	elapsed = time.Since(startTime).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxVal, elapsed)
}
