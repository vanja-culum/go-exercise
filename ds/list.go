package ds

import (
	"fmt"
	"strings"
)

type node[T comparable] struct {
    val  T
	prev *node[T]
    next *node[T]
}

type List[T comparable] struct {
    head *node[T]
	tail *node[T]
}

func (lst *List[T]) Append(t T) {
    if lst.IsEmpty() {
		newNode := &node[T]{
			val: t,
		}
		
        lst.head = newNode
		lst.tail = newNode
        return
    }

	newNode := &node[T]{
        val: t,
		prev: lst.tail,
    }


	lst.tail.next = newNode
	lst.tail = newNode

}

func (lst *List[T]) Prepend(t T) {
    newNode := &node[T]{
        val: t,
        next: lst.head,
    }

    lst.head = newNode
}

func (lst *List[T]) IndexOf(t T) (int, error) {
    i := 0

    for el := lst.head; el != nil; el = el.next {
        if el.val == t {
            return i, nil
        }

        i++
    }

    return -1, fmt.Errorf("index of value %v not found", t)
}

func (lst *List[T]) IsEmpty() bool {
	return lst.Len() == 0
}

func (lst *List[T]) Reverse() {
	lst.tail = lst.head
	curr := lst.head
	var prev *node[T]
	var nextTmp *node[T]
	for curr != nil {
		nextTmp, curr.next = curr.next, prev
		prev, curr = curr, nextTmp
	}

	lst.head = prev

}

func (lst *List[T]) Clear() {
	lst.head = nil
}

func (lst *List[T]) ToSlice() []T {
	slc := []T{}

	empty := lst.IsEmpty()

	fmt.Println(empty)
	if empty  {
		return slc
	}



	curr := lst.head

	for curr != nil {
		slc = append(slc, curr.val)
		curr = curr.next
	}

	return slc
}

func (lst *List[T]) Get(index int) (T, error) {
	len := lst.Len()

	var t T

	if index < 0 || index > len {
		return t, fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		return lst.head.val, nil
	}

	curr := lst.head
	i := 0

	for curr != nil {
		if i == index {
			return curr.val, nil
		}

		curr = curr.next
		i++
	}

	return t, fmt.Errorf("index out of bounds")

}

func (lst *List[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for curr := lst.head; curr != nil; curr = curr.next {
		sb.WriteString(fmt.Sprintf("%v", curr.val))
		if curr.next != nil {
			sb.WriteString(", ")
		}
	}

	sb.WriteString("]")
	return sb.String()
}

func (lst *List[T]) Len() int {
	len := 0

	for el := lst.head ; el != nil ; el = el.next {
		len++
	}

	return len
}

func (lst *List[T]) RemoveAtIndex(index int) (bool, error) {
    len := lst.Len()

	if index < 0 || index >= len  {
		return false, fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		lst.head = lst.head.next
		return true, nil
	}

	if index == len {
		lst.tail = lst.tail.prev
		lst.tail.next = nil
		return true, nil
	}

	prev := lst.head
	curr := lst.head.next
	i := 1

	for curr != nil {
		if(i == index) {
			prev.next = curr.next
			return true, nil
		}

		prev = curr
		curr = curr.next
		i++
    }

	return false, fmt.Errorf("index out of bounds")
}

func (lst *List[T]) InsertAt(index int, t T) (bool, error) {
	len := lst.Len()

	if index < 0 || index > len {
		return false, fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		
		newNode := &node[T]{
			val: t,
			next: lst.head,
		}

		lst.head.prev = newNode
		lst.head = newNode

		return true, nil
	}

	if index == len {
		newNode := &node[T]{
			val: t,
			prev: lst.tail.prev,
		}

		lst.tail.next = newNode
		lst.tail = newNode

		return true, nil
	}

	curr := lst.head
	i := 1

	for curr != nil {
		if i == index {
			tmpNext := curr.next
			newNode := &node[T]{
				val: t,
				prev: curr,
				next: tmpNext,
			}

			curr.next = newNode
			

			return true, nil
		}

		curr = curr.next
		i++
	}

	return false, fmt.Errorf("index out of bounds")
}