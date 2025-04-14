package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	path := "sample.txt"
	file_data, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fileScan := bufio.NewScanner(file_data)

	fileScan.Scan()
}
