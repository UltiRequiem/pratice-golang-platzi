package main

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExists, exists := r.FindHandler(request.URL.Path, request.Method)

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Page Not Found")
		return
	}

	if !methodExists {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed!")
		return

	}

	handler(w, request)

	log.Printf("%s visited!", request.URL.Path)

}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exists := r.rules[path]
	handler, methodExists := r.rules[path][method]
	return handler, methodExists, exists
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}
