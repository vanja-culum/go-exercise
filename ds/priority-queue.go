package ds

import (
	"fmt"
	"strings"
)


type pqNode[T any] struct {
	val T
	priority int
	prev *pqNode[T]
	next *pqNode[T]
}

type PriorityQueue[T any] struct {
	head *pqNode[T]
	tail *pqNode[T]
}

func (pq *PriorityQueue[T]) Enqueue(val T, priority int) {
	newNode := &pqNode[T]{
		val: val,
		priority: priority,
	}

	if pq.head == nil {
		pq.head = newNode
		pq.tail = newNode
		return
	}

	for curr := pq.head; curr != nil; curr = curr.next {
		if priority > curr.priority {
			curr.prev = newNode
			newNode.next = curr
			if curr == pq.head {
				pq.head = newNode
			}
			return
		}

		if curr.next == nil {
			curr.next = newNode
			newNode.prev = curr
			pq.tail = newNode
			return
		}

		if curr.priority == priority && curr.next.priority < priority {
			tmpNext := curr.next
			curr.next = newNode
			newNode.prev = curr
			newNode.next = tmpNext
			if curr == pq.tail {
				pq.tail = newNode
			}
			return
		}
	}
}

func (pq *PriorityQueue[T]) String() string {
	if pq.head == nil {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")

	curr := pq.head

	for curr != nil {
		sb.WriteString(fmt.Sprintf("%v", curr.val))
		if curr.next != nil {
			sb.WriteString(", ")
		}
		curr = curr.next
	}

	sb.WriteString("]")

	return sb.String()
}

func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	if pq.head == nil {
		var t T
		return t, fmt.Errorf("queue empty")
	}

	tmpNext, tmpVal := pq.head.next, pq.head.val
	pq.head = tmpNext
	return tmpVal, nil
}

