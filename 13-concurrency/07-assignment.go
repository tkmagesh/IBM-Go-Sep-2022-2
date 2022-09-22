/*
	spawn 100 goroutines that are blocked (time.Sleep) for a random time period of 0 - 10 seconds
	Wait for all the goroutines to complete and exit the application
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go fn(i, time.Duration(rand.Intn(10))*time.Second, wg)
	}
	wg.Wait()
	fmt.Println("Done")
}

func fn(idx int, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn[%d] started\n", idx)
	time.Sleep(duration)
	fmt.Printf("fn[%d] completed\n", idx)
}
