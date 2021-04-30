package main

import "fmt"

func main() {

	// var colors map[string]string

	colors := make(map[int]string)

	// colors := map[string]string{
	// 	"red":  "#523636",
	// 	"blue": "#54555",
	// }

	colors[10] = "ggsgsg"
	delete(colors, 10)
	fmt.Println(colors)
}
