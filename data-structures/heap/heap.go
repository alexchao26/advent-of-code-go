package heap

// MinHeap is an implementation of a min heap
type MinHeap struct {
	heap
}

// NewMinHeap initializes a heap with a closerToRootFunction that simply
// returns true if the first arg is smaller than the second
func NewMinHeap() *MinHeap {
	nestedHeap := heap{
		closerToRoot: func(val1, val2 int) bool {
			return val1 < val2
		},
	}
	return &MinHeap{nestedHeap}
}

// MaxHeap is an implementation of max heap
type MaxHeap struct {
	heap
}

// NewMaxHeap initializes a heap with a closerToRootFunction that simply
// returns true if the first arg is larger than the second
func NewMaxHeap() *MaxHeap {
	nestedHeap := heap{
		closerToRoot: func(val1, val2 int) bool {
			return val1 > val2
		},
	}
	return &MaxHeap{nestedHeap}
}

// heap contains a slice of heapNodes
// A heap can be represented as an array/slice with no gaps because
// calculating the indices of two children or the parent is simple
// from any given index
type heap struct {
	nodes        []HeapNode
	closerToRoot func(val1, val2 int) bool
}

// HeapNode is an interface making the type for a Min/MaxHeap node flexible
// nodes must be be able to state their value to be sorted by
type HeapNode interface {
	Value() int
}

// Front returns the first node in the heap, nil if the heap is empty
func (h *heap) Front() HeapNode {
	if len(h.nodes) == 0 {
		return nil
	}
	return h.nodes[0]
}

// Add appends a new node onto the heap and heapifies it
// to ensure correct ordering
func (h *heap) Add(newNode HeapNode) {
	h.nodes = append(h.nodes, newNode)
	h.heapifyFromEnd()
}

// Remove returns the node at the root, i.e. the minimum value node
func (h *heap) Remove() HeapNode {
	if len(h.nodes) == 0 {
		return nil
	}

	rootNode := h.nodes[0]

	// move last node to start & reduce length by one
	h.nodes[0] = h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]

	// heapify the heap from the start to sort the minimum value into the 0 index
	h.heapifyFromStart()

	return rootNode
}

func (h *heap) Length() int {
	return len(h.nodes)
}

func (h *heap) swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

// heapify from end expects an unordered value in the last index, it will compare
// it to its parent index and swapped if applicable, and repeated until the heap
// is valid
func (h *heap) heapifyFromEnd() {
	currentIndex := len(h.nodes) - 1
	for currentIndex > 0 {
		parentIndex := (currentIndex - 1) / 2
		parentNode := h.nodes[parentIndex]
		if h.closerToRoot(h.nodes[currentIndex].Value(), parentNode.Value()) {
			h.swap(parentIndex, currentIndex)
			currentIndex = parentIndex
		} else {
			break
		}
	}
}

// heapify from start expects an unordered value in the heap in index zero,
// that node's value is compared to its children, and swaps are made as needed
// until the heap is valid
func (h *heap) heapifyFromStart() {
	currentIndex := 0

	for {
		// find smaller of two children
		smallerChildIndex := currentIndex
		for i := 1; i <= 2; i++ {
			childIndex := currentIndex*2 + i
			// if a child value is closer to the root than the current node,
			// store it's index
			if childIndex < len(h.nodes) &&
				h.closerToRoot(h.nodes[childIndex].Value(), h.nodes[smallerChildIndex].Value()) {
				smallerChildIndex = childIndex
			}
		}

		// if smallerChildIndex was not reassigned, no swap is needed, return out
		if smallerChildIndex == currentIndex {
			return
		}

		// otherwise swap & update currentIndex to keep checking on next loop
		h.swap(smallerChildIndex, currentIndex)
		currentIndex = smallerChildIndex
	}
}
