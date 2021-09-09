package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	init := time.Now()
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, s := range servers {
		 checkServer(s)
	}

	fmt.Println(time.Since(init))
}

func checkServer(server string) {
	_, err := http.Get(server)

	if err != nil {
		fmt.Printf("%s is down \n", server)
	} else {
		fmt.Printf("%s is working \n", server)
	}
}
