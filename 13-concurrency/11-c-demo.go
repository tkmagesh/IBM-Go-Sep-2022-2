package main

import (
	"fmt"
	"time"
)

//Share memory by communicating

func main() { //consumer
	fmt.Println("main started")
	ch := add(100, 200)
	/*
		go func() {
			ch <- 1000
		}()
	*/
	result := <-ch //receive operation
	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int) <-chan int { //producer
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		result := x + y
		ch <- result //send operation
	}()
	return ch
}
