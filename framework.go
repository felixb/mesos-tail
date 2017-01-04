package main

type Framework struct {
	Id                 string
	Name               string
	Tasks              []Task
	CompletedTasks     []Task `json:"completed_tasks"`
	Executors          []Executor
	CompletedExecutors []Executor `json:"completed_executors"`
}

func (fw *Framework) FindTasks(search string, onlyRunning bool) []Task {
	tasks := []Task{}
	for _, t := range fw.Tasks {
		if t.Match(search) {
			tasks = append(tasks, t)
		}
	}
	if !onlyRunning {
		for _, t := range fw.CompletedTasks {
			if t.Match(search) {
				tasks = append(tasks, t)
			}
		}
	}
	return tasks
}

func (fw *Framework) Executor(taskId string) *Executor {
	for _, e := range fw.Executors {
		if e.Id == taskId {
			return &e
		}
	}
	for _, e := range fw.CompletedExecutors {
		if e.Id == taskId {
			return &e
		}
	}
	return nil
}
