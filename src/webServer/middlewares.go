package main

import (
	"log"
	"net/http"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// XD
			flag := true

			log.Println("Checking authentication...")

			if flag {
				f(w, r)
                                log.Println("Passed!")
			} else {
				return
			}
		}
	}
}
