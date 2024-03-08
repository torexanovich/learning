package main

import "fmt"

func main() {
	fmt.Println("Welcome to the BANK!")
	fmt.Println("Go select one -->")
	fmt.Println("1. Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	fmt.Println("You chose: ", choice)
}
