package main

type Executor struct {
	Id             string
	Name           string
	Directory      string
	Tasks          []Task
	CompletedTasks []Task `json:completed_tasks`
}
