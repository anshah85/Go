package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Hello, World! Here, I am learning GoLang from Scratch.\n")

	const inflationRate = 3.5

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fultureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := fultureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value of the investment is: ", fultureValue)
	fmt.Println("Future real value of the investment is: ", futureRealValue)
}
