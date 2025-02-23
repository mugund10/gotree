package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"sync"
)

func main() {
	var ds, fil int
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		return
	}

	chs := make(chan string)
	var wg sync.WaitGroup

	// Start root directory traversal
	wg.Add(1)
	go func() {
		fileSystem := os.DirFS(os.Args[1])
		recur(fileSystem, chs, &wg, &ds, &fil)
	}()

	// Close channel when all goroutines complete
	go func() {
		wg.Wait()
		close(chs)
	}()

	// Print results
	for val := range chs {
		_ = val
	}
	if ds < 2 && fil < 2 {
		fmt.Printf(" directory : %d , file : %d \n", ds, fil)
	} else if ds < 2 && fil > 1 {
		fmt.Printf(" directory : %d , files : %d \n", ds, fil)
	} else if ds > 1 && fil < 2 {
		fmt.Printf(" directories : %d , file : %d \n", ds, fil)
	} else {
		fmt.Printf(" directories : %d , files : %d \n", ds, fil)
	}
}

func recur(fileSystem fs.FS, chs chan<- string, wg *sync.WaitGroup, ds, fil *int) {
	defer wg.Done()

	entries, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		// fmt.Printf("Error reading directory: %v\n", err)
		// return
	}

	for _, entry := range entries {
		if entry.Name()[0] != '.' {
			fmt.Println(entry.Name())
			if entry.IsDir() {
				*ds++
				subFS, err := fs.Sub(fileSystem, entry.Name())
				if err != nil {
					// fmt.Printf("Error entering subdirectory %s: %v\n", entry.Name(), err)
					// continue
				}
	
				wg.Add(1)
				go recur(subFS, chs, wg, ds, fil)
			} else {
				*fil++
			}
	
			chs <- entry.Name() + "\n"
		}


	}
}

func List(fsys fs.FS, dirr string, ch chan []fs.DirEntry) {
	defer close(ch)
	dirs, err := fs.ReadDir(fsys, dirr)
	if err != nil {

	}

	// for _, dir := range dirs {
	// 	if dir.IsDir() {
	// 		nch := make(chan []fs.DirEntry)
	// 		nfsys, err := fs.Sub(fsys, dir.Name())
	// 		if err != nil {

	// 		}
	// 		go List(nfsys,".", nch)
	// 		result := <-nch ;
	// 		ch <- result

	// 	} else {

	// 	}

	// }
	ch <- dirs

}

func Walk(fsys fs.FS, name, inden string, ds, fil *int) error {

	dirs, err := fs.ReadDir(fsys, name)
	if err != nil {
		//permission error
	}

	inden1 := inden + "│     "
	inden2 := inden + "      "
	for key, dir := range dirs {

		newname := path.Join(name, dir.Name())
		if key == len(dirs)-1 {
			if dir.IsDir() {
				*ds++
			} else {
				*fil++
			}
			fmt.Printf("%v└──── %s \n", inden, dir.Name())
			Walk(fsys, newname, inden2, ds, fil)

		} else if dir.Name()[0] != '.' {
			if dir.IsDir() {
				*ds++
			} else {
				*fil++
			}
			fmt.Printf("%v├──── %s \n", inden, dir.Name())
			Walk(fsys, newname, inden1, ds, fil)
		}
	}
	return nil
}

