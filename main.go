package main

import (
	"errors"
	"fmt"
	"os"

	functions "ascii-art/Functions"
)

// define webpage structure

type Page struct {
	Title string
	Body  []byte
}

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
	// args := os.Args[1:]

	// /* ascii-art
	// if len(args) == 1 {
	// 	word := args[0]
	// 	chars := functions.FileInit("Banners/standard.txt")

	// 	result := functions.Transform(word, chars)

	// 	fmt.Print(result)
	// }
	// */

	// /* fs
	// message := "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"

	// if len(args) == 2 {
	// 	word := args[0]
	// 	banner := args[1]

	// 	chars, err := CheckBanner(banner)
	// 	if err != nil {
	// 		fmt.Println(message)
	// 		panic(err)
	// 	}

	// 	result := functions.Transform(word, chars)

	// 	fmt.Print(result)
	// }
	// */

	// message := "Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard --output=<fileName.txt>"

	// if len(args) == 3 {
	// 	word := args[0]
	// 	banner := args[1]
	// 	output := args[2]

	// 	chars, err := CheckBanner(banner)
	// 	if err != nil {
	// 		fmt.Println(message)
	// 		panic(err)
	// 	}

	// 	prefix := "--output="

	// 	if strings.HasPrefix(output, prefix) {
	// 		fileName := strings.TrimPrefix(output, prefix)

	// 		result := functions.Transform(word, chars)

	// 		os.WriteFile(fileName, []byte(result), 0666)
	// 	} else {
	// 		fmt.Println(message)
	// 	}
	// }

	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}

/*
Saves Page's Body to a text file
This returns an error value beacuse WriteFile has a return type of error. If there's no error then ir returns nil
0600 indiactes file should be created wiht read-write permissions for current user
*/
func (p *Page) save() error {
	filename := p.Title + ".text"
	return os.WriteFile(filename, p.Body, 0600)
}

/*
Construct file name from title parameter, reads file content into a new variable body and returns pointer to Page literal constructed with proper title and body values
*/
func loadPage(title string) (*Page, error) {
	filename := title + ".text"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
