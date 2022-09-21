package main

import (
	"fmt"
	"modularity-demo/calculator"
	calc_utils "modularity-demo/calculator/utils"
)

func main() {
	fmt.Println("app executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("OP count =", calculator.OpCount())
	fmt.Println(calc_utils.IsEven(21))
}
