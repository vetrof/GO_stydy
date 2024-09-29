package main

import "fmt"

func main() {
	var list []int
	fmt.Println(list == nil)
	fmt.Println(len(list) == 0)
	fmt.Println(append(list, 1))

	list = []int{}
	fmt.Println(list == nil)
	fmt.Println(len(list) == 0)

	//------

	s := []int{73, 98, 86, 61, 96}
	s2 := s[0:2:2]

	s2 = append(s2, 444)

	fmt.Println("s: ", s, "len:", len(s), "cap:", cap(s))
	fmt.Println("s: ", s2, "len:", len(s2), "cap:", cap(s2))

}
