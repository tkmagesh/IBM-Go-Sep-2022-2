/* Higher Order Functions - Assign functions as values to variables */

package main

import "fmt"

type operation func(int, int) int

func main() {
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var greet func() string
	greet = func() string {
		return fmt.Sprintf("Hi %q!", "Magesh")
	}
	fmt.Println(greet())

	//var add func(int, int) int
	var add operation
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	//var subtract func(int, int) int
	var subtract operation
	subtract = func(x, y int) int {
		return x - y
	}
	fmt.Println(subtract(100, 200))
}
