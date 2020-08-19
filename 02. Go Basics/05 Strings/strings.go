// Go Strings

package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"

	// Will print the string ("The colour of magic")
	fmt.Println(book)

	// Will print the lenght (19)
	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// uint8 is a byte

	// Strings in Go are immutable
	// book[0] = 116

	// Will slice (start, end), 0 based, ½ empty range
	fmt.Println(book[4:11])

	// Will slice (no end)
	fmt.Println(book[4:])

	// Will slice (no start)
	fmt.Println(book[:4])

	// Use + to concatenate strings
	fmt.Println("t" + book[1:])

	// Unicode
	fmt.Println("It was ½ price!")

	// Multi line
	poem := `
	The road goes ever on
	Down from the door where it began
	...
	`
	fmt.Println(poem)
}
