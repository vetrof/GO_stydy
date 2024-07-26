package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("CONCURRENCY")
	go fmt.Println("CONCURRENCY")
	go fmt.Println("CONCURRENCY")
	go fmt.Println("CONCURRENCY")
	fmt.Println("NO")
	time.Sleep(time.Second)
}
