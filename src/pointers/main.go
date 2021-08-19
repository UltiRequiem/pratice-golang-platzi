package main

import "fmt"

func moreFive(a *int) {
	*a += 5
}

func main() {
	var a int = 25

	fmt.Println(a)

	moreFive(&a)

	fmt.Println(a)
}
