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
	fmt.Println(divide(100, 7))

}

//3rd party
func divide(x, y int) int {
	return x / y
}
