package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	arr := strings.Split(s, "")

	re, _ := regexp.Compile("[a-z0-9]")
	var contains []string

	for _, item := range arr {
		match := re.FindStringIndex(item)
		if match != nil {
			contains = append(contains, item)
		}
	}
	// fmt.Println(contains)

	// filteredSource := strings.Join(contains , "")

	var reversed []string

	for i := len(contains) - 1; i >= 0; i-- {
		reversed = append(reversed, contains[i])
	}

	if strings.Join(reversed, "") == strings.Join(contains, "") && len(contains) > 1 {
		fmt.Printf("%v == %v \n", strings.Join(reversed, ""), strings.Join(contains, ""))
		return true
	}
	fmt.Printf("%v != %v \n", strings.Join(reversed, ""), strings.Join(contains, ""))
	return false

}

func main() {
	var value int
	fmt.Print("Enter your input : ")
	fmt.Scan(&value)
	pow := 0
	result := false
	for result == false {
		fmt.Printf("%v^%v : %v\n", value, pow, math.Pow(float64(value), float64(pow)))
		time.Sleep(1 * time.Second)
		pow += 1
	}

}
