package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Random order",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:     "With duplicates",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			expected: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name:     "Negative numbers",
			input:    []int{-5, 3, -2, 8, -10, 0, 1},
			expected: []int{-10, -5, -2, 0, 1, 3, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BubbleSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBubbleSortInPlace(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Random order",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying test data
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			BubbleSortInPlace(input)

			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("BubbleSortInPlace() = %v, want %v", input, tt.expected)
			}
		})
	}
}

func TestBubbleSortNil(t *testing.T) {
	result := BubbleSort(nil)
	if result != nil {
		t.Errorf("BubbleSort(nil) = %v, want nil", result)
	}

	// Test in-place with nil - should not panic
	BubbleSortInPlace(nil)
}

func BenchmarkBubbleSort(b *testing.B) {
	data := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 42}
	for i := 0; i < b.N; i++ {
		BubbleSort(data)
	}
}

func BenchmarkBubbleSortInPlace(b *testing.B) {
	data := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 42}
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration
		testData := make([]int, len(data))
		copy(testData, data)
		BubbleSortInPlace(testData)
	}
}
