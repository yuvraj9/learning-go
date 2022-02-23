package main

import "fmt"

const x int = 10 // declaring typed constant
const (
	name = "hello"
)

func main() {
	// Declaring variables
	var hello bool
	var isAwesome = true
	var yello int

	yello = 20

	const y = "lalas"
	yello = 40

	// Printing variables and constants
	fmt.Println(yello)
	fmt.Println(name)
	fmt.Println(isAwesome)
	fmt.Println(hello)
	fmt.Println(x, y)
}
