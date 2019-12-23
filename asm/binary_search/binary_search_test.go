package binary_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if result := BinarySearch(data, 8); result != true {
		t.Errorf("Expected true value for binary search.")
	}

	if result := BinarySearch(data, 1); result != true {
		t.Errorf("Expected true value for binary search.")
	}

	if result := BinarySearch(data, 10); result != true {
		t.Errorf("Expected true value for binary search.")
	}

	if result := BinarySearch(data, 12); result != false {
		t.Errorf("Expected false value for binary search.")
	}
}
