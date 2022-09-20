package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("	deferred @ main")
	}()
	f1()
}

func f1() {
	/*
		defer func() {
			fmt.Println("	deferred @ f1 [1]")
		}()
		defer func() {
			fmt.Println("	deferred @ f1 [2]")
		}()
		defer func() {
			fmt.Println("	deferred @ f1 [3]")
		}()
	*/

	defer fmt.Println("	deferred @ f1 [1]")
	defer fmt.Println("	deferred @ f1 [2]")
	defer fmt.Println("	deferred @ f1 [3]")
	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
}

func f2() {
	defer func() {
		fmt.Println("	deferred @ f2")
	}()
	fmt.Println("f2 started")
	//fmt.Println(divide(100, 0))
	fmt.Println("f2 completed")
}

func divide(x, y int) int {
	return x / y
}
