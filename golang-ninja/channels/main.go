package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		msg <- "yo ho ho 1"
		msg <- "yo ho ho 2"
		close(msg)
	}()

	//fmt.Println(msg)
	//fmt.Println(<-msg)
	//fmt.Println(<-msg)

	for {
		v, ok := <-msg
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println(v)
	}

}
