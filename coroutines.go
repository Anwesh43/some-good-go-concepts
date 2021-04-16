package main 

import (
	"fmt"
	"time"
)

func printSquares(n int) {
	for j := 0; j < n; j++ {
		fmt.Println("square of", j, j * j )
		time.Sleep(time.Second)
	}
}

func main() {
	go printSquares(10)
	time.Sleep(10 * time.Second)
}