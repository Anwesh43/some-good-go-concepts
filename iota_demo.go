package main

import (
	. "fmt"
)

const (
	A = 10 * iota 
	B 
	C 
	_
	_ 
	_ 
	_ 
	H
)

const (
	apple, pear = iota + 1, iota + 2 
	banana, melon 
)


func main() {
	Println(A, B, C, H)
	Println(apple, pear, banana, melon)
}