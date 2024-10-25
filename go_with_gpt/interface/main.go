package main

import "fmt"

type Greeter interface {
	Greet()
}

type User struct {
	Name    string
	PassNum int
}

type Admin struct {
	Role    string
	Name    string
	PassNum int
	Access  bool
}

func (u User) Greet() {
	fmt.Println("Hello", u.Name)
}

func (a Admin) Greet() {
	fmt.Println("Hello", a.Name, "| you access ->", a.Access, "| you role ->", a.Role)
}

func main() {
	user := User{Name: "vasya", PassNum: 123}
	admin := Admin{Name: "vasya", PassNum: 123, Role: "admin", Access: true}
	greeters := []Greeter{user, admin}

	for _, greter := range greeters {
		greter.Greet()
	}

}
