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

import "fmt"

func main() {

	for {
		userChoice := getUserChoice()
		if userChoice == 5 {
			break
		}

		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		n1, n2 := getOperands()
		result := processUserChoice(userChoice, n1, n2)
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

func processUserChoice(userChoice, n1, n2 int) int {
	switch userChoice {
	case 1:
		return n1 + n2
	case 2:
		return n1 - n2
	case 3:
		return n1 * n2
	case 4:
		return n1 / n2
	default:
		return 0
	}
}
