package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	expectedReturnRate := 5.5
	years := 10

	fmt.Println("Investment Amount")
	fmt.Scan(&investmentAmount)

	fmt.Println("Expected Return Rate")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("Years")
	fmt.Scan(&years)

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue := futureValue * math.Pow(1+expectedReturnRate/100, float64(years))

	fmt.Printf("future value %.2f", futureValue)
	fmt.Println("futureRealValue ", futureRealValue)

}
