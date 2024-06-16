package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Outside gorutine")

	go func() {
		fmt.Println("Inside gorutine")
	}()

	fmt.Println("Outside gorutine again")

	time.Sleep(1 * time.Second)
}
