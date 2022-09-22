package main

import (
	"fmt"
	"time"
)

//Share memory by communicating

func main() {
	ch := make(chan int)
	fmt.Println("main started")
	go add(100, 200, ch)
	result := <-ch //receive operation
	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int, ch chan int) {
	time.Sleep(3 * time.Second)
	result := x + y
	ch <- result //send operation
}
