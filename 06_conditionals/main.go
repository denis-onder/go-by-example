package main

import "fmt"

func main() {
	if 2+2 == 4 {
		fmt.Println("2+2=4")
	}
	if num := 4; num > 2 {
		fmt.Printf("%d is greater than 2 \n", num)
	}
	if number := 5; number < 0 {
		fmt.Printf("%d is negative\n", number)
	} else if number < 10 {
		fmt.Printf("%d has one digit\n", number)
	} else {
		fmt.Printf("%d has multiple digits\n", number)

	}
}
