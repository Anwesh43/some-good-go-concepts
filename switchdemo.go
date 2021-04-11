package main 

import "fmt"

func main() {
	num := 4
	switch num {
		case 1:
		 fmt.Println("<=1")
		 fallthrough
		
		case 2:
			fmt.Println("<=2")
			fallthrough
		
			
		case 3:
			fmt.Println("<=3")
			fallthrough
			
		case 4:
			fmt.Println("<=4")
			fallthrough
			
		case 5:
			fmt.Println("<=5")
			fallthrough
			
		case 6:
			fmt.Println("<=6")
			fallthrough
			
		case 7:
			fmt.Println("<=7")
			fallthrough
			
		case 8:
			fmt.Println("<=8")

	}
}