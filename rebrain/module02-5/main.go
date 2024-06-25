package main

import "fmt"

func main() {
	days := []string{"пн", "вт", "ср", "чт", "пт", "сб", "вс"}
	workDays := make([]string, 5, 5)
	copy(workDays, days)

	holydays := days[5:7]
	days = holydays

	fmt.Println(workDays, days)
}
