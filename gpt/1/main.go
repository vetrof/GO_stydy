package main

import (
	"fmt"
	"strconv"
)

func main() {

	name := getName()
	age := getAge()

	fmt.Println("Привет,", name, "Вам", age, "лет.")
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