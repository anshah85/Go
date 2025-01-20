package main

import "fmt"

func main() {
	fmt.Println("Welcome to Profit Calculator!")

	revenue := getUserInput("Please enter revenue: ")

	expenses := getUserInput("Please enter expenses: ")

	taxRate := getUserInput("Please enter Tax rate: ")

	earningsBeforeTax, earningsAfterTax, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %0.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings After tax(Profit): %0.2f\n", earningsAfterTax)
	fmt.Printf("Ratio: %0.2f\n", ratio)
}

func calculateProfit(revenue, expenses, taxRate float64) (earningsBeforeTax float64, earningsAfterTax float64, ratio float64) {
	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax * (1 - (taxRate / 100))

	ratio = earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}
