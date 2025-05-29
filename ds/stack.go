package ds

import (
	"fmt"
	"strings"
)

type stackNode[T any] struct {
	val T
	next *stackNode[T]
}

type Stack[T any] struct {
	base *stackNode[T]
}

func (s *Stack[T]) Push(t T)  {
	newNode := &stackNode[T]{
		val: t,
	}

	if s.base == nil {
		s.base = newNode
		return
	}

	curr := s.base

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode	
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var t T
		return t, fmt.Errorf("stack empty")
	}


	curr := s.base
	prev := curr

	for curr.next != nil {
		prev = curr
		curr = curr.next
	}

	prev.next = nil

	return curr.val, nil
}

func (s *Stack[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	for curr := s.base; curr != nil ; curr = curr.next {
		sb.WriteString(fmt.Sprintf("%v", curr.val))

		if curr.next != nil {
				sb.WriteString(", ")
		}

	}

	sb.WriteString("]")

	return sb.String()

}

func (s *Stack[T]) Peek() (T, error) {
	empty := s.IsEmpty()
	var val T
	
	if empty {
		return val, fmt.Errorf("stack empty")
	}

	curr := s.base

	for curr.next != nil {
		curr = curr.next
	}

	return curr.val, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.base == nil
}

func (s *Stack[T]) Clear() {
	s.base = nil
}