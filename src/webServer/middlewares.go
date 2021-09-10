package main

import (
	"log"
	"net/http"
	"time"
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

func Loggin() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()

			f(w, r)
		}
	}
}
