package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Type your mood : ")
	fmt.Scan(&input)

	if input == "sad" {
		fmt.Println("go watch some standup show")
	} else if input == "happy" {
		fmt.Println("You're good . Cheers")
	} else {
		fmt.Println("Sorry do not undestand")
	}
}
