package main

import (
	"fmt"
)

func add_two_integers(firstInt int, secondInt int) int {
	result := firstInt + secondInt
	return result
}

func printMenu(input_slice []string) {
	for i, input_slice_item := range input_slice {
		fmt.Printf("%d) %s\n", i, input_slice_item)
	}
}

func main() {
	var input int
	menu_slice := []string{"Say Hello", "Add two numbers", "Exit"}
	for input != 2 {
		printMenu(menu_slice)
		fmt.Print("Enter Number : ")
		fmt.Scan(&input)

		switch input {
		case 0:
			fmt.Println("Hello")
		case 1:
			var int1 int
			var int2 int
			fmt.Print("Enter Value 1 : ")
			fmt.Scan(&int1)
			fmt.Print("Enter Value 2 : ")
			fmt.Scan(&int2)
			add_result := add_two_integers(int1, int2)
			fmt.Println("adding ... : ", add_result)
		case 2:
			fmt.Println("Exiting ...")
		}
	}

}
