package main 

import . "fmt"

var Build string 

func main() {
	Println("hello", Build)
}

//go build -ldflags "-X main.Build=World" receive_variable.go