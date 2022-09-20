/* Higher Order Functions - Functions can be returned as return result from other functions */

package main

import (
	"fmt"
	"time"
)

type operation func(int, int) int

func main() {
	/*
		fn := getFn()
		fn()
	*/

	/*
		fmt.Println(add(100, 200))
		fmt.Println(subtract(100, 200))
	*/
	/*
		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/
	/*
		logAdd := getLogOperation(add)
		fmt.Println(logAdd(100, 200))

		logSubtract := getLogOperation(subtract)
		fmt.Println(logSubtract(100, 200))

		profiledAdd := getProfiledOperation(add)
		fmt.Println(profiledAdd(100, 200))
	*/

	composedAdd := getProfiledOperation(getLogOperation(add))
	fmt.Println(composedAdd(100, 200))
}

func getLogOperation(oper operation) operation {
	return func(x, y int) int {
		fmt.Println("Invocation started")
		result := oper(x, y)
		fmt.Println("Invocation completed")
		return result
	}
}

func getProfiledOperation(oper operation) operation {
	return func(x, y int) int {
		start := time.Now()
		result := oper(x, y)
		elapsed := time.Now().Sub(start)
		fmt.Println("time taken :", elapsed)
		return result
	}
}

//3rd party library
func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func getFn() func() {
	fn := func() {
		fmt.Println("fn invoked")
	}
	return fn
}
