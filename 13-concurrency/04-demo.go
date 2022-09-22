package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go f1() //=> handing over this function to the scheduler to be scheduled for execution
	f2()
	wg.Wait()
	fmt.Println("exiting main")
}

func f1() {
	fmt.Println("f1 invocation - started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 invocation - completed")
	wg.Done() //=> decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
