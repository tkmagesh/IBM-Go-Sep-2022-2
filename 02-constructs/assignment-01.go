/* Print all the prime numbers between 3 and 100 */
package main

import "fmt"

func main() {
LOOP:
	for no := 3; no <= 100; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		fmt.Println(no)
	}
}
