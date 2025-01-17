// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143?

package main

import "fmt"

// take a number, set the divisor to be 2, as dividing by 1 is not interesting
// if the number divides by the divisor
// set the result to be the divisor
// set the new number to be the original number divided by the divisor
// increase the divisor by 1
// iterate

func sol3(maximum int) (result int) {
	var divisor int = 2
	for maximum > 1 {
		for maximum % divisor == 0 {
			result = divisor
			maximum /= divisor
		}
		divisor++
	}
	return result
}


func main() {
	fmt.Println(sol3(600851475143))
}