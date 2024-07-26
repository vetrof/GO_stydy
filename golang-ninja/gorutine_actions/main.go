package main

import (
	"fmt"
	"time"
)

var actions = []string{
	"log in",
	"log out",
	"create record",
	"delete record",
	"update record",
}

type LogItem struct {
	action    string
	timeStamp time.Time
}

type User struct {
	id    int
	email string
	logs  []LogItem
}

func (u User) getActivity() string {
	out := fmt.Sprint("id %d   email %s \nActivity log: \n", u.id, u.email)

}
