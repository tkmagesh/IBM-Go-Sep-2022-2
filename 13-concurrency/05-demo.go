package main

import (
	"concurrency-demo/utils"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	wg.Add(1)
	go func() {
		utils.Fn()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("main completed")
}
