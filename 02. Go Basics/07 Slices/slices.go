package main

import (
	"fmt"
)

func main() {
	// Same type
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	fmt.Println("----")
	// Length
	fmt.Println(len(loons)) // 3

	fmt.Println("----")
	// 0 indexing
	fmt.Println(loons[1]) // daffy

	fmt.Println("----")
	// slices
	fmt.Println(loons[1:]) // [daffy taz]

	// Iterate in slices ________________________________________________
	fmt.Println("----")
	for i := 0; i < len(loons); i++ {
		fmt.Println(i)
	}

	fmt.Println("----")
	// single value range
	for i := range loons {
		fmt.Println(loons[i])
	}

	fmt.Println("----")
	// double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("----")
	// double value range, ignore index by using _
	for _, name := range loons {
		fmt.Println(name)
	}

}
