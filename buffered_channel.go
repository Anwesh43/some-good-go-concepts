package main 

import (
	"fmt"
	"time"
)

func sendEven(num int, chan1 chan int) {
	for j := 0; j <= num; j++ {
		chan1 <- (j)
		chan1 <- (j * 2)
		time.Sleep(time.Second)
	}
}

func main() {
	num := 10
	chan1 := make(chan int, 2)
	go sendEven(num, chan1)
	for k:=0; k < 20 ; k++ {
		fmt.Println("num", <-chan1)
	}
}