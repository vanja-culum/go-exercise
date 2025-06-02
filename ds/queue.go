package ds

import (
	"fmt"
	"strings"
)

type queueNode[T any] struct {
	val T
	next *queueNode[T]
}

type Queue[T any] struct {
	head *queueNode[T]
	tail *queueNode[T]
}

func (q *Queue[T]) Enqueue(t T) {
	newNode := &queueNode[T]{
		val: t,
	}

	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		return
	}

	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue[T]) String() string {
	var sb strings.Builder

	sb.WriteString("[")

	for curr := q.head ; curr != nil ; curr = curr.next {
		sb.WriteString(fmt.Sprintf("%v", curr.val))
		if curr.next != nil {
			sb.WriteString(", ")
		} 
	}

	sb.WriteString("]")

	return sb.String()
}

func (q *Queue[T]) Clear() {
	q.head = nil
	q.tail = nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.head == nil {
		var t T
		return t, fmt.Errorf("queue empty")
	}	

	val := q.head.val
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return val, nil
}

func (q *Queue[T]) Size() int {
	curr := q.head
	i := 0

	for curr != nil {
		curr = curr.next
		i++
	}

	return i
}