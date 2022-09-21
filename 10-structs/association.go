package main

import "fmt"

type Organization struct {
	Name     string
	Location string
}

type Employee struct {
	Id   int
	Name string
	City string
	Org  *Organization
}

func main() {
	ibm := &Organization{
		Name:     "IBM",
		Location: "Cochin",
	}
	emp := Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bengaluru",
		Org:  ibm,
	}
	fmt.Println(emp)

	emp2 := Employee{
		Id:   101,
		Name: "Suresh",
		City: "Bengaluru",
		Org:  ibm,
	}
	fmt.Println(emp2)

	fmt.Println("After changing the Org Location of emp")
	emp.Org.Location = "Pune"

	fmt.Println(emp.Org.Location)
	fmt.Println(emp2.Org.Location)
}
