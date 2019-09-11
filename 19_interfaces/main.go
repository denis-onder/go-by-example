package main

import "fmt"

type geometry interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type square struct {
	side float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (s square) area() float64 {
	return s.side * s.side
}

func main() {
	s := square{side: 4.5}
	r := rectangle{width: 5.0, height: 3.0}
	measure := func(g geometry) {
		fmt.Println(g.area())
	}
	measure(s)
	measure(r)
}
