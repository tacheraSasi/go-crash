package todo

import "slices"

import "fmt"

var TodoList []string


func AddTodo(todo string){
	TodoList = append(TodoList, todo)
	fmt.Println("Todo added successfully")
}

func RemoveTodo(id int){
	TodoList = slices.Delete(TodoList, id-1, id)
	fmt.Println("Todo removed successfully")

}

func AllTodos(){
	for i, todo := range TodoList{
		fmt.Printf("%d. %s\n", i+1, todo)
	}

}