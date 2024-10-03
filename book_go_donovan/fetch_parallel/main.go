package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	channel := make(chan string)

	for i, url := range os.Args[1:] {
		go fetch(url, channel, i)
	}

	for range os.Args[1:] {
		fmt.Println("<-channel: ", <-channel)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, channel chan<- string, i int) {
	start := time.Now()
	response, err := http.Get(url)

	if err != nil {
		channel <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, response.Body)
	response.Body.Close()
	if err != nil {
		channel <- fmt.Sprint("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	channel <- fmt.Sprintf("%d -- %.2fs %7d %s", i, secs, nbytes, url)

}
