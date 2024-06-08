package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	printMessage("yohoho")
	h := sayHello("vasya", 29)
	fmt.Println(h)

	msg, _ := ageCheck(30)
	fmt.Println(msg)

	msg, err := errorReturn(12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Hello, %s! %d", name, age)
}

func ageCheck(age int) (string, bool) {
	if age < 18 {
		return "NO ", false
	} else {
		return "OK ", true
	}
}

func errorReturn(age int) (string, error) {
	if age < 18 {
		return "NO ", errors.New("oh oh")
	} else {
		return "OK ", nil
	}
}
