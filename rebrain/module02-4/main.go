package main

import "fmt"

type AmericanVelocity float64
type EuropeanVelocity float64

func main() {

	speedMS := 120.4
	speedKmH := (speedMS / 1000) * 3600

	speedMS2 := 130.0
	speedMH := speedMS2 * 2.237

	var EU EuropeanVelocity = EuropeanVelocity(speedKmH)
	var AM AmericanVelocity = AmericanVelocity(speedMH)

	fmt.Println(EU)
	fmt.Println(AM)
}
