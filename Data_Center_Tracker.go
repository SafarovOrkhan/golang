package main

import (
	"fmt"
)

type Datacenter struct {
	name     string
	location string
	servers  []Server
}

type Server struct {
	name   string
	ip     string
	active bool
}

func (e *Datacenter) status() {
	for _, val := range e.servers {
		fmt.Printf("datacenter : %v  -> status of %v is %v\n", e.name, val.name, val.active)
	}
}

func main() {
	server1 := Server{name: "ubt-p-1", ip: "10.1.1.1", active: true}
	server2 := Server{name: "ubt-p-2", ip: "10.1.1.2", active: true}
	server3 := Server{name: "ubt-p-3", ip: "10.1.2.1", active: true}
	server4 := Server{name: "ubt-p-4", ip: "10.1.2.2", active: false}
	server5 := Server{name: "ubt-p-5", ip: "10.1.2.2", active: true}

	dc1_servers := []Server{}
	dc1_servers = append(dc1_servers, server1, server2)

	dc2_servers := []Server{}
	dc2_servers = append(dc2_servers, server3, server4, server5)

	dc1 := Datacenter{
		name:     "az_zone1",
		location: "us-east-1",
		servers:  dc1_servers,
	}

	dc2 := Datacenter{
		name:     "az_zone2",
		location: "us-east-2",
		servers:  dc2_servers,
	}

	dc1.status()
	dc2.status()

}
