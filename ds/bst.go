package ds

import (
	"fmt"
)

type ComparableSortable interface {
		~int | ~float64 | ~string
}

type bstNode[T ComparableSortable] struct {
	val T
	left *bstNode[T]
	right *bstNode[T]
}

type BST[T ComparableSortable] struct {
	root *bstNode[T]
}

func (t *BST[T]) insertNode(node *bstNode[T], val T) *bstNode[T] {
	if node == nil {
		return &bstNode[T]{
			val: val,
		}
	}

	if node.val < val {
		node.right = t.insertNode(node.right, val)
		return node
	} else if node.val > val {
		node.left = t.insertNode(node.left, val)
		return node
	}

	 return node
}

func (t *BST[T]) Insert(val T) *bstNode[T] {
	if t.root == nil {
		t.root = t.insertNode(t.root, val)
		return t.root
	}

	return t.insertNode(t.root, val)
}

func (t *BST[T]) String() string {
	return fmt.Sprintf("root: %v\n", t.root)
}

func (t *BST[T]) StringRoot() string {
	return fmt.Sprintf("root: %v\nright: %v\nleft: %v\n", t.root, t.root.right, t.root.left)
}

func (t *BST[T]) remove(node *bstNode[T], val T) error {
	if node == nil {
		return fmt.Errorf("node doesn't exist")
	}

	if node.val == val {
		
		return nil
	}

	return nil
}

func (t *BST[T]) Remove(val T) error {
	return t.remove(t.root, val)
}

func (t *BST[T]) find(node *bstNode[T], val T) (*bstNode[T], error) {
	if node == nil {
		return nil, fmt.Errorf("node not found")
	}

	if node.val == val {
		return node, nil
	}

	if node.val > val {
		return t.find(node.left, val)
	}

	return t.find(node.right, val)
}

func (t *BST[T]) Find(val T) (*bstNode[T], error) {
	return t.find(t.root, val)
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
