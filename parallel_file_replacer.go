package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func replacer(mapItem map[int]*string, wg *sync.WaitGroup) {

}

func main() {
	var wg sync.WaitGroup
	path := "sample.txt"
	file_data, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fileScan := bufio.NewScanner(file_data)
	num := 1
	fileMap := make(map[int]string)
	for fileScan.Scan() {
		fileMap[num] = fileScan.Text()
		num++
		go replacer(&fileMap[num], &wg)
	}

	fmt.Println(fileMap)
}
