package main

import "fmt"

func twoPartHello() func(s string) {
	const msg = "Hello,"
	return func(s string) {
		fmt.Println(msg, s)
	}
}

func main() {
	twoPartHello()("Test One")
	twoPartHello()("Test Two")
	twoPartHello()("Test Three")
}
