package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "Hello"
	s[1] = "From"
	s[2] = "Golang"
	fmt.Println(s)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	fmt.Println(s[1:2])
	t := []string{"g", "h", "j"}
	fmt.Println(t)
}
