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
	add := flag.String("add", "", "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	delete := flag.Int("delete", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")

	flag.Parse()

	todos := &taskmaster.Todos{}

	if err := todos.Load(filename); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case len(*add) != 0:
		todos.Add("first todo")
		err := todos.Save(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *complete > 0:
		todos.Complete(*complete)
		err := todos.Save(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *delete > 0:
		todos.Delete(*delete)
		err := todos.Save(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
	}
}
