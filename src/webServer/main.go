package main

func main() {
	server := NewServer(3000)

	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/home", HandleHome)

	server.Handle("GET", "/api", server.AddMidleware(HandleAPI, CheckAuth(), Loggin()))

	server.Listen()
}
