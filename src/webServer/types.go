package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct {
	Name  string
	Email string
	Phone string
}

func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
