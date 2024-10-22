package main

import (
	"fmt"
	"strconv"
)

func main() {
	users := make(map[string]int)

	for {
		name := getName()

		if name == "stop" {
			break
		}

		age := getAge()

		users[name] = age

	}

	fmt.Printf("Введенные пары:")
	for name, age := range users {
		fmt.Printf("%s - %d  ", name, age)
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
