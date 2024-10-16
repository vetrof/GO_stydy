package main

import "fmt"

type animal interface {
	MakeSound()
}

type cat struct{}
type dog struct{}

func (c *cat) MakeSound() {
	fmt.Println("Meow")
}

func (c *dog) MakeSound() {
	fmt.Println("Woof")
}

func main() {
	var c animal = &cat{}
	var d animal = &dog{}

	c.MakeSound()
	d.MakeSound()
}
