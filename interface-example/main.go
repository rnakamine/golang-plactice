package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u User) Hi() string {
	return fmt.Sprintf("Hi! %s", u.name)
}

type Surperman interface {
	Hi() string
}

func main() {
	var user1 User
	user1 = User{name: "Ryo Nakamine", age: 28}
	fmt.Println(user1.Hi())

	var s Surperman = User{name: "Surperman", age: 30}
	fmt.Println(s.Hi())
}
