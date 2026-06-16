package main

import (
	"fmt"
	"os"
	"encoding/json"
	
)

// 1. Our Blueprint
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func main() {
	// Guard Clause: Ensure a command was typed
	if len(os.Args) < 2 {
		fmt.Println("Error: Please type a command (e.g., add, list)")
		return
	}

	command := os.Args[1]

	// Clean Routing
	if command == "add" {
		handleAddTask()
	} else if command == "list" {
		handleListTasks()
	} else {
		fmt.Println("Unknown command. Use 'add' or 'list'.")
	}
}

func handleAddTask() {
	// 2. Mess-free zone for adding tasks
	if len(os.Args)<3 {
		fmt.Println("Please provide a task description")
		return
	}
	description:= os.Args[2]
	fmt.Println("Adding task:", description)

	//instantiating a task 
	newTask := Task{ID: 1, Description: description, Status:"todo"}
	tasks := []Task{newTask}
	
	// we are marshalling tasks so that it can be saved 
	jsonData, err := json.Marshal(tasks)
	if err != nil{
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	//make a choic logic after adding task to either make a new file or add to existing file
	var choice int 
	fmt.Println("Do you want to add to existing file or create a new file.\n Press 1 for add to existing file or 2 to create a new task file")
	fmt.Scanln(&choice)

	if choice == 1{
		//do the adding to existing file action here with like error handling 
		fmt.Println("Enter the existing file name")
		var filename string
		fmt.Scanln(&filename)
		var tasksList []Task
	
		
		existingTask, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading the file/ file doesn't exist")
		}

		err = json.Unmarshal(existingTask,&existingTask) 
		if err != nil {
			fmt.Println("Error unmarshalling JSON")
		}
		
		tasksList = append(tasksList, newTask)

		
		updatedJsonBytes, err := json.Marshal(tasksList)
		if err != nil{
			fmt.Println("Error marshalling the JSON")
		}

		err = os.WriteFile(filename, updatedJsonBytes, 0644)
		if err != nil {
			fmt.Println("Error saving changes to the file")
		}
	}else {
		//do the new file action here.
		fmt.Print("Enter file name: ")
		var filename string 
		fmt.Scanln(&filename)

		//make a json file that will save the tasks
		err = os.WriteFile(filename, jsonData, 0644)
		if err != nil {
		fmt.Println("the file cannot be created", err)
		return
		}
		fmt.Println("Task saved to the file successfully")

	}



	//we make a file with the users preferred name 
	}

func handleListTasks() {
	// 3. Mess-free zone for listing tasks
	
	fmt.Print("What file do you need: ")
	var filename string
	fmt.Scanln(&filename)
	

	fileByte, err := os.ReadFile(filename)
	if err != nil{
		fmt.Println("Error reading the task file")
		return
	}
	
	var listTask []Task

	err = json.Unmarshal(fileByte, &listTask)
	if err != nil {
		fmt.Println("Error unmarshalling the JSON", err)
		return
	}
	for _, task := range listTask{
		fmt.Printf("[%s] ID: %d - %s\n", task.Status, task.ID, task.Description)	
	}
}
