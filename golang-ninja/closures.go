package main

import "fmt"

func main() {
	inc := increment()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

}

// closures func
func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
