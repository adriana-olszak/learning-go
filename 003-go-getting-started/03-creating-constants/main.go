package main

import "fmt"

func main() {
	// [1] Constant declaration
	// we need to initialize them at the point of declaration
	// constants cannot be reassigned
	// value of the constant HAS TO BE ABLE TO DETERMINED at COMPILE time
	// for example return value of a function cannot be used as value of contant as functions are only interpreted on run time
	const c = 3

	fmt.Println(c)

	// [2] Datatype will be interpreted on each line.
	// to limit it we would need to assign a specific type during initialisation
	fmt.Println(c + 3)
	fmt.Println(c + 1.2)

	// [3] Constant with specified data type can be casted to another datatypes

	const d int = 4
	fmt.Println(d)

	fmt.Println(float32(c) + 1.2)
}
