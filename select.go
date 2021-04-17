package main 

import (
	"fmt"
	"time"
)
func waitForChannels(c1 chan int, c2 chan int) {
	i := 0
	for {
		select {
		case c1 <- (i * 2):
			 time.Sleep(time.Second)
			 i++
		case <- c2: 
			 fmt.Println("quitting")
			 return 
		} 
	}
}

func main() {
	c1 := make(chan int, 10)
	c2 := make(chan int)
	go func() {
		for j := 0; j <= 10; j++ {
			fmt.Println("num", <-c1)
		}
		c2 <- 0
	}()
	waitForChannels(c1, c2)
	
}