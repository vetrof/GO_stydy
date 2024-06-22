package main

import "fmt"

func main() {
	x := 123
	func(number int) {
		fmt.Println(number)
	}(x)
}
