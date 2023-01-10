package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// define id
var id = 0

// user option
var userOption string

// create a list of UserData
var todos = make([]taskData, 0)

// make a userData struct
type taskData struct {
	id   int
	task string
}

func main() {
	for {
		//greet user
		fmt.Printf("Welcome to the Todo System\n")
		fmt.Printf("Please choose an option below to proceed. \n")

		fmt.Println("###################################")

		//display option to choose from
		fmt.Printf("1: Add a new Todo\n")
		fmt.Printf("2: View all Todos \n")
		fmt.Printf("3: Edit a Todo \n")
		fmt.Printf("4: Mark a todo as complete \n")
		fmt.Printf("5: Remove a todo \n")
		fmt.Printf("6: Quit the program \n")

		fmt.Println("###################################")

		//get option from user
		getOption()

		processOption()
	}
}

func getOption() {

	reader := bufio.NewReader(os.Stdin)
	//Ask user to enter his/her option
	fmt.Printf("Enter a an option: \n")
	option, _ := reader.ReadString('\n')

	//remove delimiter from string
	userOption = strings.Trim(option, "\r\n")

}

func processOption() {
	switch userOption {
	//add task
	case "1":
		var task string

		reader := bufio.NewReader(os.Stdin)
		//Ask user to enter his/her option
		fmt.Printf("Enter Task:\n")
		data, _ := reader.ReadString('\n')

		//remove delimiter from string
		task = strings.Trim(data, "\r\n")

		//generate new ID
		id += 1

		//add data to struct
		var taskData = taskData{
			id:   id,
			task: task,
		}

		//add todoData struct to todos list
		todos = append(todos, taskData)

		fmt.Print("Your Task  has successfully been added to list.\n \n")

		//view all todos
	case "2":
		//loop through the todos list
		for _, todo := range todos {
			//get task
			fmt.Printf("%v: %v \n\n", todo.id, todo.task)
		}

		//edit task
	case "3":
		var taskId string
		reader := bufio.NewReader(os.Stdin)
		//Ask user to enter his/her option
		fmt.Printf("Enter Task ID: \n")
		data, _ := reader.ReadString('\n')

		taskId = strings.Trim(data, "\n, \r")

		//loop throug each todo
		for index, todo := range todos {
			//convert todo.id(int) to string
			if strconv.Itoa(todo.id) == taskId {
				fmt.Printf("%v \n", todo.task)

				var newTask string
				reader := bufio.NewReader(os.Stdin)
				//Ask user to edit previous task
				fmt.Printf("Edit task: \n")
				data, _ := reader.ReadString('\n')

				//remove delimiters
				newTask = strings.Trim(data, "\n,\r")

				//use index from slice to edit task in struct
				todos[index].task = newTask

				//display message to user
				fmt.Println("You successfully updated the task.")

			} else {
				fmt.Printf("No task with Id of %v. Try Again \n\n", taskId)
			}
		}

	case "4":
		var taskId string
		reader := bufio.NewReader(os.Stdin)
		//Ask user to enter his/her option
		fmt.Printf("Enter Taskid: \n")
		data, _ := reader.ReadString('\n')

		taskId = strings.Trim(data, "\n, \r")

		for index, todo := range todos {
			//convert todo.id(int) to string
			if strconv.Itoa(todo.id) == taskId {

				reader := bufio.NewReader(os.Stdin)

				//Ask user to confirm to mark task as completed
				fmt.Printf("Do you want to mark %v as completed? (y/n)\n", todo.task)

				data, _ := reader.ReadString('\n')

				answer := strings.Trim(data, "\n, \r")

				if answer == "y" {
					//use sprintF to save formated output to completed Task
					//strikethrough task to make as completed
					completedTask := fmt.Sprintf("\033[9m%s\033[0m", todo.task)
					//use index from slice to edit task in struct
					todos[index].task = completedTask

					fmt.Printf("Successfully marked task as completed \n\n")
				} else {
					continue
				}

			} else {
				fmt.Printf("No task with Id of %v. Try Again \n\n", taskId)
			}
		}

		//remove a todo
	case "5":

		var taskId string

		//ask user the task id he/she wants to remove
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("Enter Taskid: \n")

		data, _ := reader.ReadString('\n')

		taskId = strings.Trim(data, "\n, \r")

		for index, todo := range todos {
			if strconv.Itoa(todo.id) == taskId {
				//send a confirm message to remove todo task

				reader := bufio.NewReader(os.Stdin)

				fmt.Printf("Are you sure you want to delete %v from todos task (y/n) \n", todo.task)

				data, _ := reader.ReadString('\n')

				userInput := strings.Trim(data, "\n\r")

				if userInput == "y" {
					//using index remove task with same id as that entered by user
					//creates a new todos slice
					todos = append(todos[:index], todos[index+1:]...)

					//send message to user
					fmt.Printf("You successfully removed task. \n\n")
				} else {
					continue
				}

			} else {
				fmt.Printf("No task with Id of %v. Try Again \n\n", taskId)
			}
		}

	default:
		fmt.Printf("You have choose the wrong option.\n\n")
	}
}
