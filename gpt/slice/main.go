package main

import "fmt"

func main() {
	var name string
	var names []string

	for {
		fmt.Println("Как вас зовут?")
		fmt.Scan(&name)

		if name == "stop" {
			break
		}

		names = append(names, name)
	}

	for i, name := range names {
		if i == len(names)-1 {
			fmt.Printf("%s", name) // Для последнего элемента выводим без запятой
		} else {
			fmt.Printf("%s, ", name)
		}
	}

}
