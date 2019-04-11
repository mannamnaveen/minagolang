package main

import "fmt"

func printer(a []string) {
	for _, i := range a {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Arrays and Slices in Go Lang")
	arr := []string{"Hello", "World", "How", "Are", "You"}
	arr = append(arr, "!!!")
	// Copies the slice to new arr2 variable
	// But the memory reference will be same
	arr2 := arr
	arr2[2] = "Porapo"
	arr3 := make([]string, len(arr))
	copy(arr3, arr)
	arr3[2] = "Rara"
	printer(arr2)
	printer(arr)
	printer(arr3)
}

// Arrays are passes as values to function
// A new copy is created in the memory and then assigned to the function

// Slices are passed as reference with the same memory location.
// need to use copy keyword to copy a slice to other varaible
