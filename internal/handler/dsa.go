package handler

import (
	"go-exercise/helper"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type DsaHandler struct{}

func NewDsaHandler() DsaHandler {
	return DsaHandler{}
}

type FileDetails struct {
	Filename string
	Path     string
	Label    string
}

func (h DsaHandler) HandleGetDSAFiles(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()

	bst := FileDetails{
		Filename: "bst",
		Label:    "Binary Search Tree",
		Path:     filepath.Join(path, "ds", "bst.go"),
	}
	list := FileDetails{
		Filename: "list",
		Label:    "List",
		Path:     filepath.Join(path, "ds", "list.go"),
	}
	maxHeap := FileDetails{
		Filename: "max-heap",
		Label:    "Max Heap",
		Path:     filepath.Join(path, "ds", "max-heap.go"),
	}
	minHeap := FileDetails{
		Filename: "min-heap",
		Label:    "Min Heap",
		Path:     filepath.Join(path, "ds", "min-heap.go"),
	}
	priorityQueue := FileDetails{
		Filename: "priority-queue",
		Label:    "Priority Queue",
		Path:     filepath.Join(path, "ds", "priority-queue.go"),
	}
	queue := FileDetails{
		Filename: "queue",
		Label:    "Queue",
		Path:     filepath.Join(path, "ds", "queue.go"),
	}
	stack := FileDetails{
		Filename: "stack",
		Label:    "Stack",
		Path:     filepath.Join(path, "ds", "stack.go"),
	}

	files := []FileDetails{bst, list, maxHeap, minHeap, priorityQueue, queue, stack}
	var mu sync.Mutex
	var wg sync.WaitGroup

	results := make([]helper.FileResult, 0)
	for _, fileDetails := range files {
		wg.Add(1)
		go helper.ReadFile(fileDetails.Path, fileDetails.Filename, fileDetails.Label, &wg, &mu, &results)
	}

	wg.Wait()

	helper.Render(w, r, http.StatusOK, results)

}

type BstGenerate struct {
	NodeCount int `json:"nodeCount" validate:"required,min=3,max=50"`
}

type BstGenerateTreeNode struct {
	Value int                  `json:"value"`
	Left  *BstGenerateTreeNode `json:"left,omitempty"`
	Right *BstGenerateTreeNode `json:"right,omitempty"`
}

type BstGenerateResponse struct {
	root *BstGenerateTreeNode `json:"tree"`
}

func (h DsaHandler) HandlePostBstGenerate(w http.ResponseWriter, r *http.Request) {

	body, err := helper.Bind[BstGenerate](r)
	if err != nil {
		helper.ErrInvalidRequest(err, w, r)
		return
	}

	bstTree := &BstGenerateResponse{}

	insertLeftOrRight(bstTree.root, rand.Intn(body.NodeCount), body.NodeCount, 0)

	helper.Render(w, r, http.StatusOK, bstTree)
}

func insertLeftOrRight(node *BstGenerateTreeNode, value, nodeCount, count int) *BstGenerateTreeNode {
	if count == nodeCount {
		return node
	}

	if node == nil {
		node = &BstGenerateTreeNode{
			Value: value,
		}

		count++

		if count == nodeCount {
			return node
		}
	}

	shouldInsertLeft := rand.Intn(2) == 0

	if shouldInsertLeft {
		node.Left = insertLeftOrRight(node.Left, value, nodeCount, count)
	} else {
		node.Right = insertLeftOrRight(node.Right, value, nodeCount, count)
	}

	return node

}
