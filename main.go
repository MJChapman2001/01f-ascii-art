package main

import (
	"fmt"
	"os"

	"ascii-art/Functions"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		word := args[0]
		chars := functions.FileInit("Banners/standard.txt")

		result := functions.Transform(word, chars)

		fmt.Print(result)
	}
}