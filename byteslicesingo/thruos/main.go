package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Byte Slices in Golang...")
	// Reading the file into a variable
	// storing the buffer into f and error into e
	f, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("Error opening the file")
	} else {
		// Printing the address of the f
		fmt.Println(f)
	}

	b := make([]byte, 200)

	_, err = f.Read(b)
	stringValue := string(b)

	fmt.Printf("%s \n", stringValue)
}
