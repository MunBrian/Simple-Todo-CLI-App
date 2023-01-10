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

// message
var message string

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

	//pass paremeter message to readUserInput
	userInput := readUserInput("Enter an option: \n")

	//get returned value and assign to userOption
	userOption = userInput

}

func processOption() {
	switch userOption {
	//add task
	case "1":

		//get userinput and assign to task
		task := readUserInput("Enter Task: \n")

		//generate new ID
		id += 1

		//add data to struct
		var taskData = taskData{
			id:   id,
			task: task,
		}

		//add todoData struct to todos list
		todos = append(todos, taskData)

		fmt.Printf("Your Task  has successfully been added to list.\n \n")

		//view all todos
	case "2":

		if len(todos) == 0 {
			fmt.Printf("No available tasks. Please select option 1 to add. \n \n")
		} else {
			//loop through the todos list
			for _, todo := range todos {
				//get task
				fmt.Printf("%v: %v \n\n", todo.id, todo.task)
			}
		}

		//edit task
	case "3":
		var taskId string

		taskId = readUserInput("Enter Task ID: \n")

		//loop throug each todo
		for index, todo := range todos {
			//convert todo.id(int) to string
			if strconv.Itoa(todo.id) == taskId {
				fmt.Printf("%v \n", todo.task)

				//get userinput value and assign to newTask
				newTask := readUserInput("Edit task: \n")

				//use index from slice to edit task in struct
				todos[index].task = newTask

				//display message to user
				fmt.Println("You successfully updated the task.")

			}
		}

		//make task as complete
	case "4":

		//get value of userInput and assign to taskId
		taskId := readUserInput("Enter Taskid: \n")

		for index, todo := range todos {
			//convert todo.id(int) to string
			if strconv.Itoa(todo.id) == taskId {

				//save formatted output to message var
				message = fmt.Sprintf("Do you want to mark %v as completed? (y/n)\n", todo.task)

				//get userinput value
				userInput := readUserInput(message)

				if userInput == "y" {
					//use sprintf to save formated output to completed Task
					//strikethrough task to make as completed
					completedTask := fmt.Sprintf("\033[9m%s\033[0m", todo.task)
					//use index from slice to edit task in struct
					todos[index].task = completedTask

					fmt.Printf("Successfully marked task as completed \n\n")
				}

			}

		}

		//remove a todo
	case "5":

		taskId := readUserInput("Enter Taskid: \n")

		for index, todo := range todos {
			if strconv.Itoa(todo.id) == taskId {
				//send a confirm message to remove todo task
				//save formatted output to message var
				message = fmt.Sprintf("Are you sure you want to delete %v from todos task (y/n) \n", todo.task)

				//get userinput value
				userInput := readUserInput(message)

				if userInput == "y" {
					//using index remove task with same id as that entered by user
					//creates a new todos slice
					todos = append(todos[:index], todos[index+1:]...)

					//send message to user
					fmt.Printf("You successfully removed task. \n\n")
				}
			}

		}

		//quit the program
	case "6":

		//get userinput value
		userInput := readUserInput("Are you sure you want to quit the program? (y/n) \n")

		if userInput == "y" {
			//Exit cli
			os.Exit(0)
		}

	default:
		fmt.Printf("You have choosen the wrong option. Try again!! \n\n")
	}
}

func readUserInput(message string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)

	data, _ := reader.ReadString('\n')

	userInput := strings.Trim(data, "\n\r")

	return userInput
}
