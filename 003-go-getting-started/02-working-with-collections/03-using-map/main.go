package main

import "fmt"

// with maps we can use arbitraty values as keys
func main() {
	m := map[string]int{"foo": 1, "bar": 2} // map[key_type]value_type{initialisation_comma_delimited}

	fmt.Println(m)
	fmt.Println("foo")

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}
