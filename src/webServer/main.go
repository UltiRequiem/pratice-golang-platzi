package main

func main() {
	server := NewServer(3000)

	server.Handle("/", HandleRoot)
	server.Handle("/home", HandleHome)

	server.Handle("/api", server.AddMidleware(HandleAPI, CheckAuth()))

	server.Listen()
}
