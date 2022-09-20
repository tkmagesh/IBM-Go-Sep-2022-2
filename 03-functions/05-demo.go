/* Higher Order Functions - Pass functions as arguments to other functions */

package main

import (
	"fmt"
	"time"
)

type operation func(int, int) int

func main() {
	exec(fn)

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(100, 200, add)
	logOperation(100, 200, subtract)

	profileOperation(100, 200, add)
	profileOperation(100, 200, subtract)
}

func profileOperation(x, y int, oper operation) {
	start := time.Now()
	fmt.Println(oper(x, y))
	elapsed := time.Now().Sub(start)
	fmt.Println("elapsed = ", elapsed)
}

/*

 */
func logOperation(x, y int, oper operation) {
	fmt.Println("Invocation started")
	fmt.Println(oper(x, y))
	fmt.Println("Invocation completed")
}

/*
func logAdd(x, y int) {
	fmt.Println("Invocation started")
	fmt.Println(add(x, y))
	fmt.Println("Invocation completed")
}

func logSubtract(x, y int) {
	fmt.Println("Invocation started")
	fmt.Println(subtract(x, y))
	fmt.Println("Invocation completed")
}
*/
//3rd party library
func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func exec(f func()) {
	f()
}

func fn() {
	fmt.Println("fn invoked")
}
