package pq

type PriorityQueue struct {
	heap     []int
	heapType string
}

func NewPriorityQueue(heapType string) *PriorityQueue {
	return &PriorityQueue{
		heap:     []int{},
		heapType: heapType,
	}
}

func (pq *PriorityQueue) Len() int {
	return len(pq.heap)
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue) Swap(i, j int) bool {
	if i < 0 || j < 0 || i >= len(pq.heap) || j >= len(pq.heap) {
		return false
	}

	if pq.heapType == "min" {
		if pq.heap[i] < pq.heap[j] {
			pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
			return true
		}

		return false
	}

	if pq.heap[i] > pq.heap[j] {
		pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
		return true
	}

	return false
}

// bubbleUp is a helper function that do the process of upheapify in the heap
// in case of min heap, it compares the last element with its parent and
// swap them if the parent is greater than the last element
// in case of max heap, it compares the last element with its parent and
// swap them if the parent is smaller than the last element
func (pq *PriorityQueue) bubbleUp() {
	index := len(pq.heap) - 1
	parentIndex := (index - 1) / 2

	for index > 0 {
		if pq.Swap(index, parentIndex) {
			index = parentIndex
			parentIndex = (index - 1) / 2
		} else {
			break
		}
	}
}

func (pq *PriorityQueue) Push(value int) {
	pq.heap = append(pq.heap, value)
	pq.bubbleUp()
}

func (pq *PriorityQueue) getChildIndex(leftChildIndex, rightChildIndex int) int {
	if leftChildIndex < 0 || leftChildIndex >= pq.Len() {
		return -1
	}

	index := leftChildIndex
	if pq.heapType == "min" {
		if rightChildIndex < pq.Len() && pq.heap[leftChildIndex] > pq.heap[rightChildIndex] {
			index = rightChildIndex
		}

		return index
	}

	if rightChildIndex < pq.Len() && pq.heap[leftChildIndex] < pq.heap[rightChildIndex] {
		index = rightChildIndex
	}

	return index
}

// bubbleDown is a helper function that do the process of downheapify in the heap
// in case of min heap, it compares the first element with its children and
// swap them with the smallest child if the child is smaller than the parent
// in case of max heap, it compares the first element with its children and
// swap them with the largest child if the child is greater than the parent
func (pq *PriorityQueue) bubbleDown() {
	index := 0

	for index < pq.Len() {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2

		childIndex := pq.getChildIndex(leftChildIndex, rightChildIndex)
		if childIndex == -1 {
			break
		}

		if pq.Swap(index, childIndex) {
			index = childIndex
		} else {
			break
		}
	}
}

func (pq *PriorityQueue) Pop() int {
	if pq.IsEmpty() {
		return -1
	}

	top := pq.heap[0]
	pq.heap[0] = pq.heap[pq.Len()-1]
	pq.heap = pq.heap[:pq.Len()-1]
	pq.bubbleDown()

	return top
}
