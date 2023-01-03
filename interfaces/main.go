package main

import "fmt"

type Shape interface {
	area() float64
}

type Square struct {
	side float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

type myownShape struct{}

func (myownShape) area() float64 {
	//TODO implement me
	return 42
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (r Rectangle) area() float64 {
	return 0.5 * r.length * r.breadth
}

func info(z Shape) {
	fmt.Printf("Shape is %T\n", z)
	fmt.Println("Dimensions are :", z)
	fmt.Println("Area is", z.area())
}

func main() {
	s := Square{side: 10.0}
	r := Rectangle{length: 12.0, breadth: 14.0}
	info(s)
	info(r)
}
