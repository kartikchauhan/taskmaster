package main

import (
	"flag"
	"fmt"
	"os"

	taskmaster "github.com/cmd/taskmaster"
)

var (
	filename = ".todo.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "complete a task")

	flag.Parse()

	todos := &taskmaster.Todos{}

	if err := todos.Load(filename); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("make an app")
		err := todos.Save(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *complete > 0:
		todos.Complete(*complete)
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
	}
}
