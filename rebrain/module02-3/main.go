package main

import (
	"fmt"
	"math"
)

func main() {
	r := new(float64)
	c := 35.0
	pi := math.Pi
	*r = c / (2 * pi)

	fmt.Println(*r)

}
