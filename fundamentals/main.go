package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name   string
	age    int
	gender string
}

func addStudent(name string, age int, gender string) Student {
	var currentStudent Student
	currentStudent.name = name
	currentStudent.age = age
	currentStudent.gender = gender
	return currentStudent
}
func printStudent(current Student) {
	fmt.Println("Name:" + current.name)
	fmt.Println("Age: " + strconv.Itoa(current.age))
	fmt.Println("Gender: " + current.gender)
}
func main() {

	input := true
	var class []Student
	var name string
	var age int
	var gender string
	var response string

	fmt.Println("Welcome to University!")
	for {
		var tempStudent Student
		fmt.Println("Would you like to add a student?")
		fmt.Scanln(&response)
		if response == "No" || response == "no" {
			input = false
		}
		if !input {
			break
		}
		fmt.Println("What is your name?")
		fmt.Scanln(&name)
		fmt.Println("What is your age?")
		fmt.Scanln(&age)
		fmt.Println("What is your gender?")
		fmt.Scanln(&gender)
		tempStudent = addStudent(name, age, gender)
		class = append(class, tempStudent)
	}
	for i := 0; i < len(class); i++ {
		printStudent(class[i])
		fmt.Println()
	}

}
