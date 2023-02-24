package main

import "fmt"

// struct is the only collection type that allows us to associate any data type together
// fields that we expose have to be fixed at the compile time
func main() {
	// struct definition
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	// initialisation of a struct
	var u user
	fmt.Println(u) // output: {0  }
	// reason for such output is that struct does exsits. when we initialized the u variable we get them initilized with their zero value
	// for int this value is 0
	// for string it's an empty string

	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)           // {1 Arthur Dent}
	fmt.Println(u.FirstName) // Arthur

	// [2] shorthand initialisation of a struct
	u2 := user{ID: 2, FirstName: "Arthur2", LastName: "Dent2"}

	fmt.Println(u2)           // {2 Arthur2 Dent2}

}
