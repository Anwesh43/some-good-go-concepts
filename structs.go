package main 

import "fmt"

type Car struct {
	model string 
	price float64 
	make string 
}

func main() {
	car := Car {
		model: "Suzuki",
		price: 300.0,
		make: "Maruti", 
	}
	fmt.Println("Car", car.model, car.price, car.make)
	car1 := Car {}
	fmt.Printf("%+v\n", car1)
	car1.model = "i10"
	car1.price = 400.0 
	car1.make = "Hyundai"
	fmt.Printf("%+v\n", car1)
}