package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	t.Run("zero size", func(t *testing.T) {
		elements := generateRandomElements(0)
		if len(elements) != 0 {
			t.Errorf("generateRandomElements(0) returned slice of length %d, want empty slice", len(elements))
		}
	})

	t.Run("negative size", func(t *testing.T) {
		elements := generateRandomElements(-10)
		if len(elements) != 0 {
			t.Errorf("generateRandomElements(-10) returned slice of length %d, want empty slice", len(elements))
		}
	})

	t.Run("single element", func(t *testing.T) {
		elements := generateRandomElements(1)
		if len(elements) != 1 {
			t.Fatalf("generateRandomElements(1) returned slice of length %d, want 1", len(elements))
		}

		if elements[0] < 1 {
			t.Errorf("generateRandomElements returned non positive element: %d", elements[0])
		}
	})

	t.Run("requested size and positive values", func(t *testing.T) {
		size := 1000
		elements := generateRandomElements(size)
		if len(elements) != size {
			t.Fatalf("generateRandomElements(%d) returned slice of length %d, want %d", size, len(elements), size)
		}

		for _, element := range elements {
			if element < 1 {
				t.Fatalf("generateRandomElements returned non-positive element: %d", element)
			}
		}
	})

	t.Run("random values", func(t *testing.T) {
		elements := generateRandomElements(1000)
		allSame := true

		for i := 1; i < len(elements); i++ {
			if elements[i] != elements[0] {
				allSame = false
				break
			}
		}

		if allSame {
			t.Errorf("all 1000 elements are equal (%d), generation is not random", elements[0])
		}
	})
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"nil slice", nil, 0},
		{"single element", []int{42}, 42},
		{"maximum at start", []int{100, 1, 2, 3}, 100},
		{"maximum at end", []int{1, 2, 3, 100}, 100},
		{"duplicate values", []int{7, 7, 7, 7}, 7},
		{"negative values", []int{-5, -1, -10}, -1},
		{"mixed values", []int{0, -3, 8, 2, -9, 8}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		if got := maxChunks(nil); got != 0 {
			t.Errorf("maxChunks(nil) = %d, want 0", got)
		}
	})

	t.Run("matches maximum", func(t *testing.T) {
		for _, size := range []int{1, 7, 8, 9, 100, 1001} {
			data := generateRandomElements(size)
			want := maximum(data)

			if got := maxChunks(data); got != want {
				t.Errorf("maxChunks on slice of length %d = %d, want %d", size, got, want)
			}
		}
	})
}
