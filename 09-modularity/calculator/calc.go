package calculator

import "fmt"

var operationCount int

func init() {
	//any initialization logic
	fmt.Println("calculator package initialized [1]")
}

func init() {
	//any initialization logic
	fmt.Println("calculator package initialized [2]")
}

func OpCount() int {
	return operationCount
}
