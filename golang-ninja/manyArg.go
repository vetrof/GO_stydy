package main

import "fmt"

func main() {
	min := findMin(18, 6, 45, 7, 5, 7, 1, 3, 2, 43, 65, 65, 3, 2354234, 6, 536, 56, 233)

	fmt.Println(min)
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]
	for i, n := range numbers {
		if n < min {
			min = numbers[i]
		}
	}
	return min
}
