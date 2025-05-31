package ds

import "fmt"

type bstNode[T any] struct {
	val T
	left *bstNode[T]
	right *bstNode[T]
}

type BST[T any] struct {
	root *bstNode[T]
}

func (t *BST[T]) Insert(val T) {

}

func (t *BST[T]) Remove(val T) bool {
	return false
}

func (t *BST[T]) Find(val T) (T, error) {
	
	var value T
	return value, fmt.Errorf("value not found")
}

func (t *BST[T]) FindMin() (T, error) {
	var value T
	return value, fmt.Errorf("value not found")
}

func (t *BST[T]) InOrder() (T, error) {
	var value T
	return value, fmt.Errorf("value not found")
}

func (t *BST[T]) PreOrder() (T, error) {
	var value T
	return value, fmt.Errorf("value not found")
}

func (t *BST[T]) PostOrder() (T, error) {
	var value T
	return value, fmt.Errorf("value not found")
}

func (t *BST[T]) Height() int {
	
	return 0
}


func (t *BST[T]) LevelOrder() {

}
