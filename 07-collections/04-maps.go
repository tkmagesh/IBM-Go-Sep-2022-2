package main

import "fmt"

func main() {
	//var productRanks map[string]int = make(map[string]int)
	/*
		productRanks := make(map[string]int)
		productRanks["pen"] = 3
	*/

	/*
		productRanks := map[string]int{}
		productRanks["pen"] = 3
	*/

	//productRanks := map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	productRanks := map[string]int{
		"pen":    3,
		"pencil": 1,
		"marker": 5,
	}
	fmt.Println(productRanks)

	//adding a new item
	productRanks["notebook"] = 4
	fmt.Println(productRanks)

	//checking for the existence of an item
	keyToCheck := "notebook"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is : %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key [%q] does not exist\n", keyToCheck)
	}

	//remove an item
	delete(productRanks, "notebook")
	fmt.Println("After deleting 'notebook'")
	fmt.Println(productRanks)

	//iterating a map
	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println("No of item in the map :", len(productRanks))
}
