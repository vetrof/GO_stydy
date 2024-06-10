package main

import "fmt"

func Names() (first string, second string) {
	first = "one"
	second = "two"
	return
}

func main() {
	x := 11
	fmt.Print("123", x)
}
