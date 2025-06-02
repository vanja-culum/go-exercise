package main

import (
	"fmt"
	"go-exercise/ds"
)

func main() {
	s := ds.Stack[int]{}
	q := ds.Queue[int]{}
	s.Push(12)
	s.Push(14)
	s.Push(16)

	q.Enqueue(5)
	q.Enqueue(10)
	q.Enqueue(15)
	fmt.Println(s.String())

	fmt.Println(q.String())

	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	fmt.Println(s.String())
	q.Dequeue()
	q.Dequeue()

fmt.Println(q.String())
}