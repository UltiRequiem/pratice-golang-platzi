package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["a"] = 8
	myMap["b"] = 8

	fmt.Println(myMap)
}
