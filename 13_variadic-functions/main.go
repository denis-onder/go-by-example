package main

import "fmt"

func formMessage(strings ...string) {
	for _, s := range strings {
		fmt.Println(s)
	}
}

func sumInts(numbers ...int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println(sum)
}

func main() {
	formMessage("Hello", "There")
	sumInts(1, 2, 3, 4)
}
