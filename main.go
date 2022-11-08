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

	flag.Parse()

	todos := &taskmaster.Todos{}
	if err := todos.Load(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("make an app")
		todos.Save(filename)
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
	}
}
