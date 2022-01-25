package main

import (
	"os"

	"ascii-art/Functions"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		word := []rune(args[0])
		chars := functions.FileInit("Banners/standard.txt")

		functions.Transform(word, chars)
	}
}