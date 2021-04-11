package main 

import "fmt"

func main() {
	squares := make([]int, 10)
	for i, _ := range squares {
		if a := i + 1; a % 2 == 0 {
			squares[i] = a * a 
		} else {
			squares[i] = a * a * a 
		}
		fmt.Println("num", squares[i])
	}
}