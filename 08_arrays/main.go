package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	fmt.Println(a)
	var twoDimensional [2][3]int
	fmt.Println(twoDimensional)
	for j := 0; j < 2; j++ {
		for k := 0; k < 3; k++ {
			twoDimensional[j][k] = j + k
		}
	}
	fmt.Println(twoDimensional)
}
