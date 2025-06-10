package helper

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type FileResult struct {
	Filename string "json:\"filename\""
	Content  string "json:\"content,omitempty\""
	Error    error  "json:\"error,omitempty\""
	Label    string "json:\"label,omitempty\""
}

func ReadFile(path string, filename string, label string, wg *sync.WaitGroup, mu *sync.Mutex, results *[]FileResult) FileResult {
	defer wg.Done()

	res := FileResult{
		Filename: filename,
		Label:    label,
	}

	f, err := os.Open(path)
	if err != nil {
		res.Error = fmt.Errorf("error reading file: %v", filename)
		return res
	}

	defer f.Close()

	content, err := io.ReadAll(f)

	if err != nil {
		res.Error = err
		return res
	}

	res.Content = string(content)
	mu.Lock()
	*results = append(*results, res)
	mu.Unlock()
	return res
}
