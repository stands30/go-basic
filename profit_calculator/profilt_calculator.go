package main

import (
	"fmt"
	"os"
)

func storeResults(ebt, profit, ratio float64) {

	results := fmt.Sprintf("EBT: %.1f\n Profit: %.1f\n Ratio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("./profit_calculator/profit_calculator.txt", []byte(results), 0644)
}

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.1f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0
	}

	return userInput
}
