package main

import (
	"flag"
	"fmt"
)

func main() {

	debug_pointer := flag.Bool("debug", false, "enable/disable debug")
	port_pointer := flag.Int("port", 8080, "enable/disable debug")

	flag.Parse()

	fmt.Println("Starting Personal web server ")
	if *debug_pointer == true {
		fmt.Println("Debugging is on!")
	}
	fmt.Printf("Server listening in http://localhost:%v \n", *port_pointer)
}
