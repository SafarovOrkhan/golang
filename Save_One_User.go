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

	username_ptr := flag.String("username", "", "define username")
	role_ptr := flag.String("role", "", "define role")

	flag.Parse()

	if *username_ptr == "" {
		log.Fatal("username is not defined")
	} else if *role_ptr == "" {
		log.Fatal("role is not defined")
	}

	user := User{Username: *username_ptr, Role: *role_ptr}
	path := "/home/orkhan/golang-files/user-list.json"
	file_data, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal("Error on opening file")
	}

	defer file_data.Close()

	data, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := file_data.WriteString(string(data)) // cannot figure out newline

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Printf("wrote %v to %v \n", user, path)

}
