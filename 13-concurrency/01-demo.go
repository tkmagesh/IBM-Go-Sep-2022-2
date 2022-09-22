package main

import "fmt"

func main() {
	go f1() //=> handing over this function to the scheduler to be scheduled for execution
	f2()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
