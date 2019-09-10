package main

import "fmt"

// Person struct
type Person struct {
	name string
	age  int
}

// CreatePerson generator
func CreatePerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func main() {
	newPerson := CreatePerson("test", 22)
	fmt.Println("New record:", newPerson)
	fmt.Printf("%s is %d years old.\n", newPerson.name, newPerson.age)
}
