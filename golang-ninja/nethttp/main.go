package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const sampleUrl = "http://api.sampleapis.com/futurama/info"

type LoggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.Response.Status)
			fmt.Println("REDIRECT")
			return nil
		},
		Transport: &LoggingRoundTripper{
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},
		Timeout: time.Second * 10,
	}

	resp, err := client.Get(sampleUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(">>>>>>>>>>>>>>> STATUS %d <<<<<<<<<<<<<<<<", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(body))
}
