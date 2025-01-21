package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
//   => Show error and exit if invalid input
//      - No negative numbers
//      - Not 0
// 2) Show calculated results into file

var fileName = "Profit.txt"

func main() {
	fmt.Println("Welcome to Profit Calculator!")

	revenue, err1 := getUserInput("Please enter revenue: ")

	expenses, err2 := getUserInput("Please enter expenses: ")

	taxRate, err3 := getUserInput("Please enter Tax rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
		// can use panic too
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculateProfit(revenue, expenses, taxRate)
	writeToFile(earningsBeforeTax, earningsAfterTax, ratio)

	fmt.Printf("Earnings before tax: %0.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings After tax(Profit): %0.2f\n", earningsAfterTax)
	fmt.Printf("Ratio: %0.2f\n", ratio)
}

func writeToFile(earningsBeforeTax, earningsAfterTax, ratio float64) {
	profit := fmt.Sprintf("Earnings before tax: %.2f, Earnings after tax: %.2f, Ratio: %.2f", earningsBeforeTax, earningsAfterTax, ratio)
	os.WriteFile(fileName, []byte(profit), 0644)
}

func calculateProfit(revenue, expenses, taxRate float64) (earningsBeforeTax float64, earningsAfterTax float64, ratio float64) {
	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax * (1 - (taxRate / 100))

	ratio = earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number")
	}
	return userInput, nil
}
