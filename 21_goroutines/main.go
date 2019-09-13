package main

import "fmt"

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, ":", i)
	}
}

func main() {
	f("test")
	go f("wow")
	fmt.Scanln()
}
