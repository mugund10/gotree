package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var target string

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <dir> <file or folder>")
		return
	}
	rootDir := os.Args[1]
	target = os.Args[2]
	rootAbs, err := filepath.Abs(rootDir)
	if err != nil {
		fmt.Printf("Error resolving absolute path: %v\n", err)
		return
	}
	results := make(chan string)
	var wg sync.WaitGroup
	fsys := os.DirFS(rootAbs)
	wg.Add(1)
	go Lister(fsys, ".", results, &wg, rootAbs)
	go func() {
		wg.Wait()
		close(results)
	}()
	for path := range results {
		fmt.Println(path)
	}
}

func Lister(fsys fs.FS, currentRelPath string, results chan<- string, wg *sync.WaitGroup, rootAbs string) {
	defer wg.Done()
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return
	}
	for _, entry := range entries {
		entryRelPath := path.Join(currentRelPath, entry.Name())
		if strings.Contains(entry.Name(), target) {
			absPath := filepath.Join(rootAbs, entryRelPath)
			results <- absPath
		}
		if entry.IsDir() {
			subfs, err := fs.Sub(fsys, entry.Name())
			if err != nil {
				continue
			}

			wg.Add(1)
			go func(name string) {
				Lister(subfs, path.Join(currentRelPath, name), results, wg, rootAbs)
			}(entry.Name())
		}
	}
}
