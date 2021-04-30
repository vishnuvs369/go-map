package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#523636",
		"blue":  "#54555",
		"white": "#fffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}
}
