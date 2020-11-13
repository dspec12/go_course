package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#348284",
		"green": "#48295",
		"white": "#fffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(hex, "is makes the color", color)
	}
}
