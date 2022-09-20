/*
Write a program that will print the prime numbers between start & end
accept the start & end from the user

refactor the program to encapsulate the prime number generation logic in a function accepting start and end as parameters
the main function should be responsible for accpeting inputs and printing the result
*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the start and end :")
	fmt.Scanln(&start, &end)
	primes := generatePrimes(start, end)
	fmt.Println(primes)
}

func generatePrimes(start, end int) []int {
	primes := []int{}
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
