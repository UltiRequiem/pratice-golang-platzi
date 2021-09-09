package main

import (
	"fmt"
	"time"
)

func DoneAsync() chan int {
	r := make(chan int)
	fmt.Println("Warming up ...")
	go func() {
		time.Sleep(3 * time.Second)
		r <- 1
		fmt.Println("Done ...")
	}()
	return r
}

func main () {
	fmt.Println("Let's start ...")
	val := DoneAsync()
	fmt.Println("Done is running ...")
	fmt.Println(<- val)
}
