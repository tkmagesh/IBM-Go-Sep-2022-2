package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan string)

	go func() {
		//time.Sleep(5 * time.Second)
		ch1 <- 10
	}()

	go func() {
		//time.Sleep(3 * time.Second)
		ch2 <- 20
	}()

	go func() {
		//time.Sleep(4 * time.Second)
		fmt.Println(<-ch3)
	}()

	for i := 0; i < 3; i++ {
		select {
		case data1 := <-ch1:
			fmt.Println(data1)
		case data2 := <-ch2:
			fmt.Println(data2)
		case ch3 <- "from select":
			fmt.Println("Sent data to ch3")
			/* default:
			fmt.Println("No channel operations are successful") */
		}
	}
}
