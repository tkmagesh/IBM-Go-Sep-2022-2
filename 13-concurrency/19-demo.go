package main

/* Modify the program in such a way that the program keeps generating the numbers until the user hits the ENTER key */

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan string)
	go func() {
		fmt.Println("Hit ENTER to stop....")
		fmt.Scanln()
		stopCh <- "stop"
	}()

	ch := generateNos(stopCh)
	for data := range ch {
		fmt.Println(data)
	}
}

func generateNos(stopCh chan string) <-chan int {
	ch := make(chan int)
	go func() {
		var no int = 1
	LOOP:
		for {
			select {
			case ch <- no * 10:
				time.Sleep(500 * time.Millisecond)
				no++
			case <-stopCh:
				break LOOP
			}
		}
		close(ch)
	}()
	return ch
}
