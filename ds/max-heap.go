package ds

import (
	"fmt"
)

// Ordered types that support comparison
type Ordered interface {
	~int | ~int64 | ~float64 | ~string
}

type MaxHeap[T Ordered] struct {
	arr []T
}

// Insert adds a new value and maintains the max-heap property
func (h *MaxHeap[T]) Insert(value T) {
	h.arr = append(h.arr, value)
	h.heapifyUp(len(h.arr) - 1)
}

// ExtractMax removes and returns the largest value (the root)
func (h *MaxHeap[T]) ExtractMax() (T, error) {
	var zero T
	if len(h.arr) == 0 {
		return zero, fmt.Errorf("heap is empty")
	}

	max := h.arr[0]
	last := h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	if len(h.arr) > 0 {
		h.arr[0] = last
		h.heapifyDown(0)
	}

	return max, nil
}

// Peek returns the maximum value without removing it
func (h *MaxHeap[T]) Peek() (T, error) {
	var zero T
	if len(h.arr) == 0 {
		return zero, fmt.Errorf("heap is empty")
	}
	return h.arr[0], nil
}

// heapifyUp restores the heap by bubbling the value up
func (h *MaxHeap[T]) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.arr[i] <= h.arr[parent] {
			break
		}
		h.arr[i], h.arr[parent] = h.arr[parent], h.arr[i]
		i = parent
	}
}

// heapifyDown restores the heap by bubbling the value down
func (h *MaxHeap[T]) heapifyDown(i int) {
	n := len(h.arr)
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i

		if left < n && h.arr[left] > h.arr[largest] {
			largest = left
		}
		if right < n && h.arr[right] > h.arr[largest] {
			largest = right
		}
		if largest == i {
			break
		}
		h.arr[i], h.arr[largest] = h.arr[largest], h.arr[i]
		i = largest
	}
}

// String returns a string representation of the heap
func (h MaxHeap[T]) String() string {
	return fmt.Sprintf("%v", h.arr)
}