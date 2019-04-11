package main

import "fmt"

func printer(msgs ...string) {
	for _, msg := range msgs {
		fmt.Println(msg)
	}
}

func main() {
	fmt.Println("Functions in Go...")
	printer("Hello ", "World!!")
}
