package main

import (
	"fmt"
	"time"
)

func main() {

	ch := generateNos()
	for data := range ch {
		fmt.Println(data)
	}
}

func generateNos() <-chan int {
	ch := make(chan int)
	timeoutCh := time.After(10 * time.Second)
	go func() {
		var no int = 1
	LOOP:
		for {
			select {
			case ch <- no * 10:
				time.Sleep(500 * time.Millisecond)
				no++
			case <-timeoutCh:
				fmt.Println("Timeout tiggered")
				break LOOP
			}
		}
		close(ch)
	}()
	return ch
}

/*
//using time.After instead

func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
*/
