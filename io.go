package formatter

import "os"

func ReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func WriteFile(filename, text string) {
	// os.WriteFile()
}
