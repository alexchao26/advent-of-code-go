package heap_test

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/data-structures/heap"
)

type mockNode int

func (n mockNode) Value() int {
	return int(n)
}

func TestMinHeap(t *testing.T) {
	h := heap.NewMinHeap()
	h.Add(mockNode(5))
	h.Add(mockNode(93))

	if h.Front().(mockNode).Value() != 5 {
		t.Errorf("After adding 5, h.Front().Value() = %d, want 5", h.Front().Value())
	}
	if h.Front().(mockNode).Value() != 5 {
		t.Errorf("After adding 5 & 93, h.Front() = %d, want 5", h.Front().Value())
	}

	// Add a bunch of nodes, make sure they are removed in order
	h.Add(mockNode(10))
	h.Add(mockNode(2))
	h.Add(mockNode(1))

	if h.Front().(mockNode).Value() != 1 {
		t.Errorf("After adding 5, 93, 10, 2 & 1, h.Front() = %d, want 1", h.Front().Value())
	}

	h.Add(mockNode(3))
	h.Add(mockNode(4))
	h.Add(mockNode(123))
	h.Add(mockNode(32))
	h.Add(mockNode(-15))

	// Ensure removing nodes returns in ascending order
	for _, want := range []int{-15, 1, 2, 3, 4, 5, 10, 32, 93, 123} {
		if got := h.Remove(); got.Value() != want {
			t.Errorf("h.Remove().Value() = %d, want %d", got.Value(), want)
		}
	}
}

func TestMaxHeap(t *testing.T) {
	h := heap.NewMaxHeap()
	h.Add(mockNode(5))
	h.Add(mockNode(93))

	if h.Front().Value() != 93 {
		t.Errorf("After adding 93, h.Front().Value() = %d, want 93", h.Front().Value())
	}
	if h.Front().Value() != 93 {
		t.Errorf("After adding 93 & 5, h.Front().Value() = %d, want 93", h.Front().Value())
	}

	// Add a bunch of nodes, make sure they are removed in order
	h.Add(mockNode(10))
	h.Add(mockNode(2))
	h.Add(mockNode(1))
	h.Add(mockNode(3))
	h.Add(mockNode(4))
	h.Add(mockNode(123))
	if h.Front().(mockNode).Value() != 123 {
		t.Errorf("After adding 5, 93, 10, 2, 1, 3, 4 & 123, h.Front() = %d, want 123", h.Front().Value())
	}

	h.Add(mockNode(32))
	h.Add(mockNode(-15))

	// Ensure removing returns in descending order
	for _, want := range []int{123, 93, 32, 10, 5, 4, 3, 2, 1, -15} {
		if got := h.Remove(); got.Value() != want {
			t.Errorf("h.Remove().Value() = %d, want %d", got.Value(), want)
		}
	}
}
