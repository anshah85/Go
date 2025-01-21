package main

import (
	"fmt"
	"os"
	"strconv"
)

var fileName = "balance.txt"

func readFromFile() float64 {
	data, _ := os.ReadFile(fileName)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)

	return balance

}
func writeToFile(accountBalance float64) {
	balance := fmt.Sprint(accountBalance)
	os.WriteFile(fileName, []byte(balance), 0644)
}

func main() {

	accountBalance := readFromFile()

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

		switch choice {
		case 1:
			fmt.Println("Your account balance is:", accountBalance)
		case 2:
			fmt.Println("Please enter deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Please try again.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Updated account balance: ", accountBalance)
			writeToFile(accountBalance)
		case 3:
			fmt.Println("Please enter withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid withdrawal amount. Please try again.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid withdrawal amount. Please try again.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Updated account balance: ", accountBalance)
			writeToFile(accountBalance)
		default:
			fmt.Println("Exit.")
			fmt.Println("Thank you for choosing ABSA Bank.")
			return
		}
	}
}
