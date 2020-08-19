/*

How many times each word appears in a text

Print out how many times each word appears in the text use strings.Fields
to split to words and strings.ToLower to convert to lowercase

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	words := strings.Fields(text)
	counts := map[string]int{} // empty map
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println("\n")
	fmt.Println(counts)
	fmt.Println("\n")

}
