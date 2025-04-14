package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sync"
)

func commandExecuter(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	cmd := fmt.Sprintf("wc -w %v | awk '{print $1}'", filePath)
	result, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v contains %v words\n", filePath, string(bytes.TrimSpace(result)))
}

func main() {
	var wg sync.WaitGroup

	var slice_files []string

	slice_files = []string{"texter.go", "initial.go", "Command_Menu.go", "age-and-loop.go"}

	for _, item := range slice_files {
		wg.Add(1)
		go commandExecuter(item, &wg)
	}

	wg.Wait()
}
