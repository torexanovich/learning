package main

import "fmt"

// ask for revenue, expenses & tax rate
// calculate earnings before tax(EBT) and after tax(profit)
// calculate ratio (EBT / profit)
// output EBT, profit and the ratio
func main() {
	revenue := getInput("Revenue: ")
	expenses := getInput("Expenses: ")
	taxeRate := getInput("Taxe Rate: ")

	calculateProfit(revenue, expenses, taxeRate)
}

func getInput(printText string) float64 {
	var toScan float64
	fmt.Print(printText)
	fmt.Scan(&toScan)

	return toScan
}

func calculateProfit(revenue, expenses, taxeRate float64) {
	EBT := revenue - expenses
	profit := EBT * (1 - taxeRate/100)
	ratio := EBT / profit

	fmt.Println("\nEBT: ", EBT)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}
