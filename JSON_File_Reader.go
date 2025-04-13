package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func main() {

	file_data, err := os.Open("/home/orkhan/golang-files/users.json")
	if err != nil {
		log.Fatal(err)
	}

	var users []User

	decoder := json.NewDecoder(file_data)

	decoder.Decode(&users)

	// fmt.Println("decoded", users)

	for _, data := range users {
		fmt.Printf("username : %v , role : %v\n", data.Username, data.Role)
	}
}
