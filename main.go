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
var File_Path string
var Dir_Path string
var Tm_Path string = "/home/manavya/task-manager"

type Todo struct {
	Id         string `json:"id"`
	Start_time string `json:"time"`
	End_time   string `json:"start_time"`
	Task       string `json:"task"`
	Mark_done  bool   `json:"mark_done"`
}

func main() {
	dir_err := os.MkdirAll(Tm_Path, 0755)
	if dir_err != nil {
		fmt.Println(dir_err)
		return
	}

	// check if dir_path.txt exits in the tmp folder or not
	_, path_err := os.Lstat("/tmp/dir_path.txt")

	// if dir_path.txt doesn't exists- create
	if path_err != nil {
		data := []byte(Tm_Path)                              // default path would be task manager home path
		err := os.WriteFile("/tmp/dir_path.txt", data, 0644) // create by writing default path in dir_path.txt
		if err != nil {
			fmt.Printf("Error while creating path.txt : %v", err)
			return
		} else {
			Dir_Path = Tm_Path
		}
	}

	// if dir_path already exists

	// then read the current directory from dir_path.txt
	data, readingFile_err := os.ReadFile("/tmp/dir_path.txt")
	if readingFile_err != nil {
		fmt.Println(readingFile_err)
	}

	// set the read directory as Dir_Path
	Dir_Path = string(data)

	_, err := os.Lstat("/tmp/path.txt")
	if err != nil {
		path := fmt.Sprintf("%v/todos.json", Dir_Path)
		data := []byte(path)
		err := os.WriteFile("/tmp/path.txt", data, 0644)
		if err != nil {
			fmt.Printf("Error while creating path.txt : %v", err)
			return
		} else {
			File_Path = path
		}
	}

	data, err = os.ReadFile("/tmp/path.txt")
	if err != nil {
		fmt.Println(err)
	}
	File_Path = string(data)

	TodoList = readTodos(File_Path)

	command := os.Args

	switch command[1] {
	case "add", "a":
		newTodos := command[2:]
		fmt.Println(File_Path)
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

		// check if user has given the first task in command or not
		if len(command) < 4 {
			// default first task
			first_task = "New Todo file Created"
		} else {
			first_task = command[3]
		}
		fmt.Println(Dir_Path)

		// return the formated file path
		path_string := fmt.Sprintf("%v/%v.json", Dir_Path, file_name)

		// saves the formated file path in path.txt
		data := []byte(path_string)
		err := os.WriteFile("/tmp/path.txt", data, 0644)
		if err != nil {
			fmt.Printf("Error while creating path.txt : %v", err)
			return
		}

		// creates a new file with a dummy task
		var todolist []Todo
		// creating first task
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

		// creating the file with first task
		err = os.WriteFile(path_string, todolist_byte, 0644)
		if err != nil {
			fmt.Println("Error while creating", path_string, err)
		}

	case "cf":
		file_name := command[2]
		path_string := fmt.Sprintf("%v/%v.json", Dir_Path, file_name)
		data := []byte(path_string)
		err := os.WriteFile("/tmp/path.txt", data, 0644)
		if err != nil {
			fmt.Printf("Error while changing directory to %v.json : %v", file_name, err)
			return
		}
	case "cd":
		dir_name := command[2]
		path_string := fmt.Sprintf("%v/%v", Dir_Path, dir_name)
		data := []byte(path_string)
		err := os.WriteFile("/tmp/dir_path.txt", data, 0644)
		if err != nil {
			fmt.Printf("Error while changing directory to %v.json : %v", dir_name, err)
			return
		}

	case "pwd":
		fmt.Println(File_Path)
		fmt.Println(Dir_Path)

	case "mkdir":
		dir_name := command[2]
		path_string := fmt.Sprintf("%v/%v", Dir_Path, dir_name)
		err = os.MkdirAll(path_string, 0755)
		if err != nil {
			fmt.Println("Error while creating", path_string, err)
		}

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
	err = os.WriteFile(File_Path, data, 0644)
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
