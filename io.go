package formatter

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func GetFileNames() (string, string, error) {
	if len(os.Args) != 3 {
		return "", "", errors.New("Error: The name of the input and output file must be included")
	}

	return os.Args[1], os.Args[2], nil
}

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
