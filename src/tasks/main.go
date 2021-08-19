package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (tl *taskList) addToList(t *task) {
	tl.tasks = append(tl.tasks, t)

}

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
	myTasks := &taskList{
		tasks: []*task{
			createTask("End my Go course.", "End mi Golang course on Platzi today.", false),
			createTask("End my Python course.", "End mi Python course on Platzi today.", false),
		},
	}

	fmt.Printf("%+v\n", myTasks)

}
