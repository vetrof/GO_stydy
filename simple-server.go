package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleMain)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(403)
	w.Write([]byte("yohoho\n"))
}
