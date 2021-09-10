package main

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello world!")
	log.Println("Page Refreshed.")
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}
