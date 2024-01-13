package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFileName = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFileName, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFileName)

	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}

	return balance, nil
}

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------")
		// panic("error")
	}

	for {

		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Balance is: ", accountBalance)
			break
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0. ")
				return
			}

			accountBalance += depositAmount
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
			break
		case 3:
			fmt.Print("Your withdrawl: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0. ")
				return
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
			break
		case 4:
			fmt.Println("Goodbye!")
			return
		}

		fmt.Println("Thank you for choosing bank!")
	}

}
