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
	result := <-ch //receive operation
	wg.Wait()
	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	result := x + y
	ch <- result //send operation
}
