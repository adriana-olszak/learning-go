package main

import "fmt"

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{Method: "GET"}

	// by default in GO WE DO NOT Fall through the cases 
	// to allow it to fallthough keyword
	switch r.Method {
	case "GET":
		fmt.Println("GET request")
	case "PUT":
		fmt.Println("PUT request")
	case "POST":
		fmt.Println("POST request")
	case "DELETE":
		fmt.Println("DELETE request")
	default:
		fmt.Println("Unhandled method")
	}
}
