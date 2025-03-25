package main

import (
	"flag"
	"fmt"
	"tests/todo"
)

func main() {
	//using the flag module to parse command line arguments
	viewFlag := flag.Bool("view", false, "View all todos")
	addFlag := flag.String("add", "", "Add a new task")
	removeFlag := flag.Int("remove", -1, "Remove a task by index")


	if *viewFlag {
		todo.AllTodos()
		return
	}

	if *addFlag != "" {
		todo.AddTodo(*addFlag)
		return
	}

	if *removeFlag != -1 {
		todo.RemoveTodo(*removeFlag)
		return
	}

	fmt.Println("Usage:")
	flag.PrintDefaults()
}
