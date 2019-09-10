package main

import "fmt"

func addTwoNumber(a int, b int) int {
	return a + b
}

func threeWordedMessage(s1, s2, s3 string) {
	fmt.Printf("%s %s %s\n", s1, s2, s3)
}

func returnTwoInts() (int, int) {
	return 2, 5
}

func main() {
	fmt.Println(addTwoNumber(2, 2))
	threeWordedMessage("Hello", "from the", "other side")
	int1, int2 := returnTwoInts()
	fmt.Println(int1, int2)
	_, val := returnTwoInts()
	fmt.Println(val)
}
