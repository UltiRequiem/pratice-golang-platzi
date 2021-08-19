package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) completeTask() {
	t.completed = true
}

func (t *task) updateName(name string) {
	t.name = name
}
func (t *task) updateTask(description string) {
	t.description = description
}

func createTask(name string, description string, completed bool) *task {
	return &task{
		name:        name,
		description: description,
		completed:   completed,
	}
}

func main() {
	var myTask *task = createTask("End my Go course.", "End mi Golang course on Platzi today.", false)
	fmt.Printf("%+v\n", myTask)

	// Some hours later
	myTask.completeTask()
	fmt.Printf("%+v\n", myTask)
}
