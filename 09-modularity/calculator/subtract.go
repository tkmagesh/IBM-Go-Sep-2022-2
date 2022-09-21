package calculator

import "fmt"

func init() {
	//any initialization logic
	fmt.Println("calculator package initialized [subtract]")
}

func Subtract(x, y int) int {
	operationCount++
	return x - y
}
