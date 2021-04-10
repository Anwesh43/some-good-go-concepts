package main 

import "fmt"

func main() {
	arr := [...]string{"hello", "more", "demo", "pilage"}
	arr2 := []string{}
	arr1 := arr[1:3]
	
	arr2 = append(arr2, arr[0:len(arr)]...)
	fmt.Printf("%q\n", arr1)
	arr1[1] = "kii"

	fmt.Printf("%q\n", arr)
	fmt.Printf("arr2 %q\n", arr2)
	arr3 := make([]string, 0)
	arr3 = append(arr3, "city")
	arr3 = append(arr3, "country")
	arr2 = append(arr2, arr3...)
	fmt.Printf("%q\n", arr2)
	var k []int 
	fmt.Println(cap(k), len(k))
	fmt.Println(k == nil)
}