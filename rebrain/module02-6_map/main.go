package main

import (
	"fmt"
	"sort"
)

func main() {

	book := map[string]int{
		"book1": 1,
		"book2": 2,
		"book3": 3,
		"book4": 4,
	}

	keys := make([]string, 0, len(book))

	for k := range book {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, v := range keys {
		fmt.Println(v, book[v])
	}

}
