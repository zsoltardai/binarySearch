package binarySearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {

	var sequence = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	t.Run("withExistingValue", func(t *testing.T) {
		var item int = 8
		var result int = BinarySearch(sequence, item)
		if result != 7 {
			t.Fatalf("Got: %v\n", result)
		}
	})

	t.Run("withNonExistingValue", func(t *testing.T) {
		var item int = 11
		var result int = BinarySearch(sequence, item)
		if result != -1 {
			t.Fatalf("Got: %v\n", result)
		}
	})
}
