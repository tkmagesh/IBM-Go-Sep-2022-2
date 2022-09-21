package calculator

import "fmt"

func init() {
	//any initialization logic
	fmt.Println("calculator package initialized [add.go]")
}

func Add(x, y int) int {
	operationCount++
	return x + y
}
