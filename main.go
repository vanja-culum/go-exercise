package main

import (
	"fmt"
	"go-exercise/ds"
)

func main() {
	h := ds.MinHeap[int]{}

	h.Insert(5)
	h.Insert(10)
	h.Insert(30)

	h.Insert(12)
	h.Insert(21)
	fmt.Println(h)

}