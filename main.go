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

	case "help", "h":
		printGuide()

	case "done", "d":
		id := command[2]
		markDone(id, TodoList)
	default:
		printDefault()
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

func markDone(id string, TodoList []Todo) {
	mark_done := "[done]"
	for index := range TodoList {
		if strconv.FormatUint(TodoList[index].Id, 10) == id {
			TodoList[index].Task = fmt.Sprintf("%v %s", TodoList[index].Task, mark_done)
		}
	}

	msg, err := saveFile(TodoList)
	if err != nil {
		fmt.Printf(msg, err)
	}
	fmt.Println(msg)
}

func printGuide() {
	fmt.Printf("\nTo List all task from the todo list\n")
	fmt.Printf("-> tm list or \n-> tm ls [recomended] \nExample - tm ls  //lists all the task\n\n")

	fmt.Printf("To add a task in todo list\n")
	fmt.Printf("-> tm add 'your task here' \n-> tm a 'your task here'[recomended] \nExample - tm a 'code mf'\n\n")

	fmt.Printf("To remove a task from the todo list\n")
	fmt.Printf("-> tm remove task no. \n-> tm rm task no.[recomended] \nExample - tm rm 1 //removes the first task in the todolist\n\n")

	fmt.Printf("To Rename a task \n")
	fmt.Printf("-> tm edit task no. 'renamed task' or \n-> tm rm task no. 'renamed task'[recomended] \nExample - tm e 2 'did you start?' //rename the 2 task in the todolist to word renamed\n\n")

	fmt.Printf("To remove a task from the todo list\n")
	fmt.Printf("-> tm remove task no. \n-> tm rm task no.[recomended] \nExample - tm rm 2 //removes the 2 task in the todolist\n\n")

	fmt.Printf("To mark a task as done\n")
	fmt.Printf("-> tm done task no. \n-> tm d task no.[recomended] \nExample - tm d 2 //mark the task 2 as done in the todolist\n\n")

	fmt.Printf("For help\n")
	fmt.Printf("-> tm help \n-> tm h [recomended] \n\n")

}

func printDefault() {
	fmt.Printf("Invalid command: Below are the valid commands that you can use")
	printGuide()
}
