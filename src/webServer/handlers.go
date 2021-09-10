package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{\"years\":8}")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the home page!")
}
