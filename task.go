package main

import (
	"strings"
	"fmt"
)

type Task struct {
	Id          string
	Name        string
	SlaveId     string `json:"slave_id"`
	FrameworkId string `json:"framework_id"`
	State       string
}

func (t *Task) String() string {
	return fmt.Sprintf("%s (%s) %s", t.Name, t.Id, t.State)
}

func (t *Task) Match(search string) bool {
	return strings.Contains(t.Id, search) || strings.Contains(t.Name, search)
}