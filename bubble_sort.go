package main

import "fmt"

// BubbleSort implements the bubble sort algorithm
// It sorts the input slice in ascending order
func BubbleSort(arr []int) []int {
	if arr == nil || len(arr) <= 1 {
		return arr
	}

	// Make a copy of the array to avoid modifying the original
	sorted := make([]int, len(arr))
	copy(sorted, arr)

	n := len(sorted)
	for i := 0; i < n-1; i++ {
		// Flag to optimize - if no swaps occur, array is sorted
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				// Swap elements
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
				swapped = true
			}
		}

		// If no swaps occurred in this pass, array is already sorted
		if !swapped {
			break
		}
	}

	return sorted
}

// BubbleSortInPlace sorts the input slice in place using bubble sort
func BubbleSortInPlace(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}

	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func main() {
	// Example usage
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", numbers)

	// Using BubbleSort (returns new slice)
	sorted := BubbleSort(numbers)
	fmt.Println("Sorted array (new slice):", sorted)
	fmt.Println("Original array after sorting:", numbers)

	// Using BubbleSortInPlace (modifies original)
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)
	BubbleSortInPlace(numbersCopy)
	fmt.Println("Sorted array (in place):", numbersCopy)
}
