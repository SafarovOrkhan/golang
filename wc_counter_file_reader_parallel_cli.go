package main

import "flag"

type fileInfo struct {
	fileName  string
	lineCount int
	wordCount int
	hasTarget bool
}

func main() {
	var fileSlice []string

	file_ptr := flag.String("files", "", "List of files separated by comma")
	target_ptr := flag.String("target", "", "query string to be searched")

	flag.Parse()

}
