package handler

import (
	"encoding/json"
	"fmt"
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

func (h CodeHandler) HandleGetDSAFiles(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()

	bst := filepath.Join(path, "ds", "bst.go")
	list := filepath.Join(path, "ds", "list.go")
	maxHeap := filepath.Join(path, "ds", "max-heap.go")
	minHeap := filepath.Join(path, "ds", "min-heap.go")
	priorityQueue := filepath.Join(path, "ds", "priority-queue.go")
	queue := filepath.Join(path, "ds", "queue.go")
	stack := filepath.Join(path, "ds", "stack.go")

	files := []string{bst, list, maxHeap, minHeap, priorityQueue, queue, stack}
	var mu sync.Mutex
	var wg sync.WaitGroup

	results := make([]helper.FileResult, 0)
	for _, filename := range files {
		wg.Add(1)
		go helper.ReadFile(filename, &wg, &mu, &results)
	}

	wg.Wait()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(results)

	fmt.Println("res", results)

}
