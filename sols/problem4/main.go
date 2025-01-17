// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 x 99
// Find the largest palindrome made from the product of two 3-digit numbers

package main

import (
	"fmt";
	"math"
)

func is_palindrome(i int64) (result bool) {
	if i < 0 {
        return false
    }
	var original = i
	var reversed int64 = 0
	for i != 0 {
		reversed = reversed*10 + i%10
		i /= 10
	}
	return original == reversed

}

func sol4(num_digits float64) (result int64) {
	max_input := math.Pow(10, num_digits) - 1
	min_input := math.Pow(10, num_digits-1)

	for i := max_input; i > min_input; i-- {
		for j := max_input; j > min_input; j-- {
			product := int64(i * j)
			if is_palindrome(product) && (product > result) {
				result = product
				fmt.Println(i)
				fmt.Println(j)
			}
		}
	}
	return
	}


func main() {
	//fmt.Println(is_palindrome(437734))
	fmt.Println(sol4(4))
}