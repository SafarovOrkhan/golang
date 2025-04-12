package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func main() {
	user1 := User{Username: "Alice", Role: "Developer"}
	user2 := User{Username: "Sam", Role: "Administrator"}
	user3 := User{Username: "Lisa", Role: "SRE Engineer"}

	users_slice := []User{}
	users_slice = append(users_slice, user1, user2, user3)

	jsonData, err := json.Marshal(users_slice)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
