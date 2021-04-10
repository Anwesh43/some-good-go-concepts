package main 

import (
	"fmt"
)

func createStrs(k interface{}) {
	a, _ := k.(string)
	fmt.Println("a is", a)
}	

func main() {
	createStrs("hello")
}