package ds

import (
	"fmt"
	"slices"
)

type MinHeapType interface {
	~int | ~int64 | ~float64 | ~string
}

type MinHeap[T MinHeapType] struct {
	arr []T
}

func (h *MinHeap[T]) insert(i int, t T) error { 
	var nilT T
	len := len(h.arr)

	// empty or end
	if len == 0 || i >= len {
		h.arr = append(h.arr, t)
		return nil
	}

	val := h.arr[i]

	// bigger than current
	if t > val {
		h.arr = slices.Insert(h.arr, i, t)
		return nil
	}

	if val == nilT {
		h.arr[i] = t
		return nil
	}

	if val == t {
		return nil
	}

	if val > t {
		next := h.arr[i+1]
		if next < t {
			h.arr = slices.Insert(h.arr, i+1, t)
			return nil
		}

		next2 := h.arr[i+2]
		if next2 < t {
			h.arr = slices.Insert(h.arr, i+2, t)
			return nil
		}
	}

	return h.insert(i+3, t)
}

func (h *MinHeap[T]) Insert(t T) error {
	if len(h.arr) == 0 {
		h.arr = append(h.arr, t)
		return nil
	}

	return h.insert(0, t)
}

func (h MinHeap[T]) String() string {
	return fmt.Sprintf("%v", h.arr)
}