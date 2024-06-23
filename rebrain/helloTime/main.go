package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	formattedTime := now.Format("02.01.2006 15:04")
	fmt.Println("Hello - now -->> ", formattedTime)

	var s = "ПРивет"
	fmt.Println(len(s))
	fmt.Println(s[3])
	fmt.Println(s[1:10])
}