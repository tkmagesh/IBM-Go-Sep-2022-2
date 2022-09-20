package main

import "fmt"

func main() {
	//var nos []int
	//var nos []int = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	nos = append(nos, 100)
	fmt.Println(nos)

	nos = append(nos, 200, 300)
	fmt.Println(nos)

	thousands := []int{1000, 2000, 3000}
	nos = append(nos, thousands...)
	fmt.Println(nos)

	fmt.Println("Iterating an array using Indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("Iterating an array using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	/*
		newNos := nos
		newNos[0] = 99
		fmt.Println(nos, newNos)
	*/

	/*
		slicing
		[lo:hi] => from lo to hi-1
		[lo:] => from lo to end of the slice
		[:hi] => from 0 to hi-1
	*/

	fmt.Println("nos[3:6] =", nos[3:6])
	fmt.Println("nos[3:] =", nos[3:])
	fmt.Println("nos[:6] =", nos[:6])

	subset := nos[3:6]
	/*
		subset[0] = 99
		fmt.Println(subset, nos)
	*/

	//creating a copy of the slice
	copyOfSubset := make([]int, len(subset), len(subset))
	copy(copyOfSubset, subset)
	copyOfSubset[0] = -100
	fmt.Println(subset, copyOfSubset)

}
