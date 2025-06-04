package main

import (
	"go-exercise/ds"
)

func main() {
	t := ds.BST[int]{}

	t.Insert(20)
	t.Insert(10)
	t.Insert(7)
	t.Insert(12)
	t.Insert(30)
	t.Insert(25)
	t.Insert(40)
	t.Insert(2)
	t.Insert(8)
	t.Insert(11)
	t.Insert(15)

	t.Remove(40)

}