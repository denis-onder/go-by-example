package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["key1"] = "value1"
	m["key2"] = "value2"
	fmt.Println("Map:", m)
	v1 := m["key1"]
	fmt.Println(v1)
	delete(m, "key2")
	fmt.Println("Map:", m)
	_, val := m["key1"]
	fmt.Println(val)
	n := map[string]int{"k1": 1, "k2": 2}
	fmt.Println(n)
}
