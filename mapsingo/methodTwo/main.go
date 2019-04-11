package main

import "fmt"

func main() {
	// One way of declaring an empty map with name colors
	// using var keyword
	var colors = make(map[string]string)

	// One way of declaring an empty map with name colors
	// using the literal way
	// colors := make(map[string]string)

	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"

	fmt.Println(colors)
}
