package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8888", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	age := query.Get("age")
	if name == "" {
		name = "indigo Montoya"
	}
	fmt.Fprintf(res, "Hello my name %s is %s", name, age)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	fmt.Fprintf(res, "Bye %s", name)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "HomePage")
}
