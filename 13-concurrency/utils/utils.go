package utils

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Fn() {
	fmt.Println("fn started")
	wg.Add(3)
	go f1()
	go f2()
	go f3()
	wg.Wait()
	fmt.Println("fn completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done()
}

func f2() {
	fmt.Println("f2 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f2 completed")
	wg.Done()
}

func f3() {
	fmt.Println("f3 started")
	time.Sleep(2 * time.Second)
	fmt.Println("f3 completed")
	wg.Done()
}
