package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type webWriter struct{}

func (webWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func fetch(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error while fetching %s", url)
	}
	return resp
}

func main() {
	io.Copy(webWriter{}, fetch("http://google.com").Body)
}
