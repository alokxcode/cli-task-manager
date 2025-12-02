package main

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var TodoList []Todo

type Todo struct {
	Id   uint64 `json:"id"`
	Task string `json:"task"`
}

func main() {
	path := "todos.json"
	TodoList = readTodos(path)

	command := os.Args

	switch command[1] {
	case "add", "a":
		newTodos := command[2:]
		addTodo(newTodos)

	case "list", "ls":
		listTodos(TodoList)

	case "rm", "remove":
		ids := command[2:]
		removeTodos(ids, TodoList)

	case "edit", "e":
		id := command[2]
		newTask := command[3]
		editTodo(id, newTask, TodoList)

	}

}

func readTodos(path string) []Todo {
	var todoList []Todo
	_, err := os.Lstat(path)
	if err != nil {
		return todoList
	} else {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error while Reading todos.json")
		}
		err = json.Unmarshal(data, &todoList)
		if err != nil {
			fmt.Println("Error while unmarshalling todos.json", err)
		}
	}

	return todoList

}

func saveFile(TodoList []Todo) (string, error) {
	data, err := json.Marshal(TodoList)
	if err != nil {
		return "Error while removing todo : marshaling", err
	}
	err = os.WriteFile("todos.json", data, 0644)
	if err != nil {
		return "Error while saving todos", err
	} else {
		return "saved", nil
	}
}

func addTodo(newTodos []string) {
	for _, i := range newTodos {
		noOfTodos := len(TodoList)
		NewTodo := Todo{
			Id:   uint64(noOfTodos) + 1,
			Task: i,
		}
		TodoList = append(TodoList, NewTodo)
	}

	msg, err := saveFile(TodoList)
	if err != nil {
		fmt.Printf(msg, err)
	}
	fmt.Println(msg)
}

func listTodos(TodoList []Todo) {
	if len(TodoList) == 0 {
		fmt.Printf("Todos is empty\n")
	}

	fmt.Printf("\n-----------Todo List-------------\n")
	var todo Todo
	for _, todo = range TodoList {
		fmt.Printf("[%v] -> %v\n", todo.Id, todo.Task)
	}
}

func removeTodos(ids []string, TodoList []Todo) {
	var newTodoList []Todo
	newTodoList = TodoList
	for _, id := range ids {
		var temp []Todo
		for _, todo := range newTodoList {
			if strconv.FormatUint(todo.Id, 10) != id {
				temp = append(temp, todo)
			}
		}
		newTodoList = temp
	}

	for index := range newTodoList {
		newTodoList[index].Id = uint64(index + 1)
	}

	msg, err := saveFile(newTodoList)
	if err != nil {
		fmt.Printf(msg, err)
	}
	fmt.Println("removed")
}

func editTodo(id string, newTask string, TodoList []Todo) {
	for index := range TodoList {
		if strconv.FormatUint(TodoList[index].Id, 10) == id {
			TodoList[index].Task = newTask
		}
	}

	msg, err := saveFile(TodoList)
	if err != nil {
		fmt.Printf(msg, err)
	}
	fmt.Println(msg)
}
