package handler

import (
	"go-exercise/ds"
	"go-exercise/helper"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/google/uuid"
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
	NodeCount int    `json:"nodeCount" validate:"required,min=3,max=50"`
	Type      string `json:"type" validate:"required,oneof=balanced unbalanced random perfect"`
}

type BstGenerateTreeNode struct {
	Value int                  `json:"value"`
	Left  *BstGenerateTreeNode `json:"left,omitempty"`
	Right *BstGenerateTreeNode `json:"right,omitempty"`
}

type BstEdge struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type BstNodeData struct {
	Label string `json:"label"`
}

type BstNodePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type BstNode struct {
	Id       string          `json:"id"`
	Data     BstNodeData     `json:"data"`
	Position BstNodePosition `json:"position"`
}

type BstGenerateResponse struct {
	Nodes []BstNode `json:"nodes"`
	Edges []BstEdge `json:"edges"`
}

func (h DsaHandler) HandlePostBstGenerate(w http.ResponseWriter, r *http.Request) {
	body, err := helper.Bind[BstGenerate](r)
	if err != nil {
		helper.ErrInvalidRequest(err, w, r)
		return
	}

	respTree := &BstGenerateResponse{}

	t := ds.BST[int]{}

	count := 0

	for count < body.NodeCount {
		randVal := rand.Intn(body.NodeCount + 1)
		_, err := t.Find(randVal)

		if err != nil {
			t.Insert(randVal)
			count++
		}
	}

	for _, node := range t.InOrderNode() {
		position, _ := t.GenerateCoords(node.Val)
		source := uuid.NewString()

		node := BstNode{
			Position: BstNodePosition{
				X: position.X,
				Y: position.Y,
			},
			Id: source,
			Data: BstNodeData{
				Label: strconv.Itoa(node.Val),
			},
		}

		respTree.Nodes = append(respTree.Nodes, node)
	}

	links := t.GenerateLinks()

	for _, link := range links {
		edge := BstEdge{
			Id: uuid.NewString(),
		}

		if link.Parent != nil {
			for _, respNode := range respTree.Nodes {
				if respNode.Data.Label == strconv.Itoa(link.Child.Val) {
					edge.Source = respNode.Id
				}

				if respNode.Data.Label == strconv.Itoa(link.Parent.Val) {
					edge.Target = respNode.Id
				}
			}

			respTree.Edges = append(respTree.Edges, edge)
		} else {
			// these are root links
			edge.Source = respTree.Nodes[0].Id

			for _, respNode := range respTree.Nodes {
				if respNode.Data.Label == strconv.Itoa(link.Child.Val) {
					edge.Target = respNode.Id
					break
				}
			}

			respTree.Edges = append(respTree.Edges, edge)
		}

	}

	helper.Render(w, r, http.StatusOK, respTree)
}
