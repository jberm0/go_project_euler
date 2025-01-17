// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23
// Find the sum of all multiples of 3 or 5 below 1000.

package main

import "fmt"

func sol1(maximum int) int {
	var result int = 0
	for i := 1; i < maximum; i++ {
		if (i % 3 == 0) || (i % 5 == 0) {
			result += i
		}
	}
	return result
}

func main() {
	fmt.Println(sol1(1000))
}
