package main

import (
	"fmt"
	"formatter"
)

func main() {
	// Get the input and output file names
	input, output, err := formatter.GetFileNames()

	if err != nil {
		fmt.Println(err)
		return
	}

	// read input file
	text, err := formatter.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// process the text and return a formatted text
	editedText := formatter.ProcessText(text)

	// write the formatted text to the output file
	formatter.WriteFile(output, editedText)
}
