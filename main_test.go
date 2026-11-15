package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	t.Run("zero size", func(t *testing.T) {
		got := generateRandomElements(0)
		if len(got) != 0 {
			t.Fatalf("expected empty slice, got length %d", len(got))
		}
	})

	t.Run("negative size", func(t *testing.T) {
		got := generateRandomElements(-5)
		if len(got) != 0 {
			t.Fatalf("expected empty slice, got length %d", len(got))
		}
	})

	t.Run("positive size", func(t *testing.T) {
		const size = 100
		got := generateRandomElements(size)
		if len(got) != size {
			t.Fatalf("expected length %d, got %d", size, len(got))
		}
		for i, value := range got {
			if value <= 0 {
				t.Errorf("element at index %d must be positive, got %d", i, value)
			}
		}
	})
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "empty slice", data: []int{}, want: 0},
		{name: "nil slice", data: nil, want: 0},
		{name: "one element", data: []int{42}, want: 42},
		{name: "maximum at beginning", data: []int{9, 2, 3, 1}, want: 9},
		{name: "maximum in middle", data: []int{1, 10, 3, 4}, want: 10},
		{name: "maximum at end", data: []int{1, 2, 3, 11}, want: 11},
		{name: "negative values", data: []int{-8, -2, -10}, want: -2},
		{name: "duplicate maximum", data: []int{5, 9, 9, 2}, want: 9},
		{name: "all equal", data: []int{7, 7, 7}, want: 7},
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
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "empty slice", data: []int{}, want: 0},
		{name: "nil slice", data: nil, want: 0},
		{name: "one element", data: []int{42}, want: 42},
		{name: "fewer elements than chunks", data: []int{2, 8, 3}, want: 8},
		{name: "exactly eight elements", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "remainder is distributed", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 20, 10}, want: 20},
		{name: "maximum at end", data: []int{1, 2, 3, 4, 5, 6, 7, 100}, want: 100},
		{name: "negative values", data: []int{-8, -2, -10, -4, -5, -6, -7, -9, -3}, want: -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	for size := 1; size <= 100; size++ {
		data := generateRandomElements(size)
		if got, want := maxChunks(data), maximum(data); got != want {
			t.Fatalf("size %d: maxChunks = %d, maximum = %d", size, got, want)
		}
	}
}
