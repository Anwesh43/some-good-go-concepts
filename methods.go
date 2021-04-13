package main 

import (
	"fmt"
)


type Person struct {
	age int64 
	name string 
	id int64
}

func (p *Person) display() {
	fmt.Println("My name is", p.name, "my age is", p.age, "my id is", p.id)
}

func (p * Person) doubleTheAge() {
	p.age *= 2
}

func main() {
	p := Person {
		age:24, 
		name:"Mk",
		id: 1,
	}
	p.display()
	p.doubleTheAge()
	p.display()
}