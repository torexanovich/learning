package main

import (
	"fmt"
	"strings"
)

func main() {
	// arrays
	// var a [5]int
	// b := [5]int{1, 2, 3, 4}
	// c := [...]int{1, 2, 3, 4, 5, 6} // ... makes the compiler determine the length
	// a = b
	// fmt.Println(c)

	// slices
	// var d []int = a[1:4] //creates a slice from a[1] to a[3]
	// d = append(d, 2)
	// e := make([]int, len(d)) // create with make
	// copy(e, d) // copies from d to e
	// fmt.Println(e)

	// structs
	// type Name struct {
	// 	FirstName string
	// 	LastName string
	// }
	// type Person struct {
	// 	FullName Name
	// 	Age int
	// }
	// FlutterDeveloper := Person{
	// 	FullName: Name{
	// 		FirstName: "Batir",
	// 		LastName: "Muratbaev",
	// 	},
	// 	Age: 19,
	// }
	// fmt.Println(FlutterDeveloper)

	// if_else()
	// switch_statement()
	// for_loop()
}

// Task: Write a program that takes an integer as input and prints whether the number is positive, negative, or zero
func if_else() {
	var num int
	fmt.Print("num: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("The number is positive")
	} else if num < 0 {
		fmt.Println("The number is negative")
	} else {
		fmt.Println("The number is zero")
	}
}

// Task: Write a program that takes a day of the week as input
// (e.g., "Monday", "Tuesday", etc.) and prints whether it is a weekday or a weekend day
func switch_statement() {
	var day string
	fmt.Print("day: ")
	fmt.Scan(&day)

	switch strings.ToLower(day) {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("It's a weekday")
	case "saturday", "sunday":
		fmt.Println("It's a weekend day")
	default:
		fmt.Println("Invalid day")
	}
}

// Task: Write a program that takes a string as input and prints each character of the string along with its index
func for_loop() {
	var str string
	fmt.Print("String: ")
	fmt.Scan(&str)

	for i := 0; i < len(str); i++ { // for loop
		fmt.Printf("Character at index %d is --> %c\n", i, str[i])
	}
	// for i, v := range str { // range
	// 	fmt.Printf("Character at index %d is --> %c\n", i, v)
	// }
}
