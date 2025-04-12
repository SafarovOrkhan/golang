package main

import (
	"errors" // I will use this
	"fmt"
)

func divide(a, b *int) (float64, error) { // I changed to float to make sure I return with full result
	var result float64
	var err error
	if *b == 0 {
		err = errors.New("divider can't be 0")
		return result, err
	} else {
		result = float64(*a) / float64(*b)
		return result, nil
	}
}

func main() {
	var int1, int2 int //defined in one line

	fmt.Print("Please enter integer 1 : ")
	fmt.Scan(&int1)
	fmt.Print("Please enter integer 2 : ")
	fmt.Scan(&int2)

	result, err := divide(&int1, &int2) // I used pointers , I know it is not what you want
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("result is : ", result)
	}

}
