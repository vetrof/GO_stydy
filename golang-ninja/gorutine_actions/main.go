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
	out := fmt.Sprintf("id %d   email %s \nActivity log: \n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d) %s  %s \n", i+1, item.action, item.timeStamp)
	}
	return out
}

func main() {
	u := User{
		id:    1,
		email: "qwer@t.y",
		logs: []LogItem{
			{actions[0], time.Now()},
			{actions[1], time.Now()},
			{actions[2], time.Now()},
			{actions[3], time.Now()},
			{actions[4], time.Now()},
		},
	}
	fmt.Println(u.getActivity())
}

func generateUsers(count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@t.y", i+1),
			logs:  generateLogs(),
		}
	}
	return users
}

func generateLogs(count int) []LogItem {
	logs := make([]LogItem, count)

	for i := 0; i < count; i++ {
	}
}
