package main

import "fmt"

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func (z square) area2() float64 {
	return z.side * z.side * z.side
}

type shape interface {
	area() float64
	area2() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
	fmt.Println(z.area2())
}

func main() {
	s := square{10}
	fmt.Printf("%T\n",s)
	info(s)
}
