package main

import "fmt"

// [OUTPUT]
// go run .
// 42
// 3.14
// 3.14
// Arthur
// true
// (3+4i)
// 3
// 4
func main() {
	// [1] example of how this is done in other languages
	var i int // keyword | name of variable | type
	i = 42    // assignment
	fmt.Println(i)

	// [2] simpler declaration syntax
	var f float32 = 3.14 // we also have float64

	fmt.Println(f)

	// [3] implicit initialization syntax
	t := 3.14 // this will use implicit initialization syntax, it uses float64
	fmt.Println(t)

	// [4] String type
	firstName := "Arthur"
	fmt.Println(firstName)

	// [5] Boolean
	b := true
	fmt.Println(b)

	// [6] Complex datatype, very usefull to do complex math
	c := complex(3, 4)
	fmt.Println(c)

	// [7] Pulling out real and imaginary parts of complex type
	r, im := real(c), imag(c)

	fmt.Println(r)
	fmt.Println(im)
}
