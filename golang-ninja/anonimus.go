package main

import "fmt"

func main() {
	x := 10
	func() {
		fmt.Println(x)
	}()
}
