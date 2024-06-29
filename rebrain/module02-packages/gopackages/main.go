package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/vetrof/utils"
	utilsV2 "github.com/vetrof/utils/v2"
	"gopackage/wordz"
)

func main() {
	fmt.Println("Hello!")
	color.Red("Hello!")

	for i := 0; i < 5; i++ {
		fmt.Println(wordz.Hello)
		fmt.Println(wordz.Random())
	}

	s := []string{"asd", "fdds"}
	x := "qwerty"
	fmt.Println(utils.Contains(s, x))
	fmt.Println(utilsV2.Containss(s, x))

}
