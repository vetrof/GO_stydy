package main

import "fmt"

func main() {
	a := []string{"qwert", "ty"}
	x := "ty"
	fmt.Println(a, x, contains(a, x))

	arr := []int{1, 2, 3, 4, 5, 34, 2, 43, 6, 8, 9, 78, 5, 3, 2, 12}
	fmt.Println(getMax(arr...))
}

func contains(a []string, x string) bool {
	for _, element := range a {
		if element == x {
			return true
		}
	}
	return false
}

func getMax(x ...int) int {
	maxDigit := 0
	for _, n := range x {
		if n > maxDigit {
			maxDigit = n
		}
	}
	return maxDigit
}
