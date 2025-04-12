package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Enter Number : ")
	fmt.Scan(&input)

	if s, err := strconv.ParseInt(input, 10, 64); err == nil {
		if s%2 == 0 {
			fmt.Printf("%v is considered even number\n", s)
		} else {
			fmt.Printf("%v is ODD number\n", s)
		}
	} else {
		fmt.Printf("%s is not an integer\n", input)
	}
}
