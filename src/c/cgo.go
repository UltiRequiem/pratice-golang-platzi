package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

func Random() int {
	return int(C.random())
}

func main() {
	fmt.Println(Random())
}
