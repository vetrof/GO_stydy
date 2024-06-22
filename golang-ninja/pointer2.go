package main

import "fmt"

func main() {
	msg := "yohoho and bottle rome"
	MessageMod(&msg)
	fmt.Println("final print: ", msg)

	summ := 100
	doubler(&summ)
	fmt.Println("summ: ", summ)

}

func MessageMod(message *string) {
	*message += "aaaaa oooo yyyyy"
	//fmt.Println("print in func", *message)
}

func doubler(summ *int){
	*summ = *summ * 2
}

