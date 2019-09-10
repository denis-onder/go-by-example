package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	sum := 0
	for _, val := range s {
		sum += val
	}
	fmt.Printf("Sum: %d\n", sum)
	for i, number := range s {
		fmt.Printf("Slice iteration: %d - Value: %d\n", i, number)
	}
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range m {
		fmt.Printf("%s -> %d\n", key, value)
	}
	for key := range m {
		fmt.Printf("%s \n", key)
	}
}
