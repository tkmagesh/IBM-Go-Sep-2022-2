package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileReader, err := os.Open("data1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fileReader.Close()
	bufferedReader := bufio.NewReader(fileReader)
	for {
		sentence, err := bufferedReader.ReadString('.')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(sentence)
	}
}
