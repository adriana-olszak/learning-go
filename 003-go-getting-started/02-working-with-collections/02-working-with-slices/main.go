package main

import "fmt"

func main() {
	// [1] Creating the slice for a array
	arr := [3]int{1, 2, 3} // [size_of_array]type{initialization_values}

	slice := arr[:] // varialbe_name[:] colon makes it create a slice from a collection

	fmt.Println(arr, slice)

	arr[1] = 20
	slice[2] = 40

	// Slice is pointing into the array that is kept in memory
	// so any changes made to the array will be reflected back into the slice
	// and any changes made to the slice will be reflected back into the array
	fmt.Println("afterReAssingment", arr, slice)

	// [2] defining a slice without underlaying array
	slice2 := []int{3, 4, 5} // no size of array specified means that complires will be responsible for manging the underlying array

	fmt.Println(slice2)

	// [3] adding elements to a slice
	slice2 = append(slice2, 4) // override slice2 with a new slice

	fmt.Println(slice2)

	// what happens with size bounded underlaying array?
	// go compiler will fill the array till it's size is exceeded
	// at that stage compiler will create another array in the memory with needed size
	// copy the data and destroy the prev array

	// [4] using : operator with more finesse
	s := slice2[1:]   // it's gonna be a slice created out of slice2 starting from index 1
	s2 := slice2[:2]  // gonna grab first 2 elements (0, 1)
	s3 := slice2[1:2] // gonna take only element at index 1

	fmt.Println(s, s2, s3)

}
