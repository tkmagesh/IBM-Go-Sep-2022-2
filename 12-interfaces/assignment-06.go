package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

type ProductPredicate func(Product) bool

func (products Products) IndexOf(p Product) int {
	for idx, product := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products Products) Print() {
	for _, product := range products {
		fmt.Println(product.Format())
	}
}

func (products Products) Filter(predicate ProductPredicate) Products {
	result := Products{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) All(predicate ProductPredicate) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

func (products Products) Any(predicate ProductPredicate) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

/*
	Write the apis for the following
		Sort => Sort the products collection by any attribute
			IMPORTANT : MUST Use sort.Sort() function
            use cases:
                1. sort the products collection by cost
                2. sort the products collection by name
                3. sort the products collection by units
                4. sort the products collection by cost in descending order
                5. sort the products collection by name in descending order
                6. sort the products collection by units in descending order
				etc

*/

//sort.Interface interface implementation
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

//sort by Name
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

//sort by Cost
type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

//sort by Units
type ByUnits struct {
	Products
}

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

//sort by Category
type ByCategory struct {
	Products
}

func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

//High level method for products
func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Units":
		sort.Sort(ByUnits{products})
	case "Cost":
		sort.Sort(ByCost{products})
	case "Category":
		sort.Sort(ByCategory{products})
	}
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	products.Print()

	fmt.Println()
	fmt.Println("Default Sort [By Id]")
	//sort.Sort(products)
	products.Sort("Id")
	products.Print()

	fmt.Println()
	fmt.Println("Default Sort [By Name]")
	//sort.Sort(ByName{products})
	products.Sort("Name")
	products.Print()

	fmt.Println()
	fmt.Println("Default Sort [By Cost]")
	//sort.Sort(ByCost{products})
	products.Sort("Cost")
	products.Print()

	fmt.Println()
	fmt.Println("Default Sort [By Units]")
	//sort.Sort(ByUnits{products})
	products.Sort("Units")
	products.Print()

	fmt.Println()
	fmt.Println("Default Sort [By Category]")
	//sort.Sort(ByCategory{products})
	products.Sort("Category")
	products.Print()

	/*
		fmt.Println("Index Of Stove : ")
		stove := Product{102, "Stove", 5000, 5, "Utencil"}
		fmt.Println(products.IndexOf(stove))

		fmt.Println()
		fmt.Println("Filter")
		fmt.Println("Costly products [cost >= 1000]")
		costlyProductPredicate := func(p Product) bool {
			return p.Cost >= 1000
		}
		costlyProducts := products.Filter(costlyProductPredicate)
		costlyProducts.Print()

		fmt.Println()
		fmt.Println("Stationary Products")
		stationaryProductPredicate := func(p Product) bool {
			return p.Category == "Stationary"
		}
		stationaryProducts := products.Filter(stationaryProductPredicate)
		stationaryProducts.Print()

		fmt.Println()
		fmt.Println("All")
		fmt.Println("Are all the products costly products ? :", products.All(costlyProductPredicate))

		fmt.Println()
		fmt.Println("Any")
		fmt.Println("Are there any costly product ? :", products.Any(costlyProductPredicate))
	*/
}
