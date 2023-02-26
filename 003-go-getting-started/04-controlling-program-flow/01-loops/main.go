package main

import "fmt"

func main() {
	// Types of functions
	// All functions in go are created using for keywod

	var a int // initial value is always gonna be 0
	// [1] till condition met
	fmt.Println("initial value of a", a)
	for a < 5 {
		fmt.Println("simple while loop", a)

		if a == 3 { // you can use continue to omit run of a loop
			a++
			continue
		}

		fmt.Println("simple while loop", "continuing...")
		a++
	}

	// [2] conditional with post clause
	// when using implicic initalization we dont have acccess to the variable outside of the scope of the function
	for b := 0; b < 4; b++ {
		fmt.Println("conditional with post clause", b)
	}

	// [3] Infinite loop
	var c int
	for {
		fmt.Println("infinite loop", c)
		if c == 4 { // also break keyword can be used to break out of the loop
			fmt.Println("breaking out")
			break
		}
		c++

	}

	// [4] Looping over collections
	slice := []int{1, 2, 3}

	// completelty valid syntax
	for i := 0; i < len(slice); i++ {
		fmt.Println("iterate through collection using simple syntax", slice[i])
	}

	// improved syntax to iterate through collections returns index and value
	for i, v := range slice {
		fmt.Println("index and value", i, v)
	}

	// improved syntax to iterate through collections
	m := map[string]int{"http": 80, "https": 443}
	for k, v := range m {
		fmt.Println("looping through map", k, v)
	}

}
