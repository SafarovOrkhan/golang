package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func main() {

	file_ptr := flag.String("file", "", "file input")
	flag.Parse()

	if *file_ptr == "" {
		log.Fatal("No file is specified")
	}

	file_data, err := os.Open(*file_ptr)
	if err != nil {
		log.Fatal(err)
	}
	var users []User

	decoder := json.NewDecoder(file_data)
	decoder.Decode(&users)
	fmt.Printf("Loaded %d users from users.json\n", len(users))
	fmt.Printf("========= users =======\n")

	for _, data := range users {
		fmt.Printf("username : %v , role : %v\n", data.Username, data.Role)
	}

}
