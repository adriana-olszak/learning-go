package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Coments contain significant to me comments made by tutor dring the course video

func main() {
	f, err := os.Open("myapp.log") // it will always return an Error instead of throwing the Error

	if err != nil { // if we have a pointer empty value it will equal nil
		log.Fatal(err)
	}

	defer f.Close() // defer will be run after the function it's called in exits it will clean up
	// defer can be run immidietyly after the successufll initialization which makes it easier to rembember

	r := bufio.NewReader(f) // buffer intpu output package. It will setu a reader any input String n this case it's a file

	for { // with no params it will create inifinie foop
		s, err := r.ReadString('\n')

		if err != nil {
			break
		}

		if strings.Contains(s, "ERROR") {
			fmt.Println((s))
		}
	}

}

// NEWBE problems
// when running the script uing go run . it failed with expected 'package', found 'EOF
