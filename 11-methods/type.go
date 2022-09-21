package main

import "fmt"

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	s := MyStr("This is a string")
	fmt.Println(s.Length())
}
