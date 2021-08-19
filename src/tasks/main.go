package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

func createTask(name string, description string, completed bool) task {
	return task{
		name:        name,
		description: description,
		completed:   completed,
	}
}

func main() {
	var myTask task = createTask("End my Go course.", "End mi Golang course on Platzi today.", false)

	fmt.Printf("%+v\n", myTask)
}
