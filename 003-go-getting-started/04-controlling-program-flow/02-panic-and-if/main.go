package main

import "fmt"

type User struct {
	ID   int
	name string
}

// OUTPUT with PANIC
// go run .
// starting...
// panic: PANIC!

// goroutine 1 [running]:
// main.main()
//
//	/Users/adrianaoiszak/RubymineProjects/learning-go/003-go-getting-started/04-controlling-program-flow/02-panic-and-if/main.go:8 +0x6c
//
// exit status 2
func main() {
	fmt.Println("starting...")

	// panic("PANIC!")

	fmt.Println("started")

	// If statement
	u1 := User{
		ID:   1,
		name: "A",
	}

	u2 := User{
		ID:   2,
		name: "B",
	}

	if u1.ID == u2.ID {
		fmt.Println("User same")
	} else if u1.name == u2.name {
		fmt.Println("Similar user")
	} else {
		fmt.Println("User different")
	}
}
