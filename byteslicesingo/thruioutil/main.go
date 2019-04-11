package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// writeToFile will create a file and write data into it.
func writeToFile(filename string) {
	// Use some random data
	message := "Hello world, how are you. \n"

	// Convert the string to byte format
	byteData := []byte(message)

	// Open a file
	// write the byte slice data to file
	err := ioutil.WriteFile(filename, byteData, 0666)

	// Checking for errors in the file
	if err != nil {
		fmt.Println("Write operation not success")
		log.Fatal(err)
		os.Exit(1)
	}
}

// readFromFile reads a file and displays the text
func readFromFile(filename string) {
	// Read the file
	// get the data from file in byte format
	b, err := ioutil.ReadFile(filename)

	// checking for error in opening the file
	if err != nil {
		fmt.Println("Error in reading the file")
		log.Fatal(err)
		os.Exit(1)
	}

	// Conver the byte data into string format
	message := string(b)

	// print the string data
	fmt.Println()
	fmt.Println(message)
}

func main() {
	fmt.Println("Write and Read files...")
	writeToFile("sample.txt")
	readFromFile("text.txt")
}
