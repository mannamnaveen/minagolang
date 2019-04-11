package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello World!\n")
	if numberBytes, theError := fmt.Printf("Hello World!!! \n"); theError != nil {
		fmt.Printf("Error : %s ", theError)
		os.Exit(1)
	} else {
		fmt.Printf("Number Bytes : %d \n", numberBytes)
	}
}
