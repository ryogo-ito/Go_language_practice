package main

import "fmt"

type UserError interface {
	Error() string
}

type User struct {
	name string
	age  int
}

func (e *User) Error() string {
	return fmt.Sprintf("User not found %v", e.name)
}

func New(n string, a int) *User {
	return &User{name: n, age: a}
}
