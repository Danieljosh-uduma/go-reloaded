package main

import (
	"fmt"
	"formatter"
	"os"

	// "strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("sample.txt")
	if err != nil {
		panic(err)
	}
	newWords := []string{}
	var builder strings.Builder

	words := strings.Fields(string(content))
	for _, r := range words {
		if r == "(cap)" {
			prev := len(newWords) - 1
			newWords[prev] = formatter.Cap(newWords[prev])
			continue
		} else if r == "(up)" {
			prev := len(newWords) - 1
			newWords[prev] = formatter.Up(newWords[prev])
			continue
		}

		newWords = append(newWords, r)
	}

	for _, r := range newWords {
		builder.WriteString(r)
		builder.WriteString(" ")
	}

	fmt.Println(builder.String())

}
