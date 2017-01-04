package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMasterState_FindSlave(t *testing.T) {
	s1 := Slave{Id: "slave-id"}
	s2 := Slave{Id: "other-slave-id" }
	m := &MasterState{Slaves: []Slave{s1, s2}}

	assert.Equal(t, s1, *m.FindSlave("slave-id"))
}