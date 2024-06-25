package main

import "fmt"

func main() {
	days := []string{"пн", "вт", "ср", "чт", "пт", "сб", "вс"}
	workDays := make([]string, 5, 5)
	copy(workDays, days)

	var otherDay []string
	for _, day := range days {
		if !contains(workDays, day) {
			otherDay = append(otherDay, day)
		}
	}
	fmt.Println(days)
	fmt.Println(workDays)
	fmt.Println(otherDay)
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
