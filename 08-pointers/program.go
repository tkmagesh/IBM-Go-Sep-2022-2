package main

import "fmt"

func main() {
	var no int = 10
	var noPtr *int = &no
	fmt.Println(noPtr)

	//dereferencing
	fmt.Println(*noPtr)
	fmt.Println(no == *(&no))

	fmt.Println("Before incrementing, no = ", no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	n1, n2 := 10, 20
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)

	nos = [5]int{3, 1, 4, 2, 5}
	sort(nos)
	fmt.Println(nos)
}

func increment(x *int) {
	*x++
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func sort( /*  */ ) {
	/*  */
}
