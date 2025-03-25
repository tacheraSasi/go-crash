package main

import (
	"testing"
	"tests/todo"
)

// Test Add Task
func TestAddTask(t *testing.T) {
	// Reset todo list before test
	todo.TodoList = []string{}

	// Add a task
	todo.AddTodo("Learn Go")

	// Check if the task was added
	if len(todo.TodoList) != 1 || todo.TodoList[0] != "Learn Go" {
		t.Errorf("Expected 'Learn Go' to be added, but got %v", todo.TodoList)
	}
}

// Test Remove Task
func TestRemoveTodo(t *testing.T) {
	// Reset todo list before test
	todo.TodoList = []string{"Learn Go", "Write Tests"}

	// Remove the first task
	todo.RemoveTodo(1)

	// Check if the task was removed
	if len(todo.TodoList) != 1 || todo.TodoList[0] != "Write Tests" {
		t.Errorf("Expected 'Write Tests' to be remaining, but got %v", todo.TodoList)
	}

	todo.RemoveTodo(3)
}

func TestDisplayTasks(t *testing.T) {
	todo.TodoList = []string{"Learn Go", "Write Tests"}

	todo.AllTodos()
}

