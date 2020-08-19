// Example of "for loop"

package main

import (
	"fmt"
)

func main() {

	// For Loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// Break to finish loop early
	fmt.Println("----")
	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	// Continue for not run a part
	fmt.Println("----")
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	// While loop
	fmt.Println("----")
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	// While True loop
	fmt.Println("----")
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
}
