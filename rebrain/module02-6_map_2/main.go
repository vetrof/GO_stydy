package main

import "fmt"

func main() {
	// Инициализация вложенной карты
	literature := make(map[string]map[string][]string)

	// Создание внутренней карты
	innerMap1 := make(map[string][]string)
	innerMap2 := make(map[string][]string)
	innerMap1["mumu"] = []string{"poem"}
	innerMap2["mazay"] = []string{"drama"}

	// Присвоение внутренней карты внешней карте
	literature["masha"] = innerMap1
	literature["vasya"] = innerMap2

	// Вывод карты
	fmt.Println(literature)
}
