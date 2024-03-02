package main

import "fmt"

// ask for revenue, expenses & tax rate
// calculate earnings before tax(EBT) and after tax(profit)
// calculate ratio (EBT / profit)
// output EBT, profit and the ratio
func main() {
	var revenue float64
	var expenses float64
	var taxeRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Taxe Rate: ")
	fmt.Scan(&taxeRate)

	EBT := revenue - expenses
	profit := EBT * (1 - taxeRate/100)
	ratio := EBT / profit

	fmt.Println("\nEBT: ", EBT)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}
