package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go f1(wg) //=> handing over this function to the scheduler to be scheduled for execution
	f2()
	wg.Wait()
	fmt.Println("exiting main")
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() //=> decrement the counter by 1
	fmt.Println("f1 invocation - started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 invocation - completed")

}

func f2() {
	fmt.Println("f2 invoked")
}
