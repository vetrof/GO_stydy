package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := time.After(30 * time.Second)
	echo := make(chan []byte)

	go readStdin(echo)

	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Done")
			os.Exit(0)
		}
	}
}

func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		i, _ := os.Stdin.Read(data)
		if i > 0 {
			out <- data
		}
	}
}
