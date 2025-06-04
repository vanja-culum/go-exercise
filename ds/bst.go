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
	} else if node.val > val {
		node.left = t.insertNode(node.left, val)
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
		if node.left == nil && node.right == nil {
			// leaf node
			node = nil
		}

		return nil
	}

	// should be left
	if node.val > val {
		if node.left == nil {
			return fmt.Errorf("node not found, tree irregular to left")
		}

		return t.remove(node.left, val)
	}

	// should be right
	if node.right == nil {
			return fmt.Errorf("node not found, tree irregular to right")
	}

	return t.remove(node.right, val)

}

type Direction int

// const (
// 	Left iota
// 	Right
// )

// func (t *BST[T]) removeParentLink(node *bstNode[T], dir Direction) {

// }

func (t *BST[T]) Remove(val T) error {
	if t.root == nil {
		return fmt.Errorf("tree empty")
	}

	if t.root.val == val {
		t.root = nil
		return nil
	}

	parent := t.root
	var node *bstNode[T]
	var tmpNode *bstNode[T]
	var dir string
	if t.root.val > val {
		node = t.root.left
		dir = "left"
	} else {
		node = t.root.right
		dir = "right"
	}

	if node == nil {
		return fmt.Errorf("node not found from start")
	}

	for node != nil {
		tmpNode = node
		// found it
		if node.val == val {
			// no children nodes
			if node.left == nil && node.right == nil {
				if dir == "left" {
					parent.left = nil
				} else {
					parent.right = nil
				}

				node = nil
				return nil
			}

			// only left node
			if node.left != nil && node.right == nil {
				if dir == "left" {
					parent.left = node.left
				} else {
					parent.right = node.left
				}

				node = nil
				return nil
			}

			// only right node
			if node.left == nil && node.right != nil {
				if dir == "left" {
					parent.left = node.right
				} else {
					parent.right = node.right
				}

				node = nil
				return nil
			}

			// has both child nodes, replace with largest smaller node
			minParent := node
			curr := node.left
			for curr.right != nil {
				tmpCurr := curr
				curr = curr.right
				minParent = tmpCurr
			}

			// unlink parent of largest smaller node if it's not the node itself
			if minParent != node {
				minParent.right = nil
			}

			// set left of largest to the node
			if node.left != curr {
				curr.left = node.left
			}

			// set right of largest to the node
			curr.right = node.right

			// set node to current
			node = curr

			if dir == "left" {
				parent.left = node
			} else {
				parent.right = node
			}


			return nil
		} else if node.val > val {
			// to left
			node = node.left
			dir = "left"
		} else {
			node = node.right
			dir = "right"
		}

		parent = tmpNode
	}

	return fmt.Errorf("node not found")
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

func (t *BST[T]) inOrder(node *bstNode[T]) []T {
	arr := []T{}

	if node == nil {
		return arr
	}

	if node.left != nil {
		arr = append(arr, t.inOrder(node.left)...)
	}

	arr = append(arr, node.val)

	if node.right != nil {
		arr = append(arr, t.inOrder(node.right)...)
	}
	
	return arr
}

func (t *BST[T]) InOrder() []T {
	return t.inOrder(t.root)
}

func (t *BST[T]) preOrder(node *bstNode[T]) []T {
	arr := []T{}
	if node == nil {
		return arr
	}

	arr = append(arr, node.val)

	if node.left != nil {
		arr = append(arr, t.preOrder(node.left)...)
	}

	if node.right != nil {
		arr = append(arr, t.preOrder(node.right)...)
	}

	return arr
}

func (t *BST[T]) PreOrder() []T {
	return t.preOrder(t.root)
}

func (t *BST[T]) postOrder(node *bstNode[T]) []T {
	arr := []T{}
	if node == nil {
		return arr
	}

	if node.left != nil {
		arr = append(arr, t.postOrder(node.left)...)
	}

	if node.right != nil {
		arr = append(arr, t.postOrder(node.right)...)
	}

	arr = append(arr, node.val)

	return arr
}

func (t *BST[T]) PostOrder() []T {
	return t.postOrder(t.root)
}

func (t *BST[T]) BFS()  []T {
	q := Queue[*bstNode[T]]{}
	arr := []T{}
	q.Enqueue(t.root)
	for q.Size() > 0 {
		node, err := q.Dequeue()
		if err != nil {
			continue
		}

		arr = append(arr, node.val)
		if node.left != nil {
			q.Enqueue(node.left)
		}

		if node.right != nil {
			q.Enqueue(node.right)
		}
	}

	return arr
}

func (t *BST[T]) Min() (T, error) {
	if t.root == nil {
		var val T
		return val, fmt.Errorf("tree empty")
	}

	node := t.root

	for node.left != nil {
		node = node.left
	}

	return node.val, nil
}

func (t *BST[T]) Max() (T, error) {
	if t.root == nil {
		var val T
		return val, fmt.Errorf("tree empty")
	}

	node := t.root
	for node.right != nil {
		node = node.right
	}

	return node.val, nil
}

func (t *BST[T]) height(node *bstNode[T]) int {
	if node == nil {
		return 0
	}

	leftHeight := t.height(node.left)
	rightHeight := t.height(node.right)

	if leftHeight > rightHeight {
		return leftHeight + 1 
	}

	return rightHeight + 1
}

func (t *BST[T]) Height() int {
	if t.root == nil {
		return -1
	}

	return t.height(t.root)
}

