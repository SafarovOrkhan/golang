package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

type fileInfo struct {
	fileName  string
	lineCount int
	wordCount int
	hasTarget bool
}

func parser(fileSlice *string, target_ptr string, slice *[]fileInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	var splitter = regexp.MustCompile(`[A-Za-z0-9]+`) //found it from AI
	file_data, err := os.Open(*fileSlice)
	if err != nil {
		log.Fatal(err)
	}
	debug := false
	totalWords := 0
	totalLines := 0
	fileScan := bufio.NewScanner(file_data)
	for fileScan.Scan() {
		totalLines++
		if strings.Contains(fileScan.Text(), target_ptr) {
			debug = true
		}
		words := splitter.FindAllString(fileScan.Text(), -1) // again . AI ...
		totalWords += len(words)                             // AI..
	}

	structData := fileInfo{fileName: *fileSlice, lineCount: totalLines, wordCount: totalWords, hasTarget: debug}

	*slice = append(*slice, structData) // got help from AI

}

func main() {
	var wg sync.WaitGroup
	var fileSlice []string
	var slice []fileInfo

	slice = make([]fileInfo, 0)

	file_ptr := flag.String("files", "", "List of files separated by comma")
	target_ptr := flag.String("target", "", "query string to be searched")

	flag.Parse()

	fileSlice = strings.Split(*file_ptr, ",")

	for _, item := range fileSlice {
		wg.Add(1)
		go parser(&item, *target_ptr, &slice, &wg)
	}

	wg.Wait()
	fmt.Println(slice)

}
