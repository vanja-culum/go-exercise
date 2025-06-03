package main

import (
	"fmt"
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

	arr1 := t.InOrder()
	arr2 := t.PreOrder()
	arr3 := t.PostOrder()
	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)
	fmt.Println("arr3", arr3)
	arr4 := t.BFS()

	fmt.Println("arr4", arr4)

}