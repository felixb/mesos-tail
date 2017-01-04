package main

import "fmt"

type MasterState struct {
	Id         string
	Hostname   string
	Leader     string
	Slaves     []Slave
	Frameworks []Framework
}

func (state *MasterState) FindTasks(search string, onlyRunning bool) []Task {
	tasks := []Task{}
	for _, fw := range state.Frameworks {
		tasks = append(tasks, fw.FindTasks(search, onlyRunning)...)
	}
	return tasks
}

func (state *MasterState) FindSlave(id string) *Slave {
	for _, s := range state.Slaves {
		if s.Id == id {
			return &s
		}
	}
	return nil
}

func FetchMasterState(leader string) (*MasterState, error) {
	var state MasterState
	err := fetchJson(fmt.Sprintf("http://%s/master/state.json", leader), &state)
	return &state, err
}
