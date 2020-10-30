package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filenames := os.Args[1:]

	for _, filename := range filenames {
		fmt.Println("--- File Content ---")
		printFile(filename)
		fmt.Println("--- End File Content ---")
	}
}

func printFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(filename, "not found.")
	} else {
		io.Copy(os.Stdin, file)
		fmt.Println()
	}
}
