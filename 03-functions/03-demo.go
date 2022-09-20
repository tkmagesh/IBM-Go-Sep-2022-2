/* Anonymous functions */
package main

import "fmt"

func main() {
	userName := "Magesh"
	func() {
		fmt.Println("fn invoked")
	}()

	msg := func() string {
		return fmt.Sprintf("Hi %q!", userName)
	}()
	fmt.Println(msg)

	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println(result)
}
