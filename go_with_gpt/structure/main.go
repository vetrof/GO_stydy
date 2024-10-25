package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Name string
	Age  int
}

func (u User) Greet() {
	fmt.Printf("Привет, %s! Вам %d лет.\n", u.Name, u.Age)
}

func main() {
	users := make([]User, 0)

	for {
		name := getName()
		if name == "stop" {
			break
		}

		age := getAge()

		user := User{Name: name, Age: age}

		users = append(users, user)
	}

	for _, u := range users {
		u.Greet()
	}
}

func getName() string {
	var input string

	for {
		fmt.Println("Как вас зовут?")
		_, err := fmt.Scanln(&input)

		if err != nil {
			fmt.Println("Ошибка ввода!")
			continue
		}

		break

	}
	return input
}

func getAge() int {
	var input string
	var age int

	for {
		fmt.Println("Сколько вам лет?")
		_, err := fmt.Scanln(&input)

		// Пытаемся преобразовать строку в целое число
		age, err = strconv.Atoi(input)

		if err != nil {
			fmt.Println("Ошибка ввода! Введите корректный возраст числом.")
			continue
		}

		if age <= 0 {
			fmt.Println("Возраст должен быть положительным числом.")
			continue
		}

		break

	}
	return age
}
