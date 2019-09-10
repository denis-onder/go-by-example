package main

import "fmt"

func pointerApproach(num *int) {
	*num = 0
}

func nonPointerApproach(num int) {
	num = 0
}

func main() {
	i := 5
	fmt.Printf("Initial value: %d\n", i)
	// -----------------------------------------------------------
	nonPointerApproach(i)
	fmt.Printf("Non-pointer approach value: %d\n", i) // Result => 5
	// -----------------------------------------------------------
	pointerApproach(&i)
	fmt.Printf("Pointer approach value: %d\n", i) // Result => 0
}
