// FizzBuzz Challenge

package main

import (
	"fmt"
)

func main() {

	// For every number from 1 to 20
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			// if the number is divisible by 3 and 5 print fizz buzz
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			// else if the number is divisible by 3 print fizz
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			// else if the number is divisible by 5 print buzz
			fmt.Println("Buzz")
		} else {
			// else print the number
			fmt.Println(i)
		}
	}
}
