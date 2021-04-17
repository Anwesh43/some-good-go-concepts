package main 

import (
	"fmt"
	"time"
)


func addNumsAferFourSeconds(a, b int, c chan int) {
	time.Sleep(4 * time.Second)
	c <- a + b 
}

func multiplyNumsAfterTwoSecond(a, b int, c chan int) {
	time.Sleep(2 * time.Second)
	c <- a * b
}

func main() {
	c1 := make(chan int) 
	c2 := make(chan int)
	go addNumsAferFourSeconds(10, 20, c1)
	go multiplyNumsAfterTwoSecond(10, 20, c2)
	for {
		select {
			case res := <-c1:
				fmt.Println("Addition", res) 
			case res := <-c2:
				fmt.Println("Multiplication", res)

			case <-time.After(5 * time.Second):
				fmt.Println("timed out")
				return 
		}
	}
}
