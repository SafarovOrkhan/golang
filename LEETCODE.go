package main

// https://leetcode.com/problems/multiply-strings/description/
import (
	"fmt"
)

func plusOne(digits []int) []int {
	increaser := 1
	total := 0
	for i := len(digits) - 1; i >= 0; i-- {
		fmt.Println(digits[i])
		total += digits[i] * increaser
		increaser *= 10
	}
	var slc []int
	num := total + 1

	for num > 0 {
		slc = append([]int{num % 10}, slc...)
		num /= 10
	}

	return slc
}

func main() {
	var digits []int
	digits = []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	// digits = []int{4, 3, 2, 1}
	// digits = []int{9}
	// fmt.Println("ada", 1/1)
	res := plusOne(digits)

	fmt.Println(res)
}
