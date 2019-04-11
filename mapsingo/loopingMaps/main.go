package main

import "fmt"

func main() {
	fmt.Println("Looping over the items in Maps")
	colors := map[string]string{
		"green": "#00ff00",
		"red":   "#ff0000",
		"blue":  "#0000ff",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color + "\t:\t" + hex)
	}
}
