package main

import (
	"fmt"
	"os"
)

func gfgCommand(dirName string) []string {
	var array []string

	dir, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("failed to open Directory")
		return nil
	}

	for _, e := range dir {
		array = append(array, e.Name())
	}

	return array
}

func main() {
	var name string = "Orkhan"
	var anotherName = "Jason"
	var age = 30
	city := "Baku"
	helow := "String"
	gfgCommand("hellow")
	const DefaultTrue string = "True"
	fmt.Println(name, anotherName, age, city, helow)
}
