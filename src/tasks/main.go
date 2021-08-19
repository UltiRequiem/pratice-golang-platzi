package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (tl *taskList) addToList(t *task) {
	tl.tasks = append(tl.tasks, t)

}

func (tl *taskList) deleteIndex(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
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
func (tl *taskList) printCompletedTasksNames() {
	for index, task := range tl.tasks {
		if task.completed {
			fmt.Println("Task N.", index, ":", task.name, "is finished.")
		}
	}
}

func main() {
	var myTasks *taskList = &taskList{
		tasks: []*task{
			createTask("End my Go course", "End my Golang course on Platzi today.", false),
			createTask("End my Python course", "End my Python course on Platzi today.", false),
			createTask("End my Python course", "End my Nodejs course on Platzi today.", false),
		},
	}

	myTasks.tasks[1].completeTask()

	myTasks.printCompletedTasksNames()
}
