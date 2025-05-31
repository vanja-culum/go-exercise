package ds

import (
	"fmt"
	"strings"
)

type queueNode[T any] struct {
	val T
	next *queueNode[T]
	prev *queueNode[T]
}

type Queue[T any] struct {
	head *queueNode[T]
	tail *queueNode[T]
}

func (q *Queue[T]) Enqueue(t T) {
	newNode := &queueNode[T]{
		val: t,
	}

	if q.head == nil || q.tail == nil {
		q.head = newNode
		q.tail = newNode

		return
	}

	newNode.prev = q.tail
	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue[T]) String() string {
	var sb strings.Builder

	sb.WriteString("[")

	if q.head == nil {
		sb.WriteString("]")
		return sb.String()
	}

	for curr := q.head ; curr != nil ; curr = curr.next {
		sb.WriteString(fmt.Sprintf("%v", curr.val))
		if curr.next != nil {
			sb.WriteString(", ")
		} 
	}

	sb.WriteString("]")

	return sb.String()
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.head == nil {
		var t T
		return t, fmt.Errorf("queue empty")
	}	

	val := q.head.val

	if q.head.next == nil {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
		q.head.prev = nil
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