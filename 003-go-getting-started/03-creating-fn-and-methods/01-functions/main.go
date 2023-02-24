package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	startWebServer(port, 3)
	_, err := startWebServerTuple(port, 3) // underscore is write only operator, it allows us to drop some part of return value
	fmt.Println(err)
}

func startWebServer(port int, numberOfRetries int) error {
	fmt.Println("Starting the server...")

	fmt.Println("The server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	// return nil
	return errors.New("something went wrong")
}

// nothing interesting happened in this function, very basic
func startWebServerTuple(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting the server...")

	fmt.Println("The server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	// return nil
	return port, errors.New("something went wrong")
}
