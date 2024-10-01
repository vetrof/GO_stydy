package main

import (
	"fmt"
	"os"
)


func main(){
	first := "none"
	second := "none"
	third := "none"

	args := len(os.Args)

	if args > 1 {
		first = os.Args[1]
	}

	if args > 2{
		second = os.Args[2]
	}

	if args > 3{
		third = os.Args[3]
	}

	fmt.Println(first, second, third)

}