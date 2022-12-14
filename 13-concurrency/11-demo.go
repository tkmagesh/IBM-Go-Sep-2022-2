package main

import (
	"fmt"
	"sync"
)

//Share memory by communicating

func main() {
	var wg sync.WaitGroup
	//var ch chan int = make(chan int)
	ch := make(chan int)
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200, &wg, ch)
	wg.Wait()
	result := <-ch //receive operation
	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	ch <- result //send operation
	defer wg.Done()
}
