package main

import "fmt"


// array with specifed size will always have the specified length. there is no way to change the size of array.
func main(){
	// [1] How to creatie a variable - long syntax
	var arr [3]int // [size_of_the_array]type

	// assignment per index, indexed from 0 in go
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[1])
	// fmt.Println(arr[3]) // will throw as it's bund controlled

	// [2] implicit initalisation syntax
	arr2 := [3]int{4,5,6} // [size_of_array]type{initialization_values}

	fmt.Println(arr2)
	fmt.Println(arr2[1])
}