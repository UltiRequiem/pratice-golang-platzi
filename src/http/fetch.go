package main

import (
	"log"
	"net/http"
)

func fetch(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error while fetching %s", url)
	}
	return resp
}
