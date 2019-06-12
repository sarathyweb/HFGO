package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				panic(err)
			} else {
				fmt.Println(filePath)
			}
		}
	}
	return nil
}

func main() {
	defer reportPanic()
	scanDirectory("/home")
}
