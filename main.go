package main

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"os"
)

var TodoList []Todo

type Todo struct {
	Id   uint   `json:"id"`
	Task string `json:"task"`
}

func main() {
	_, err := os.Lstat("todos.json")
	var IsfileExist bool
	if err != nil {
		IsfileExist = false
	} else {
		IsfileExist = true
	}

	if IsfileExist {
		data, err := os.ReadFile("todos.json")
		if err != nil {
			fmt.Println("Error while Reading todos.json")
		}
		err = json.Unmarshal(data, &TodoList)
		if err != nil {
			fmt.Println("Error while unmarshalling todos.json", err)
		}
		fmt.Println(TodoList)
	}

	var newTodo string
	command := os.Args

	switch command[1] {
	case "add":
		var word string
		for _, word = range command[2:] {
			newTodo = fmt.Sprintf("%v %v", newTodo, word)
		}

		noOfTodos := len(TodoList)
		NewTodo := Todo{
			Id:   uint(noOfTodos) + 1,
			Task: newTodo,
		}

		TodoList = append(TodoList, NewTodo)
		data, _ := json.Marshal(TodoList)
		err = os.WriteFile("todos.json", data, 0644)
		if err != nil {
			fmt.Println("Error while saving todo", err)
		} else {
			fmt.Println(newTodo, "saved")
		}

	case "tdlist":

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
