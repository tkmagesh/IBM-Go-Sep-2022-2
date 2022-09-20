/*
	Display the following menu
		1. Add
		2. Subtract
		3. Multiply
		4. Divide
		5. Exit

	If the user choice is 1 - 4
		accept 2 numbers from the user
		perform the operation
		print the result
		display the menu again

	If the user choice = 5
		print "Thank you"
		exit the application

	If the user choice is NOT 1 - 5
		print "Invalid choice"
		display the menu again
*/
package main

import (
	"errors"
	"fmt"
)

type operation func(int, int) int

var operations = map[int]operation{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {

	for {
		userChoice := getUserChoice()
		if userChoice == 5 {
			break
		}
		result, err := processUserChoice(userChoice)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Result =", result)
	}
	fmt.Println("Thank you!")
}

func getUserChoice() (userChoice int) {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter the choice :")
	fmt.Scanln(&userChoice)
	return
}

func getOperands() (n1 int, n2 int) {
	fmt.Println("Enter the nos:")
	fmt.Scanln(&n1, &n2)
	return
}

func processUserChoice(userChoice int) (result int, err error) {
	if operation, exists := operations[userChoice]; exists {
		n1, n2 := getOperands()
		result = operation(n1, n2)
		return
	}
	err = errors.New("invalid choice")
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x + y
}

func divide(x, y int) int {
	return x + y
}
