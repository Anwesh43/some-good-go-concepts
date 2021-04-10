package main 

import "fmt"

type User struct {
	name string 
	id int64 
	age int64
}

func (u *User) display() {
	fmt.Printf("My name is %s, age is %d, id is %d", u.name, u.age, u.id)
}

type Player struct {
	User 
	gameName string 
}

func main() {
	p := Player {}
	p.name = "Anwesh"
	p.age = 22
	p.id = 1 
	p.gameName = "fifa21"
	p.display()
}