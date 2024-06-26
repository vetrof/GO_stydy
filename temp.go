package main

import "fmt"

func main() {
	for n := 0; n >= 0; n++ {

		fmt.Println(n)

		if n > 1000 {
			break
		}
	}
	fmt.Println("one print")
	
	fmt.Println("two print")
}
