package handler

import (
	"go-exercise/helper"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type CodeHandler struct{}

func NewCodeHandler() CodeHandler {
	return CodeHandler{}
}

type FileDetails struct {
	Filename string
	Path     string
	Label    string
}

func (h CodeHandler) HandleGetDSAFiles(w http.ResponseWriter, r *http.Request) {
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
