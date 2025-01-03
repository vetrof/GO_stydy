package main

import "fmt"

type UserActions interface {
	Greet()
	ChangeAccess(access bool)
}

type Greeter interface {
	UserActions
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

func (a Admin) ChangeAccess(access bool) {
	a.Access = access
	fmt.Println("Access изменён:", a.Access)
}

func (a *Admin) ChangeRole(role string) {
	a.Role = role
	fmt.Println("Role изменён:", a.Role)
}

func main() {
	user := User{Name: "vasya", PassNum: 123}
	admin := Admin{Name: "vasya", PassNum: 123, Role: "admin", Access: true}
	greeters := []Greeter{user, admin} // TODO почему ошибка?

	for _, greter := range greeters {
		greter.Greet()
	}

}
