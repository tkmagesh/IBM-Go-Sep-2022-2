package main

import (
	"fmt"
	"sync"
)

//Communicate by Sharing Memory (Not advisable)
var result int

func main() {
	var wg sync.WaitGroup
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200, &wg)
	wg.Wait()
	fmt.Println("result =", result)
	fmt.Println("main completed")
}

func add(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	result = x + y
}
