package main

import (
	"strings"
	"fmt"
)

type Slave struct {
	Id       string
	Hostname string
	Pid      string
}

func (s *Slave) Port() string {
	return strings.Split(s.Pid, ":")[1]
}

func (s *Slave) Url() string {
	return fmt.Sprintf("http://%s:%s", s.Hostname, s.Port())
}

func (s *Slave) StateUrl() string {
	dir := strings.Split(s.Pid, "@")[0]
	return fmt.Sprintf("%s/%s/state.json", s.Url(), dir)
}