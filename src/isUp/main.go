package main

import (
	"fmt"
	"net/http"
)

func main() {

	servers := []string{
		"http://facebook.com",
		"http://platzi.com",
		"http://google.com",
		"http://instagram.com",
		"http://github.com",
	}

	channel := make(chan string)

	for _, s := range servers {
		go checkServer(s, channel)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}

}

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)

	if err != nil {
		channel <- fmt.Sprintf("%s is down.", server)
		return
	}

	channel <- fmt.Sprintf("%s is not down.", server)
}
