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

type CodeHandler struct {}

func NewCodeHandler() CodeHandler {
	return CodeHandler{}
}

func (h CodeHandler) HandleGetDSAFiles(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()

	p1 := filepath.Join(path, "ds", "bst.go")
	p2 := filepath.Join(path, "ds", "list.go")
	p3 := filepath.Join(path, "ds", "max-heap.go")
	p4 := filepath.Join(path, "ds", "min-heap.go")
	p5 := filepath.Join(path, "ds", "priority-queue.go")
	p6 := filepath.Join(path, "ds", "queue.go")
	p7 := filepath.Join(path, "ds", "stack.go")


	files := []string{p1, p2, p3, p4, p5, p6, p7}
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