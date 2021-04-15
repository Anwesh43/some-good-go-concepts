package main 

import "fmt"

type Square struct {
	a int 
}

type Rect struct {
	h int 
	w int 
}

type Shape interface {
	Area() int 
	Perimeter() int 
}

func (s *Square) Area() int {
	return s.a * s.a 
}

func (s *Square) Perimeter() int {
	return 4 * s.a 
}

func (r *Rect) Perimeter() int {
	return 2 * (r.h + r.w)
}

func (r *Rect) Area() int {
	return r.h * r.w 
}

func printMetrics(s Shape) {
	fmt.Println("Area is", s.Area())
	fmt.Println("Perimeter is", s.Perimeter())
}

func main() {
	printMetrics(&Square{10})
	printMetrics(&Rect{10, 20})	
}