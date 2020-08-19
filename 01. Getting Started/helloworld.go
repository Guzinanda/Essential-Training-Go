/*
* DATASHEET
* @file  		helloworld.go
* @brief 		Traditional Hello World
* @details
* @author  		Fernanda Guzmán Gómez
* @date  		18/08/2020
* @copyright Copyright (c) 2020 Fernanda Guzmán Gómez. All rights reserved.
 */

package main

import (
	"fmt" // fmt: Package of functions for formatted printing.
)

func main() {
	fmt.Println("Hello World ☺") // Go strngs are Unicode, you don´t need special code for non-english languages
}

// Run the program: go run helloworld.go
// Build an executable: go build helloworld.go

// Help: go help
