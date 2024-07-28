package main

import (
	"fmt"
	"math/rand"
	"os"
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

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("id %d   email %s \nActivity log: \n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d) %s  %s \n", i+1, item.action, item.timeStamp)
	}
	return out
}

func main() {
	rand.Seed(time.Now().Unix())

	users := generateUsers(1000)
	for _, user := range users {
		err := safeUserInfo(user)
		if err != nil {
			return
		}
	}
}

func generateUsers(count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@t.y", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
	}
	return users
}

func generateLogs(count int) []LogItem {
	logs := make([]LogItem, count)

	for i := 0; i < count; i++ {
		logs[i] = LogItem{
			timeStamp: time.Now(),
			action:    actions[rand.Intn(len(actions)-1)],
		}
	}
	return logs
}

func safeUserInfo(user User) error {
	fmt.Println("Write info TO FILE user:%d", user.id)
	filename := fmt.Sprintf("logs/uuid_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActivityInfo())
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
