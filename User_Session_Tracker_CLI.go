package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type UserStruct struct {
	Username string `json:"username"`
	Action   string `json:"action"`
	Time     string `json:"time"`
}

func main() {

	username_ptr := flag.String("username", "", "Define Username")
	action_ptr := flag.String("action", "", "Define Action login/logout")
	list_ptr := flag.Bool("list", false, "List all active sessions")
	flag.Parse()

	path := "/home/orkhan/golang-files/users.json"

	var users []UserStruct

	file_for_decode, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file_for_decode.Close()

	decoder := json.NewDecoder(file_for_decode)
	decoder.Decode(&users)
	// fmt.Println("after decoding ", users)
	if *username_ptr == "" || *action_ptr == "" && *list_ptr != true {
		log.Fatal("bad usage : username must be set alongside with action")
	}

	if *username_ptr != "" && *action_ptr != "" && *list_ptr != true {
		user := UserStruct{Username: *username_ptr, Action: *action_ptr, Time: time.Now().Format(time.RFC3339)}

		if *action_ptr == "login" || *action_ptr == "logout" {
			fmt.Println("proceeding with creation")
		} else {
			log.Fatal("Wrong action option value")
		}

		for _, data := range users {
			if data.Username == user.Username && data.Action == user.Action {
				log.Fatal("Username with same action is already exists ")
			}
		}

		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}

		defer file_for_decode.Close()

		users = append(users, user)

		json.NewEncoder(file).Encode(users)
		// fmt.Println("last users list", users)

	} else if *action_ptr != "" && *list_ptr == true {
		log.Fatal("cannot use --list option like this")
	}

	if users == nil {
		log.Fatal("sessions are empty")
	}

	if *list_ptr == true && *username_ptr != "" {
		for _, data := range users {
			var listedUsers []UserStruct
			if data.Username == *username_ptr {
				listedUsers = append(listedUsers, data)
			}
			fmt.Println(listedUsers)
		}
	} else {
		fmt.Println(users)
	}

}
