package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend!")
	default:
		fmt.Println("It is a weekday.")
	}
	typeSwitch := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I am a boolean")
		case int:
			fmt.Println("I am an integer")
		}
	}
	typeSwitch(false)
	typeSwitch(12)
}
