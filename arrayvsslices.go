package main 

import "fmt"

func main() {
	arr := [...]string{"hello", "more", "demo", "pilage"}
	arr2 := []string{}
	fmt.Printf("%q\n", arr[1:3])
	arr2 = append(arr2, "hello")
	fmt.Printf("%q\n", arr)
	fmt.Printf("%q\n", arr2)
}