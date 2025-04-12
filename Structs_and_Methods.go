package main

import (
	"fmt"
)

type Server struct { //this is private , structs are case sensitive . also fields too.
	// in this case struct is public but fields are private
	name   string
	ip     string
	active bool
}

func printBuilder(serverName, serverIpAddress string, serverActive bool) {
	if serverActive {
		fmt.Printf("Server %v is active at %v\n", serverName, serverIpAddress)
	} else {
		fmt.Printf("Server %v is not active \n", serverName)
	}
}

// second print function is using pointers and struct function type
func (e *Server) printBuilder2() {
	if e.active {
		fmt.Printf("Server %v is active at %v\n", e.name, e.ip)
	} else {
		fmt.Printf("Server %v is not active \n", e.name)
	}
}

func main() {
	server1 := Server{name: "ubuntu", ip: "10.1.1.1", active: true}
	server2 := Server{name: "fedora", ip: "10.1.1.2", active: false}

	//normal value passing
	printBuilder(server1.name, server1.ip, server1.active)
	printBuilder(server2.name, server2.ip, server2.active)

	//directly using struct and method (cuz this method needs struct)
	server1.printBuilder2()
	server2.printBuilder2()

}
