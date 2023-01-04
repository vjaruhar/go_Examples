package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	length float64
	widht  float64
}

func (r Rect) Area() float64 {
	//fmt.Println("Area is :"+ r.length*r.widht);
	return r.length * r.widht
}

func main() {
	//var s Shape

	r := Rect{0.5, 0.5}
	fmt.Println(r.Area())
}
