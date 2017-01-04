package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFramework_FindTasks(t *testing.T) {
	t1 := Task{
		Id: "foo",
		Name: "bar",
	}
	t2 := Task{
		Id: "blubb",
		Name: "flubbelwupp",
	}
	t3 := Task{
		Id: "some-id",
		Name: "foobar tasks",
	}
	t4 := Task{
		Id: "blubb",
		Name: "flubbelwupp",
	}

	fw := &Framework{
		Tasks: []Task{t1, t2, },
		CompletedTasks: []Task{t3, t4, },
	}
	assert.Equal(t, []Task{t1, t3}, fw.FindTasks("foo", false))
	assert.Equal(t, []Task{t1}, fw.FindTasks("foo", true))
}
