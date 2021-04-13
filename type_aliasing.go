package main 

import (
	"fmt"
)

type MyFloat float64 

func (m MyFloat) inverse() float64 {
	return float64(1.0 / float64(m))
}

func main() {
	f := MyFloat(62.0)
	fmt.Println(f.inverse())
}