package main

import (
	"fmt"
	"sync"
)

var counter int

var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go fn(i, wg)
	}
	wg.Wait()
	fmt.Printf("Counter = %d\n", counter)
	fmt.Println("Done")
}

func fn(idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	{
		counter++
	}
	mutex.Unlock()
}
