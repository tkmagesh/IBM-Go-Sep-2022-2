package main

import "fmt"

//var x int
//var x int = 100
//var x = 100
//x := 100 //=> NOT applicable at the package level

var myVar int // => unused variables are allowed at the package level

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	/*
		//type inteference
		var msg = "Hello World!"
	*/

	//the below syntax is application ONLY in a function (function scope) and NOT at the package level
	msg := "Hello World"
	fmt.Println(msg)

	//unused variables are NOT allowed at the function level
	var myVar int
	myVar = 100
	fmt.Println(myVar)

	/* multiple variables */
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of 100 and 200 is"
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of 100 and 200 is"
			result int    = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "Sum of 100 and 200 is"
			result = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of 100 and 200 is"
			result    = x + y
		)
		fmt.Println(str, result)
	*/

	x, y, str := 100, 200, "Sum of 100 and 200 is"
	result := x + y
	fmt.Println(str, result)
}
