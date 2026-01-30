package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Error() string {
	return "error"
}

func (u User) Greet() string {
	return fmt.Sprint("My name is ", u.Name)
}

var user User = User{
	Name: "seki🗡️",
	Age:  28,
}

func displayUserGreeting() {
	var n *User = nil
	var e error = n
	fmt.Println(e == nil)
	fmt.Println(user.Greet())
}
