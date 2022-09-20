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
}
