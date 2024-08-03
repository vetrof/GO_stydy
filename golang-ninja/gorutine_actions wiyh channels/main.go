package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
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
	timeNow := time.Now()
	waitgroup := &sync.WaitGroup{}

	users := make(chan User)
	go generateUsers(1000, users)
	for user := range users {
		waitgroup.Add(1)
		go safeUserInfo(user, waitgroup)

	}
	waitgroup.Wait()
	fmt.Println("all time --->>> : ", time.Since(timeNow))
}

func generateUsers(count int, users chan User) {
	for i := 0; i < count; i++ {
		users <- User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@t.y", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
	}
	close(users)
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

func safeUserInfo(user User, waitgroup *sync.WaitGroup) error {
	fmt.Println("Write info TO FILE user:%d", user.id)
	time.Sleep(time.Millisecond * 10)
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
	waitgroup.Done()
	return nil
}
