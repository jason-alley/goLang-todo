package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var taskItems = []string{
	"Watch GoLang crash course.",
	"Build SSG project",
	"Reward myself"}

var userOption string

func main() {
	mainProgram()
}

/*
The mainProgram function executes the main logic of the todo list.
*/
func mainProgram() {
	printTasks(taskItems)
	fmt.Println()

	fmt.Println("Here are all the items on your todo list. What would you like to do next? add a new item or remove and existing one ? (add, delete)")

	fmt.Scanln(&userOption)

	switch userOption {
	case "add":
		for {
			fmt.Println("Please enter a new item. Type 'end' to exit!:")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')

			// Trim whitespace and newline chars.
			text = strings.TrimSpace(text)
			if "end" == text {
				mainProgram()
			}
			taskItems = addTask(taskItems, text)
			printTasks(taskItems)

		}
	case "delete":
		for {
			fmt.Println("Please enter the number of the item you wish to remove. Type 'end' to exit!:")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')

			// Trim whitespace and newline chars.
			input = strings.TrimSpace(input)

			// Convert the string to and int.
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Please enter a valid number. Type 'end' to exit!:")
			}
			if "end" == input {
				mainProgram()
			}
			taskItems = removeTask(taskItems, num)
			printTasks(taskItems)
		}
	}
}

/*
printTask function takes in a slice containing strings, and prints them to the user.
*/
func printTasks(taskItems []string) {
	fmt.Println("######## List of Tasks ########")
	for index, task := range taskItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

/*
addTask function takes the existing taskItems slice, and them add the new task item to it.
*/
func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}

/*
remmoveTask function removes a task from slice, based on the number of the task entered.
*/
func removeTask(taskItems []string, itemNumber int) []string {
	var updatedItemItems = append(taskItems[:itemNumber-1], taskItems[itemNumber:]...)
	return updatedItemItems
}
