package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func walker(path string, info fs.FileInfo, err error) error {
	var directories []string
	if info.IsDir() {
		directories = append(directories, info.Name())
	}
	return err
}

func main() {

	var searchString string
	fmt.Printf("Enter Your Search : ")
	fmt.Scan(&searchString)
	path := "/home"
	filepath.Walk(path, walker)
	result := false
	fmt.Printf("Search ==> %v\n", searchString)
	for _, item := range directories {
		if strings.Contains(item, searchString) {
			fmt.Printf("Founded Matches: %v\n", item)
			result = true
		}
	}

	if !result {
		fmt.Println("Found Nothing !!!")
	}

	// fmt.Println("-----------------------")
	// fmt.Println(directories)

}
