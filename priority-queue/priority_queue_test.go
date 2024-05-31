package pq

import "testing"

func TestMinQueuePush(t *testing.T) {
	pq := NewPriorityQueue("min")
	pq.Push(3)
	pq.Push(2)
	pq.Push(1)

	if pq.heap[0] != 1 {
		t.Errorf("Expected 1, got %d", pq.heap[0])
	}
}

func TestMaxQueuePush(t *testing.T) {
	pq := NewPriorityQueue("max")
	pq.Push(1)
	pq.Push(2)
	pq.Push(3)

	if pq.heap[0] != 3 {
		t.Errorf("Expected 3, got %d", pq.heap[0])
	}
}

func TestMinQueuePop(t *testing.T) {
	pq := NewPriorityQueue("min")
	pq.Push(4)
	pq.Push(3)
	pq.Push(2)
	pq.Push(1)

	if pq.Pop() != 1 {
		t.Errorf("Expected 1, got %d", pq.Pop())
	}

	pq.Pop()

	if pq.heap[0] != 3 {
		t.Errorf("Expected 3, got %d", pq.heap[0])
	}
}

func TestMaxQueuePop(t *testing.T) {
	pq := NewPriorityQueue("max")
	pq.Push(1)
	pq.Push(2)
	pq.Push(3)
	pq.Push(4)

	if pq.Pop() != 4 {
		t.Errorf("Expected 4, got %d", pq.Pop())
	}

	pq.Pop()

	if pq.heap[0] != 2 {
		t.Errorf("Expected 2, got %d", pq.heap[0])
	}
}

func TestPopEmptyQueue(t *testing.T) {
	pq := NewPriorityQueue("min")

	if pq.Pop() != -1 {
		t.Errorf("Expected -1, got %d", pq.Pop())
	}
}
