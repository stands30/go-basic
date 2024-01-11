package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	years := 10

	fmt.Println("Investment Amount")
	fmt.Scan(&investmentAmount)

	fmt.Println("Expected Return Rate")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("Years")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	fmt.Printf("future value %.2f\n", futureValue)
	fmt.Println("futureRealValue ", futureRealValue)

}

func calculateFutureValue(investmentAmount float64, expectedReturns float64, years int) (fv float64, rfv float64) {
	fv = float64(investmentAmount) * math.Pow(1+expectedReturns/100, float64(years))
	rfv = fv * math.Pow(1+expectedReturns/100, float64(years))

}
