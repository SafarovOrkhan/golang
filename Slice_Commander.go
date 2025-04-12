package main

import (
	"fmt"
	//"strconv" not used cuz cli_input is an integer
	"os"
)

func main() {
	var command_slice []string = []string{"kubectl", "docker", "terraform", "aws", "trivy"}
	var cli_input int
	fmt.Print("Please enter your cli index number : ")
	fmt.Scan(&cli_input)

	// checking if input is out of index
	if cli_input >= len(command_slice) || cli_input < 0 {
		fmt.Println("index out of range")
		os.Exit(1)
	}

	// value , err := strconv.ParseInt(cli_input , 10 , 64) will not work cuz I already want a string

	// directly printing slice value
	fmt.Printf("You are selected %v\n", command_slice[cli_input])
	fmt.Println("cap is : ", cap(command_slice))
}
