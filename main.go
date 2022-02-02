package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	functions "ascii-art/Functions"
)

func CheckBanner(s string) (map[int]string, error) {
	bannerList := []string{"standard", "shadow", "thinkertoy"}

	for _, i := range bannerList {
		if i == s {
			path := "Banners/" + s + ".txt"

			return functions.FileInit(path), nil
		}
	}

	return nil, errors.New("Banner not found")
}

func main() {
	args := os.Args[1:]

	/* ascii-art
	if len(args) == 1 {
		word := args[0]
		chars := functions.FileInit("Banners/standard.txt")

		result := functions.Transform(word, chars)

		fmt.Print(result)
	}
	*/

	/* fs
	message := "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"

	if len(args) == 2 {
		word := args[0]
		banner := args[1]

		chars, err := CheckBanner(banner)
		if err != nil {
			fmt.Println(message)
			panic(err)
		}

		result := functions.Transform(word, chars)

		fmt.Print(result)
	}
	*/

	message := "Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard --output=<fileName.txt>"

	if len(args) == 3 {
		word := args[0]
		banner := args[1]
		output := args[2]

		chars, err := CheckBanner(banner)
		if err != nil {
			fmt.Println(message)
			panic(err)
		}

		prefix := "--output="

		if strings.HasPrefix(output, prefix) {
			fileName := strings.TrimPrefix(output, prefix)

			result := functions.Transform(word, chars)

			os.WriteFile(fileName, []byte(result), 0666)
		} else {
			fmt.Println(message)
		}
	}
}
