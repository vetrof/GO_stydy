package main

import "fmt"

func main() {

	months := [...]string{1: "jan", "feb", "mars", "apr", "may"}
	Q1 := months[2:4]
	Q2 := months[4:]

	fmt.Println(Q1, len(Q1), cap(Q1))
	fmt.Println(Q2, len(Q2), cap(Q2))

}
