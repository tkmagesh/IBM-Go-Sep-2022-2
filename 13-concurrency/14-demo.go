package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch := generateNos()
	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
}

func generateNos() <-chan int {
	ch := make(chan int)
	go func() {
		rand.Seed(10)
		count := rand.Intn(20)
		fmt.Printf("Generating %d values\n", count)
		for i := 1; i <= count; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
		close(ch)
	}()
	return ch
}
