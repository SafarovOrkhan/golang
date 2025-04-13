package main

import (
	"encoding/json"
	"os"
)

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func main() {
	user1 := User{Username: "lisa", Role: "admin"}
	user2 := User{Username: "Sam", Role: "dev"}
	user3 := User{Username: "Sara", Role: "ba"}

	var users []User

	users = append(users, user1, user2, user3)

	// created file
	file, err := os.Create("/home/orkhan/golang-files/users.json")
	if err != nil {
		panic(err)
	}

	// pass NewEncoder to that file to make sure it automatically writes
	json.NewEncoder(file).Encode(users)

}
