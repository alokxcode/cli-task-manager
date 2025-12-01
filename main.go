package main

import (
	// "bufio"
	"fmt"
	"os"
)

var TodoList = make([]Todo, 0)

type Todo struct {
	sNo  uint
	task string
}

func main() {
	var newTodo string
	command := os.Args

	switch command[1] {
	case "add":
		var word string
		for _, word = range command[2:] {
			newTodo = fmt.Sprintf("%v %v", newTodo, word)
		}
		noOfTodos := len(TodoList)
		var todo = Todo{
			sNo:  uint(noOfTodos) + 1,
			task: newTodo,
		}
		TodoList = append(TodoList, todo)
		var todos Todo
		for _, todos = range TodoList {
			fmt.Printf("[%v] -> %v\n", todos.sNo, todos.task)
		}

	case "tdlist":
		var todo Todo
		for _, todo = range TodoList {
			fmt.Printf("%v -> %v", todo.sNo, todo.task)
		}
	}

	// reader := bufio.NewReader(os.Stdin)

	// fmt.Printf("Add Todo -> ")
	// newTodo, _ = reader.ReadString('\n')

	// var todo = Todo{
	// 	task: newTodo,
	// }

	// Todos = append(Todos, todo)

	// println("-------------------Your Todos----------------------")

}
