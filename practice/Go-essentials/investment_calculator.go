package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Hello, World! Here, I am learning GoLang from Scratch.")

	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var fultureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Print("\nFuture value of the investment is: ", fultureValue)
}
