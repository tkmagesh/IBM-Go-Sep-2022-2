package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileReader, err := os.Open("data2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fileReader.Close()

	scanner := bufio.NewScanner(fileReader)
	lineNumber := 0
	for scanner.Scan() {
		txt := scanner.Text()
		lineNumber++
		fmt.Printf("%d : %s\n", lineNumber, txt)
	}
}
