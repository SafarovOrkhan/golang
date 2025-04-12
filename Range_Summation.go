package main

import (
	"fmt"
)

func main() {
	integer_slice := []int{3, 5, 7, 10, 12}
	var total_sum int16
	for _, val := range integer_slice {
		total_sum += int16(val)
		fmt.Printf("added  %d \n", total_sum)
	}
	fmt.Println("Total number is ", total_sum)

}
