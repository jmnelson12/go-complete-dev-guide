package main

import "fmt"

/*

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"blue":  "#00FF00",
	// 	"green": "#0000FF",
	// }
	// colors["black"] = "000000"

	// delete(colors, "red")

*/
func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#00FF00",
		"green": "#0000FF",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
