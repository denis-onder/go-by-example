package main

import "fmt"

func main() {
	counter := 1
	// One argument for loop
	for counter <= 3 {
		fmt.Println(counter)
		counter++
	}
	// Classic for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// Infinite loop
	for {
		fmt.Println("loop")
		break
	}
	// Continue interation
	for j := 0; j <= 5; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}
}
