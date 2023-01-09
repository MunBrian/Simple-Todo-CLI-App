package main

import (
	"bufio"
	"fmt"
	"os"
)

func test() string {
	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&name)
	// fmt.Println("Hello, ", name)

	//fmt.Scan(&data)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a todo:\n")
	data, _ := reader.ReadString('\n')

	return data
}
