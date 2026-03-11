package main

import (
	"fmt"
	"formatter/engine"
)

func main() {
	// Get the input and output file names
	input, output, err := engine.GetFileNames()

	if err != nil {
		fmt.Println(err)
		return
	}

	// read input file
	text, err := engine.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// process the text and return a formatted text
	editedText := engine.ProcessText(text)

	// write the formatted text to the output file
	engine.WriteFile(output, editedText)
}
