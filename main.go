package main

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

var TodoList []Todo
var Path string

type Todo struct {
	Id         string `json:"id"`
	Start_time string `json:"time"`
	End_time   string `json:"start_time"`
	Task       string `json:"task"`
	Mark_done  bool   `json:"mark_done"`
}

func main() {
	_, err := os.Lstat("/tmp/path.txt")
	if err != nil {
		path := "/tmp/todos.json"
		data := []byte(path)
		err := os.WriteFile("/tmp/path.txt", data, 0644)
		if err != nil {
			fmt.Printf("Error while creating path.txt : %v", err)
			return
		} else {
			Path = path
		}
	}

	data, err := os.ReadFile("/tmp/path.txt")
	if err != nil {
		fmt.Println(err)
	}
	Path = string(data)

	TodoList = readTodos(Path)

	command := os.Args

	switch command[1] {
	case "add", "a":
		newTodos := command[2:]
		fmt.Println(Path)
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
		id := command[2:]
		markDone(id, TodoList)

	case "touch":
		file_name := command[2]
		var first_task string
		if len(command) < 4 {
			first_task = "New Todo file Created"
		} else {
			first_task = command[3]
		}

		path_string := fmt.Sprintf("/tmp/%v.json", file_name)
		data := []byte(path_string)
		err := os.WriteFile("/tmp/path.txt", data, 0644)
		if err != nil {
			fmt.Printf("Error while creating path.txt : %v", err)
			return
		}
		var todolist []Todo
		NewTodo := Todo{
			Id:         "1",
			Start_time: time.Now().Add(5*time.Hour + 30*time.Minute).Format(time.Kitchen),
			End_time:   "",
			Task:       first_task,
			Mark_done:  false,
		}
		todolist = append(todolist, NewTodo)
		fmt.Println(todolist)
		todolist_byte, _ := json.Marshal(todolist)
		err = os.WriteFile(path_string, todolist_byte, 0644)
		if err != nil {
			fmt.Println("Error while creating", path_string)
		}

	case "cd":
		file_name := command[2]
		path_string := fmt.Sprintf("/tmp/%v.json", file_name)
		data := []byte(path_string)
		err := os.WriteFile("/tmp/path.txt", data, 0644)
		if err != nil {
			fmt.Printf("Error while changing directory to %v.json : %v", file_name, err)
			return
		}

	case "pwd":
		fmt.Println(Path)

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
	err = os.WriteFile(Path, data, 0644)
	if err != nil {
		return "Error while saving todos", err
	} else {
		return "file saved", nil
	}
}

func addTodo(newTodos []string) {
	for _, i := range newTodos {
		id := len(TodoList) + 1
		NewTodo := Todo{
			Id:         strconv.FormatInt(int64(id), 10),
			Start_time: time.Now().Add(5*time.Hour + 30*time.Minute).Format(time.Kitchen),
			End_time:   "",
			Task:       i,
			Mark_done:  false,
		}
		TodoList = append(TodoList, NewTodo)
	}

	msg, err := saveFile(TodoList)
	if err != nil {
		fmt.Printf(msg, err)
	} else {
		for _, t := range newTodos {
			fmt.Printf("Added : %v\n", t)
		}
	}
}

func listTodos(TodoList []Todo) {

	fmt.Printf("\n---------------- Todo List ------------------\n")
	if len(TodoList) == 0 {
		fmt.Printf("\n		[ Empty ]				\n\n\nUse this to add tasks in todo list\n-> tm add 'your task'\n")
	}
	var todo Todo
	for _, todo = range TodoList {
		fmt.Printf("[%v] %v%v -> %v\n", todo.Id, todo.Start_time, todo.End_time, todo.Task)
	}
	fmt.Printf("\n")
}

func removeTodos(ids []string, TodoList []Todo) {
	var newTodoList []Todo
	var ids_task []string
	newTodoList = TodoList
	for _, id := range ids {
		var temp []Todo
		for _, todo := range newTodoList {
			if todo.Id != id {
				temp = append(temp, todo)
			} else {
				ids_task = append(ids_task, todo.Task)
			}
		}
		newTodoList = temp
	}

	for index := range newTodoList {
		newTodoList[index].Id = strconv.FormatInt(int64(index+1), 10)
	}

	msg, err := saveFile(newTodoList)
	if err != nil {
		fmt.Printf(msg, err)
	} else {
		for _, t := range ids_task {
			fmt.Printf("removed : %v\n", t)
		}
	}
}

func editTodo(id string, newTask string, TodoList []Todo) {
	var prev_task string
	for index := range TodoList {
		if TodoList[index].Id == id {
			prev_task = TodoList[index].Task
			TodoList[index].Task = newTask
		} else if TodoList[index].Id == fmt.Sprintf("\033[2m%v\033[0m", id) {
			fmt.Printf("[%v] -> %v : Can't rename marked tasks\n", TodoList[index].Id, TodoList[index].Task)
			return
		}
	}

	msg, err := saveFile(TodoList)
	if err != nil {
		fmt.Printf(msg, err)
	} else {
		fmt.Printf("[%v] -> %s ---> %v\n", id, prev_task, newTask)
	}
}

func markDone(ids []string, TodoList []Todo) {
	mark_done := "\033[32m[done]\033[0m"
	done_time := time.Now().Add(5*time.Hour + 30*time.Minute).Format(time.Kitchen)
	var id_task []string
	for _, id := range ids {
		id_int, _ := strconv.Atoi(id)
		if id_int > len(TodoList) {
			fmt.Printf("[%v] -> Task Id not found\n", id_int)
		}
		for index := range TodoList {
			if TodoList[index].Id == id {
				if !TodoList[index].Mark_done {
					id_task = append(id_task, TodoList[index].Task)
					TodoList[index].Id = fmt.Sprintf("\033[2m%v\033[0m", TodoList[index].Id)
					TodoList[index].Start_time = fmt.Sprintf("\033[2m%v : \033[0m", TodoList[index].Start_time)
					TodoList[index].End_time = fmt.Sprintf("\033[2m%v\033[0m", done_time)
					TodoList[index].Task = fmt.Sprintf("\033[2m%v\033[0m %s", TodoList[index].Task, mark_done)
					TodoList[index].Mark_done = true
				} else {
					fmt.Printf("%v : Aready marked as done\n", TodoList[index].Task)
				}
			} else if TodoList[index].Id == fmt.Sprintf("\033[2m%v\033[0m", id) {
				fmt.Printf("[%v] -> %v : Aready marked as done\n", TodoList[index].Id, TodoList[index].Task)
			}
		}
	}

	msg, err := saveFile(TodoList)
	if err != nil {
		fmt.Printf(msg, err)
	} else {
		for _, t := range id_task {
			fmt.Printf("marked done : %v\n", t)
		}
	}
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
