package main

import "io"

func main() {
	io.Copy(webWriter{}, fetch("http://google.com").Body)
}
