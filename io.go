package formatter

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// GetFileNames checks if the correct number os command-line argument is correct
// Return an Error if it is not correct
// Return the input and output file names if it is correct
func GetFileNames() (string, string, error) {
	if len(os.Args) != 3 {
		return "", "", errors.New("Error: The name of the input and output file must be included")
	}

	return os.Args[1], os.Args[2], nil
}

// ReadFile reads the content of a file
// Return the content of the file if it is correct
// Return an Error if it is not correct
func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", errors.New(filename + "not found.")
	}
	return string(content), nil
}

// WriteFile writes the content to the filename provided.
// If the file does not exist, it creates the file.
func WriteFile(filename, text string) {
	os.WriteFile(filename, []byte(text), 0644)
}

// ProcessText processes the text based on the commands provided.
// It splits the text into words and processes each word based on the commands.
// It returns the processed text.
func ProcessText(data string) string {
	wordSplit := strings.Fields(data)
	editSplit := []string{}

	// exchange helper function applies a function to the last word in editSplit
	exchange := func(fn func(string) string) {
		if len(editSplit) < 1 {
			return
		}
		prev := len(editSplit) - 1
		editSplit[prev] = fn(editSplit[prev])
	}

	// modify helper function applies a function to the last n words in editSplit
	// and removes the last word from editSplit
	modify := func(cmd string, value int) {
		index := len(editSplit) - 2

		for i := value; i > 0; i-- {
			switch cmd {
			case "up":
				editSplit[index] = up(editSplit[index])
			case "cap":
				editSplit[index] = cap(editSplit[index])
			case "low":
				editSplit[index] = low(editSplit[index])
			}
			index--
		}
		editSplit = editSplit[:len(editSplit)-1]
	}

	for idx, r := range wordSplit {
		switch r {
		case "(hex)":
			exchange(hex)
			continue
		case "(bin)":
			exchange(bin)
			continue
		case "(up)":
			exchange(up)
			continue
		case "(low)":
			exchange(low)
			continue
		case "(cap)":
			exchange(cap)
			continue
		}

		// if the word has a suffix of ")", it is a command
		// and it was not seen by the switch, means it is a command
		if strings.HasSuffix(r, ")") {
			word := wordSplit[idx-1]
			num, _ := strconv.Atoi(r[:len(r)-1])
			modify(word[1:len(word)-1], num)
			continue
		}

		editSplit = append(editSplit, r)
	}

	return join(editSplit)
}
