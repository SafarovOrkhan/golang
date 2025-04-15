package main

// https://leetcode.com/problems/multiply-strings/description/
import "fmt"

func multiply(num1 string, num2 string) string {
	sheet := make(map[byte]int)
	sheet[48] = 0
	sheet[49] = 1
	sheet[50] = 2
	sheet[51] = 3
	sheet[52] = 4
	sheet[53] = 5
	sheet[54] = 6
	sheet[55] = 7
	sheet[56] = 8
	sheet[57] = 9

	num1_slice := []byte(num1)
	num2_slice := []byte(num2)

	// for i , val := range sheet {

	// }
}

func main() {
	var value []byte = []byte{1, 2, 3}
	fmt.Println([]int(value))
}
