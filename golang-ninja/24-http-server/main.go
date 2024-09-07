package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

var (
	users = []User{{1, "Vasya"}, {2, "Petya"}}
)

func main() {
	http.HandleFunc("/users", handleUsers)

	fmt.Println("start server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(resp)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users = append(users, user)
}
