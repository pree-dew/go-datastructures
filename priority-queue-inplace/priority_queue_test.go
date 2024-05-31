package pqinplace

import (
	"testing"
)

func TestConvertToMaxHeap(t *testing.T) {
	arr := []int{10, 5, 3, 6, 2, 4, 8}
	ConvertToMaxHeap(arr)

	expected := []int{10, 6, 8, 5, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], arr[i])
		}
	}
}

func TestConvertToMinHeap(t *testing.T) {
	arr := []int{10, 5, 3, 6, 2, 4, 8}
	ConvertToMinHeap(arr)

	expected := []int{2, 3, 4, 10, 6, 5, 8}
	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], arr[i])
		}
	}
}

func TestSortMaxHeap(t *testing.T) {
	arr := []int{10, 5, 3, 6, 2, 4, 8}
	ConvertToMaxHeap(arr)

	SortMaxHeap(arr)

	expected := []int{2, 3, 4, 5, 6, 8, 10}
	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], arr[i])
		}
	}
}

func TestSortMinHeap(t *testing.T) {
	arr := []int{10, 5, 3, 6, 2, 4, 8}
	ConvertToMinHeap(arr)

	SortMinHeap(arr)

	expected := []int{10, 8, 6, 5, 4, 3, 2}
	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], arr[i])
		}
	}
}

func TestPushToMaxHeap(t *testing.T) {
	arr := []int{10, 5, 3, 6, 2, 4, 8}
	ConvertToMaxHeap(arr)
	PushToMaxHeap(&arr, 7)

	expected := []int{10, 7, 8, 6, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], arr[i])
		}
	}
}

func TestPushToMinHeap(t *testing.T) {
	arr := []int{10, 5, 3, 6, 2, 4, 8}
	ConvertToMinHeap(arr)
	PushToMinHeap(&arr, 1)

	expected := []int{1, 2, 4, 3, 6, 5, 8, 10}
	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], arr[i])
		}
	}
}
