package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func loadingScreen(val2 int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("Loading : [")
	if val2%2 != 0 {
		// eachTime := 21 / val2
		for i := 0; i < val2; i++ {
			fmt.Print(strings.Repeat("*", 2))
			time.Sleep(1 * time.Second)
		}
	} else {
		eachTime := 20 / val2
		for i := 1; i < val2; i++ {
			fmt.Print(strings.Repeat("*", eachTime))
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Print("]\n")
}

func main() {
	var wg sync.WaitGroup
	var val1 int
	var val2 int
	fmt.Print("Please enter value 1 : ")
	fmt.Scan(&val1)

	fmt.Print("Please enter times : ")
	fmt.Scan(&val2)
	res := val1
	wg.Add(1)
	go loadingScreen(val2, &wg)

	for i := 1; i < val2; i++ {
		res += val1
		time.Sleep(1 * time.Second)

	}

	wg.Wait()
	fmt.Println(res)
}
