package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{
			Id:       id,
			Name:     name,
			Cost:     cost,
			Units:    units,
			Category: category,
		},
		Expiry: expiry,
	}
}
func main() {
	//grapes := PerishableProduct{Product{101, "Grapes", 50, 5, "Fruits"}, "2 Days"}
	/*
		grapes := PerishableProduct{
			Product: Product{
				Id:       101,
				Name:     "Grapes",
				Cost:     50,
				Units:    5,
				Category: "Fruits"},
			Expiry: "2 Days",
		}
	*/
	grapes := NewPerishableProduct(101, "Grapes", 50, 5, "Fruits", "2 Days")
	fmt.Println(grapes)

	fmt.Println("Accessing the members of the base type")
	fmt.Println(grapes.Product.Name)
	fmt.Println(grapes.Name)
}
