package main

import (
	"fmt"
	"maps"
)

type Student struct {
	Name   string
	Scores map[string]int
}

func (e *Student) Average() int {
	sum := 0
	num := 0
	for _, v := range e.Scores {
		sum += v
		num += 1
	}
	return sum / num
}

func PassedSubjects(Scores map[string]int) []string {
	var arr []string
	for v := range maps.Keys(Scores) {
		if Scores[v] >= 50 {
			arr = append(arr, v)
		}
	}
	return arr
}

func main() {
	student1 := Student{Name: "Sara", Scores: map[string]int{"Math": 80, "Science": 75, "Language": 81}}
	student2 := Student{Name: "Lisa", Scores: map[string]int{"Math": 12, "Science": 44, "Language": 92}}
	student3 := Student{Name: "David", Scores: map[string]int{"Math": 76, "Science": 55, "Language": 22}}

	var mapOfStudents []Student
	mapOfStudents = []Student{student1, student2, student3}

	for _, student := range mapOfStudents {
		fmt.Printf("Student %v\n", student.Name)
		fmt.Printf("Average Score %v\n", student.Average())
		arr := PassedSubjects(student.Scores)
		fmt.Printf("Passed Subjects %v\n\n", arr)
	}
}
