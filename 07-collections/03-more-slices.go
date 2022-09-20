package main

import "fmt"

func main() {

	/*
		var nos []int
		//var nos []int = []int{3, 1, 4, 2, 5}
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

		nos = append(nos, 10)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

		nos = append(nos, 20)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

		nos = append(nos, 30)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

		nos = append(nos, 40)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

		nos = append(nos, 50)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
	*/
	/*
		nos := make([]int, 5)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
	*/

	nos := make([]int, 5, 10) // memory allocated for 10 int values of which 5 are initialized with 0
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

	//indexer is used to access the memory that is already initialized
	nos[0] = 10
	nos[1] = 20
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

	//append() is used to utilize the uninitialized (may be allocated) memory
	nos = append(nos, 100)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
}
