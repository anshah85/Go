package main

import "fmt"

func main() {

	accountBalance := 1000.0

	fmt.Println("Welcome to ABSA Bank!")

	for {
		fmt.Println("What would you like to do today?")
		fmt.Println("1. Check account balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Please enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your account balance is:", accountBalance)
		} else if choice == 2 {
			fmt.Println("Please enter deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Please try again.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Updated account balance: ", accountBalance)

		} else if choice == 3 {
			fmt.Println("Please enter withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid withdrawal amount. Please try again.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Updated account balance: ", accountBalance)
		} else {
			fmt.Println("Exit.")
			break
		}
	}

	fmt.Println("Thank you for choosing ABSA Bank.")

}
