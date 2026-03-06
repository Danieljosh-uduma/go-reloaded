package formatter

import (
	"errors"
	"os"
	"strings"
)

func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", errors.New(filename + "not found.")
	}
	return string(content), nil
}

func WriteFile(filename, text string) {
	os.WriteFile(filename, []byte(text), 0644)
}

func GetFileNames() (string, string, error) {
	if len(os.Args) != 3 {
		return "", "", errors.New("Error: The name of the input and output file must be included")
	}

	return os.Args[1], os.Args[2], nil
}

func ProcessText(data string) string {
	wordSplit := strings.Fields(data)
	editSplit := []string{}

	exchange := func(fn func(string) string) {
		if len(editSplit) < 1 {
			return
		}
		prev := len(editSplit) - 1
		editSplit[prev] = fn(editSplit[prev])
	}

	for _, r := range wordSplit {
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

		editSplit = append(editSplit, r)
	}

	return strings.Join(editSplit, " ")
}
