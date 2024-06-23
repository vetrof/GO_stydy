package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var x string = "104"
	var y int = 35

	xInt, _ := strconv.Atoi(x)
	yStr := strconv.Itoa(y)

	fmt.Println(xInt, reflect.TypeOf(xInt))
	fmt.Println(yStr, reflect.TypeOf(yStr))

}
