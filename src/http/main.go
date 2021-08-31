package main

import (
	"fmt"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func fetch(url string) *http.Response {
	resp, err := http.Get(url)
	checkErr(err)
	return resp
}

func main() {
	fmt.Println(fetch("https://github.com/UltiRequiem"))
}
