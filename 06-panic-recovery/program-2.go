package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Something went wrong!")
			fmt.Println(err)
			return
		}
		fmt.Println("Thank you!")
	}()
	result, err := divideClient(100, 0)
	if err != nil {
		fmt.Println("Do not attempt to divide by zero")
		return
	}
	fmt.Println("result =", result)
}

func divideClient(x, y int) (result int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	result = divide(x, y)
	return
}

//3rd party
func divide(x, y int) int {
	return x / y
}
