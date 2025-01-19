package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Hello, World! Here, I am learning GoLang from Scratch.")

	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var fultureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Print("\nFuture value of the investment is: ", fultureValue)
}
