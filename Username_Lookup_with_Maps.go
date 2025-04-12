package main

import (
	"fmt"
	"os"
)

func main() {
	var map_of_usernames map[string]string //will define it later
	var username_input string

	map_of_usernames = map[string]string{
		"user1":        "orkhan",
		"user2":        "Arif",
		"user3":        "testUser",
		"storageAdmin": "Storage privileged administrator",
		"admin":        "default Administrator",
	}

	fmt.Print("insert username : ")
	fmt.Scan(&username_input)

	// I learned go comma idioms
	// I twisted a little bit so instead of ok I used status
	_, status := map_of_usernames[username_input] // I used "_" this because I dont need value of it
	if !status {                                  // I first used status == false , but this is probably best practice ?
		fmt.Printf("User not recognized -> %v\n", username_input)
		os.Exit(1)
	} else {
		fmt.Println("Full name is : ", map_of_usernames[username_input])
	}
}
