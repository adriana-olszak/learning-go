package main

import "fmt"

func main() {
	// [1] defining a Pointer Data type
	var firstName *string
	fmt.Println(firstName) // outputs <nil> nil is a representation to a pointer which do not referenece to anything

	// [2] using a dereference operator
	// *firstName = "Arthur"
	// this will throw an error when run as we invalid memory address
	// cause we are trying to assing String "Arthur" to uninitialized pointer

	// [3] initialisation of the pointer to a string, using new() allocates memory for aspecified data type
	var firstNameInitialized *string = new(string)
	*firstNameInitialized = "Arthur"

	fmt.Println(firstNameInitialized)  // prints the memory address
	fmt.Println(*firstNameInitialized) // using dereference to access it's value

	// pointer arythmetic in form of *(firstName +1) was removed from the language.
	// pointer arythmetics is available in a different form

	// [4] Create a pointer that points to a variable
	firstNamePoint := "Arthur"
	fmt.Println(firstNamePoint)

	// address of operator - we can get the Pointer the the memory
	ptr := &firstNamePoint

	fmt.Println(ptr, *ptr)

	firstNamePoint = "Tricia"
	fmt.Println(ptr, *ptr)

	// the memory location will be always the same even though we change value that it points to it
}
