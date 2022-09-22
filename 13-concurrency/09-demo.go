package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	counter int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.counter++
	}
	c.Unlock()
}

var c Counter

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go fn(i, wg)
	}
	wg.Wait()
	fmt.Printf("Counter = %d\n", c.counter)
	fmt.Println("Done")
}

func fn(idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.Increment()
}
