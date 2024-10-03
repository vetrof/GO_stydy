package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
