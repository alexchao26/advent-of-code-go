package structures

import (
	"testing"
)

type mockNode int

func (n mockNode) Value() int {
	return int(n)
}

func TestMinHeap(t *testing.T) {
	h := NewMinHeap()
	h.Add(mockNode(5))
	h.Add(mockNode(93))

	if h.nodes[0].Value() != 5 {
		t.Errorf("After adding 5, h.nodes[0].Value() = %d, want 5", h.nodes[0].Value())
	}
	if h.nodes[1].Value() != 93 {
		t.Errorf("After adding 93, h.nodes[1].Value() = %d, want 93", h.nodes[1].Value())
	}

	// Add a bunch of nodes, make sure they are removed in order
	h.Add(mockNode(10))
	h.Add(mockNode(2))
	h.Add(mockNode(1))
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
	h := NewMaxHeap()
	h.Add(mockNode(5))
	h.Add(mockNode(93))

	if h.nodes[0].Value() != 93 {
		t.Errorf("After adding 93, h.nodes[0].Value() = %d, want 93", h.nodes[1].Value())
	}
	if h.nodes[1].Value() != 5 {
		t.Errorf("After adding 5, h.nodes[1].Value() = %d, want 5", h.nodes[0].Value())
	}

	// Add a bunch of nodes, make sure they are removed in order
	h.Add(mockNode(10))
	h.Add(mockNode(2))
	h.Add(mockNode(1))
	h.Add(mockNode(3))
	h.Add(mockNode(4))
	h.Add(mockNode(123))
	h.Add(mockNode(32))
	h.Add(mockNode(-15))

	// Ensure removing returns in descending order
	for _, want := range []int{123, 93, 32, 10, 5, 4, 3, 2, 1, -15} {
		if got := h.Remove(); got.Value() != want {
			t.Errorf("h.Remove().Value() = %d, want %d", got.Value(), want)
		}
	}
}
