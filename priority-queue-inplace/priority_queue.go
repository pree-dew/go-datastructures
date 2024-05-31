package pqinplace

// This package provides a priority queue implementation using a slice in place of a heap.
// The priority queue can be either a min heap or a max heap.
// The priority queue is implemented using a slice and the heap operations are done in place.

func ConvertToMaxHeap(arr []int) {
	for i := 1; i < len(arr); i++ {
		index := i
		parentIndex := (index - 1) / 2

		for index > 0 {
			if arr[index] > arr[parentIndex] {
				arr[index], arr[parentIndex] = arr[parentIndex], arr[index]
				index = parentIndex
				parentIndex = (index - 1) / 2
			} else {
				break
			}
		}
	}
}

func ConvertToMinHeap(arr []int) {
	for i := 1; i < len(arr); i++ {
		index := i
		parentIndex := (index - 1) / 2

		for index > 0 {
			if arr[index] < arr[parentIndex] {
				arr[index], arr[parentIndex] = arr[parentIndex], arr[index]
				index = parentIndex
				parentIndex = (index - 1) / 2
			} else {
				break
			}
		}
	}
}

func PushToMaxHeap(arr *[]int, val int) {
	*arr = append(*arr, val)
	index := len(*arr) - 1
	parentIndex := (index - 1) / 2

	for index > 0 {
		if (*arr)[index] > (*arr)[parentIndex] {
			(*arr)[index], (*arr)[parentIndex] = (*arr)[parentIndex], (*arr)[index]
			index = parentIndex
			parentIndex = (index - 1) / 2
		} else {
			break
		}
	}
}

func PushToMinHeap(arr *[]int, val int) {
	*arr = append(*arr, val)
	index := len(*arr) - 1
	parentIndex := (index - 1) / 2

	for index > 0 {
		if (*arr)[index] < (*arr)[parentIndex] {
			(*arr)[index], (*arr)[parentIndex] = (*arr)[parentIndex], (*arr)[index]
			index = parentIndex
			parentIndex = (index - 1) / 2
		} else {
			break
		}
	}
}

func maxHeapBubbleDown(arr []int, index, size int) {
	for index < size {
		left := 2*index + 1
		right := 2*index + 2

		if left >= size {
			break
		}

		larger := left
		if right < size && arr[left] < arr[right] {
			larger = right
		}

		if arr[index] < arr[larger] {
			arr[index], arr[larger] = arr[larger], arr[index]
			index = larger
		} else {
			break
		}
	}
}

func SortMaxHeap(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapBubbleDown(arr, 0, i)
	}
}

func minHeapBubbleDown(arr []int, index, size int) {
	for index < size {
		left := 2*index + 1
		right := 2*index + 2

		if left >= size {
			break
		}

		smaller := left
		if right < size && arr[left] > arr[right] {
			smaller = right
		}

		if arr[index] > arr[smaller] {
			arr[index], arr[smaller] = arr[smaller], arr[index]
			index = smaller
		} else {
			break
		}
	}
}

func SortMinHeap(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		minHeapBubbleDown(arr, 0, i)
	}
}
