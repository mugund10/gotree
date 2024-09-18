package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

func main() {

	for _, val := range os.Args[1:] {
		var ds, fil int
		fileSysytem := os.DirFS(val)
		fmt.Println(fileSysytem)
		Walk(fileSysytem, ".", "", &ds, &fil)
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
