package tree

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadTree(root string) (res []string, err error) {
	files, err := ReadFiles(root)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.isDir {
			paths, err := ReadTree(file.Root + "/" + file.Name)
			if err != nil {
				return nil, err
			}
			for _, path := range paths {
				res = append(res, path)
			}
		} else {
			res = append(res, file.Root+"/"+file.Name)
		}
	}
	return res, nil
}

func ReadFiles(root string) ([]TreeFile, error) {
	var files []TreeFile
	fileInfo, err := ioutil.ReadDir(root)

	if err != nil {
		return nil, err
	}

	for _, file := range fileInfo {
		myFile := TreeFile{Name: file.Name(), isDir: file.IsDir(), Root: root, size: file.Size()}
		myFile.readChildrens()
		files = append(files, myFile)
	}
	return files, nil
}

type TreeFile struct {
	Root   string
	Name   string
	isDir  bool
	childs []TreeFile
	size   int64
}

func (f TreeFile) String() string {
	return fmt.Sprintf("%v size %v | ", f.Name, f.size)
}

func (f *TreeFile) readChildrens() bool {
	if f.isDir {
		childs, err := ReadFiles(f.Root + "/" + f.Name)
		if err == nil {
			f.childs = childs
			return true
		}

		log.Fatalf("Error: %v", err)
		return false
	}
	return true
}
