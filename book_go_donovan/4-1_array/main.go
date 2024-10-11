package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	//var a [3]int
	//var b = [3]int{9, 8, 7}
	a := [...]int{6, 7, 8}

	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Println(i, v)
	}

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("x"))

	fmt.Printf("%x \n%x \n", c1, c2)
	fmt.Println(c1 == c2)
}
