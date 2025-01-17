// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import "fmt"

func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

func LCM(a, b int) int {
	result := a * b / GCD(a, b)
	return result
}

func sol5(maximum int) (result int) {
	result = 1
	for i:=2; i < maximum; i++ {
		result = LCM(result, i)
	}
	return
	}

func main() {
	fmt.Println(sol5(20))
}