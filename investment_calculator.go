package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	//  investmentAmount := 1000
	var investmentAmount float64
	expectedReturnRate := 5.5
	years := 10
	fmt.Println("Investment Amount")
	fmt.Scan(&investmentAmount)

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue := futureValue * math.Pow(1+expectedReturnRate/100, float64(years))

	fmt.Println("future value ", futureValue)
	fmt.Println("futureRealValue ", futureRealValue)

}
