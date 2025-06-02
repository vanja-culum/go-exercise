package ds

import (
	"fmt"
	"strings"
)

type stackNode[T any] struct {
	val T
	next *stackNode[T]
	prev *stackNode[T]
}

type Stack[T any] struct {
	head *stackNode[T]
	tail *stackNode[T]
}

func (s *Stack[T]) Push(t T)  {
	newNode := &stackNode[T]{
		val: t,
	}

	if s.head == nil {
		s.head = newNode
		s.tail = newNode
		return
	}
	
	s.tail.next = newNode
	newNode.prev = s.tail
	s.tail = newNode
}

func (s *Stack[T]) Pop() (T, error) {
	if s.head == nil {
		var t T
		return t, fmt.Errorf("stack empty")
	}

	val := s.tail.val
	s.tail = s.tail.prev

	if s.tail == nil {
		s.head = nil
	} else {
		s.tail.next = nil
	}
	
	
	return val, nil
}

func (s *Stack[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	for curr := s.head; curr != nil ; curr = curr.next {
		sb.WriteString(fmt.Sprintf("%v", curr.val))

		if curr.next != nil {
				sb.WriteString(", ")
		}
	}

	sb.WriteString("]")

	return sb.String()

}

func (s *Stack[T]) Peek() (T, error) {
	var val T
	if s.head == nil {
		return val, fmt.Errorf("stack empty")
	}

	return s.tail.val, nil
}

func (s *Stack[T]) Clear() {
	s.head = nil
	s.tail = nil
}

func (s *Stack[T]) Size() int {
	if s.head == nil {
		return 0
	}

	i := 0
	curr := s.head
	for curr != nil {
		i++
		curr = curr.next
	}

	return i
}