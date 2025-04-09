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

func (a *Admin) ChangeAccess(access bool) {
	a.Access = access
	fmt.Println("Access изменён:", a.Access)
}

func (a *Admin) ChangeRole(role string) {
	a.Role = role
	fmt.Println("Role изменён:", a.Role)
}

func main() {
	// Создаём администратора
	admin := Admin{Name: "Masha", PassNum: 456, Role: "admin", Access: true}

	// Создаём срез только с администраторами, т.к. они реализуют интерфейс UserActions
	greeters := []Greeter{&admin}

	// Приветствуем всех
	for _, greeter := range greeters {
		greeter.Greet()
	}

	// Изменяем доступ и роль администратора
	admin.ChangeAccess(false)
	admin.ChangeRole("superadmin")
}
