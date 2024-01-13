package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	// fmt.Println(" time ", randomdata.Number(1, 5))
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------")
		// panic("error")
	}

	for {

		PresentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			break
		case 4:
			fmt.Println("Goodbye!")
			return
		}

		fmt.Println("Thank you for choosing bank!")
	}

}
