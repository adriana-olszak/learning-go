package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Coments contain significant to me comments made by tutor dring the course video

func main() {
	// here we use flag standard lib that allows us to declare what command params
	//  we exepct to receive and also define helper text for them
	// we also define types that we expect to recevie
	path := flag.String("path", "myapp.log", "The path that should be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, ERROR and CRITICAL")

	flag.Parse() // this with actually populate variables with CLI params. if they were not provided we will get default value

	// it will always return an Error as 2nd param instead of throwing the Error
	// we need to use * before path as the return type of flag.String is not an actual string but pointer to it
	f, err := os.Open(*path)

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

		if strings.Contains(s, *level) {
			fmt.Println((s))
		}
	}

}

// [EXAMPLE OF A RUN]
// passing --help will print what we defined using flags
// go run . --help
// Usage of /var/folders/ks/cx852zy94_n63qzdxrc02k1m0000gn/T/go-build655536635/b001/exe/creating-a-cli-app:
//   -level string
//     	Log level to search for. Options are DEBUG, INFO, ERROR and CRITICAL (default "ERROR")
//   -path string
//     	The path that should be analyzed (default "myapp.log")

// [PROBLEMS]
// when running the script uing go run . it failed with expected 'package', found 'EOF
// the issue was simple. I forgot to save. Using VSC very rearly. Have to setup auto-save as I have in Jetbrains IDEs.
