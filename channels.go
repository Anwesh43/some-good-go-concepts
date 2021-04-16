package main 

import (
	"fmt"
	"time"
)

func printNumsAndSendSum(arr [7]int, chan1 chan int) {
	sum := 0
	for _, r := range arr {
		sum += r 
		time.Sleep(time.Second)
		fmt.Println(r)
	}
	chan1 <- sum
}

func main() {
	chan1 := make(chan int)
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	go printNumsAndSendSum(arr, chan1)
	sum := <- chan1
	fmt.Println("sum is", sum)
}