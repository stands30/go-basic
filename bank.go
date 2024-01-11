package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

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
			break
		case 3:
			fmt.Println("Goodbye!")
			break
		}

		// if choice == 1 {
		// 	fmt.Println("Your Balance is: ", accountBalance)
		// } else if choice == 2 {

		// } else if choice == 3 {
		// 	fmt.Print("Your withdrawl: ")
		// 	var withdrawalAmount float64
		// 	fmt.Scan(&withdrawalAmount)

		// 	if withdrawalAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0. ")
		// 		return
		// 	}

		// 	if withdrawalAmount > accountBalance {
		// 		fmt.Println("Invalid amount. You can't withdraw more than you have ")
		// 		return
		// 	}

		// 	accountBalance -= withdrawalAmount
		// 	fmt.Println("Balance Updated! New Amount: ", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }
		fmt.Println("Thank you for choosing bank!")
	}

}
