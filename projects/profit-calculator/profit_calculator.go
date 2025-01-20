package main

import "fmt"

func main() {
	fmt.Println("Welcome to Profit Calculator!")

	var revenue float64
	var expenses float64
	taxRate := 2.5

	fmt.Print("Please enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Please enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Please enter Tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - (taxRate / 100))

	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Println("Earnings before tax: ", earningsBeforeTax)
	fmt.Println("Earnings After tax(Profit): ", earningsAfterTax)
	fmt.Println("Ratio: ", ratio)
}
