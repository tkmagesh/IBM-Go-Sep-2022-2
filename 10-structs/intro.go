package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

func main() {
	//var emp Employee
	//var emp Employee = Employee{100, "Magesh", "Bengaluru"}
	//var emp Employee = Employee{Id: 100, Name: "Magesh"}
	/*
		var emp Employee = Employee{
			Id:   100,
			Name: "Magesh",
			City: "Bengaluru",
		}
	*/
	/*
		var emp = Employee{
			Id:   100,
			Name: "Magesh",
			City: "Bengaluru",
		}
	*/

	emp := Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bengaluru",
	}

	//pointer
	//empPtr := new(Employee)

	empPtr := &Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bengaluru",
	}

	//fmt.Println(emp)
	fmt.Printf("%#v\n", emp)

	fmt.Println("Accessing the fields")
	fmt.Printf("Name = %q\n", emp.Name)

	fmt.Println("Accessing the fileds (pointer)")
	//fmt.Printf("Name = %q\n", (*empPtr).Name)
	fmt.Printf("Name = %q\n", empPtr.Name)

	emp2 := Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bengaluru",
	}
	fmt.Println(emp == emp2)

	employees := []Employee{}
	employees = append(employees, Employee{100, "Magesh", "Bengaluru"})
	employees = append(employees, Employee{101, "Suresh", "Pune"})
	employees = append(employees, Employee{102, "Ramesh", "Chennai"})

	fmt.Println(employees)
}
