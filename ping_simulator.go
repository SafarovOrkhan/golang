package main

import (
	"fmt"
	"log"
	"sync"

	probing "github.com/prometheus-community/pro-bing"
)

func pinger(ipAddress string, wg *sync.WaitGroup) {
	defer wg.Done()
	ping, err := probing.NewPinger(ipAddress)

	if err != nil {
		log.Fatal(err)
	}

	ping.Count = 3

	err = ping.Run()
	if err != nil {
		log.Fatal(err)
	}

	// stats := ping.Statistics()
	// fmt.Println(stats)
	fmt.Printf("Server %v responded\n", ipAddress)
}

func main() {
	var wg sync.WaitGroup

	var ipAdrresses []string

	ipAdrresses = []string{"1.1.1.1", "8.8.8.8", "8.8.4.4"}

	for _, item := range ipAdrresses {
		wg.Add(1)
		go pinger(item, &wg)

	}

	wg.Wait()
}
