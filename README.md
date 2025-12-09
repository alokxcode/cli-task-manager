# **Task Manager CLI (tm)**

A simple, fast, Linux-style **CLI task manager with mini file-system simulator** written in Go.  
Manage tasks, create diffrent files, organize projects, work just like a terminal — except every command starts with **`tm`**.

Perfect for developers who like terminal workflows or want a lightweight, keyboard-driven task manager.

---

## 🚀 **Features**

### ✅ **Task Management**
- Add tasks  
- Edit tasks  
- Delete tasks
- Mark tasks as done 
- List all tasks  

### 📂 **File Operations**
Works like a Linux terminal, but prefixed with `tm`:
- `tm touch <filename>` → create a file  
- `tm ls` → list files  
- `tm cat <file name>` → open file  
- `tm rm <name>` → remove file

### ⚡ **Portable & Fast**
- Single executable  
- Written in Go  
- No dependencies  
- Works offline  

---

## 📥 **Installation**

1. **Download the executable**  
   download from ([here](https://drive.google.com/file/d/1IPxSVK6Vv9Xg3WQYjyF2ZJdzQxWzNIUy/view?usp=drive_link)).

2. **Add it to your system `$PATH`**  
   Add the following line to your `.bashrc`, `.zshrc`, or shell config:

   ```bash
   export PATH="$PATH:/path/to/your/downloaded/executable"

   Replace the path with the path of downloaded executable. - to get the path use pwd
   example :- /home/manavya/downloads
   
3. **Create an alias**  
Add this to your shell config as well:

   ```bash
   alias tm='task-manager'
   ``````


4. **Reload your terminal**  
   ```bash
   source ~/.bashrc
   ``````

or source the configuration file you edited
text

Now you can use tm command in any directory ( tm for task manager )
---

## 🧪 Usage Examples

- Create a new task file:  
```bash
tm touch todo.txt
``````

- Add a task: 
```bash 
tm add "Finish documentation"
``````


- Edit a task:  
```bash
tm edit 2 "Update README introduction"
``````

- Mark a task as done: 
```bash 
tm done 3
``````

- Remove a task: 
```bash 
tm remove 3
``````

- Create file:
```bash  
tm touch projectA
``````


- Open files:  
```bash
tm cat projectA
``````

- List tasks of a file:  
```bash
tm lt
``````

- List all files:  
```bash
tm ls
``````

---

## 📌 Why This Exists

Managing tasks with GUI apps can be heavy and distracting. This CLI tool gives you:

- A lightweight workflow  
- Familiar commands  
- A structured system to organize tasks inside folders and files  
- A more developer-friendly way to stay productive  

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome.  
Feel free to open pull requests or suggestions.

---

