package main

import (
	"fmt"
	"formatter"
)

func main() {
	input, output, err := formatter.GetFileNames()

	if err != nil {
		fmt.Println(err)
		return
	}
	text, err := formatter.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	editedText := formatter.ProcessText(text)

	formatter.WriteFile(output, editedText)
	fmt.Println(editedText)
}
