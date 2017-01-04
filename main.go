package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addr := flag.String("master", "localhost:5050", "leading mesos master")
	search := flag.String("task", "", "Id/name of target task")
	first := flag.Bool("first", false, "Choose the first task")
	onlyRunning := flag.Bool("running", true, "Show only running tasks")
	pailStdout := flag.Bool("stdout", true, "Print stdout")
	pailStderr := flag.Bool("stderr", true, "Print stderr")
	flag.Parse()

	masterState, err := FetchMasterState(*addr)
	if err != nil {
		fmt.Printf("Error fetching mesos master state: %s\n", err)
		os.Exit(3)
	}

	tasks := masterState.FindTasks(*search, *onlyRunning)

	if len(tasks) == 0 {
		fmt.Printf("Found 0 tasks matching '%s'\n", *search)
		os.Exit(2)
	} else if !*first &&  len(tasks) > 1 {
		fmt.Printf("Found %d tasks matching '%s'. You need to be more specific.\n", len(tasks), *search)
		for _, t := range tasks {
			fmt.Println(t.String())
		}
		os.Exit(1)
	}

	// 1 task found
	task := tasks[0]

	slave := masterState.FindSlave(task.SlaveId)
	if slave == nil {
		fmt.Printf("Error finding slave running your task: %s (%s) with slave id %s\n", task.Name, task.Id, task.SlaveId)
		os.Exit(3)
	}
	slaveState, err := FetchSlaveState(slave)
	if err != nil {
		fmt.Printf("Error fetching mesos agent state: %s\n", err)
		os.Exit(3)
	}

	// 1 task found
	// 1 slave found

	if *pailStdout {
		p, err := NewPailer(os.Stdout, slaveState, &task, "stdout")
		if err != nil {
			fmt.Printf("Error fetching stdout: %s\n", err)
			os.Exit(4)
		}
		p.Start()
	}
	if *pailStderr {
		p, err := NewPailer(os.Stderr, slaveState, &task, "stderr")
		if err != nil {
			fmt.Printf("Error fetching stdout: %s\n", err)
			os.Exit(4)
		}
		p.Start()
	}

	// just wait forever
	select {}
}
