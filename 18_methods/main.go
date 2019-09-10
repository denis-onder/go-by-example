package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r *rect) mutateStruct() {
	r.height = r.height * 2
	r.width = r.width * 2
}

func main() {
	rectangle := rect{height: 20, width: 50}
	fmt.Println(rectangle.area())
	rectangle.mutateStruct()
	fmt.Println(rectangle)
}
