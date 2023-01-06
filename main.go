package main

import "fmt"

//define id
var id = 0

//create a list of UserData
var todos = make([]userData, 0)

//make a userData struct
type userData struct {
	id   int
	data string
}

func main() {
	for {
		//greet user
		fmt.Printf("Welcome to the Todo System\n")
		fmt.Printf("Please choose an option. \n")

		//display option to choose from
		fmt.Printf("1: Add a new Todo\n")
		fmt.Printf("2: View all Todos \n")
		fmt.Printf("3: Edit a Todo \n")
		fmt.Printf("4: Mark a todo as incomplete \n")
		fmt.Printf("5: Remove a todo \n")
		fmt.Printf("6: Quit the program \n")

		//define user input data type
		var userInput string

		//get user's input
		fmt.Println("Enter your option: ")
		fmt.Scan(&userInput)

		//check input validity
		if userInput > "6" || userInput <= "0" {
			println("You entered a wrong option value. Pick option 1-6 as stated above.")
			continue
		}

		switch userInput {
		case "1":
			var data string

			fmt.Printf("Enter a todo:\n")
			//get user's todo
			fmt.Scan(&data)

			//increment value of ID
			id += 1

			var userData = userData{
				id:   id,
				data: data,
			}

			//add userData to todos list
			todos = append(todos, userData)
			fmt.Printf("List of todos %v\n", todos)
		}
	}
}
