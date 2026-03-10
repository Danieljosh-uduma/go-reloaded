package formatter

import (
	"strconv"
	"strings"
	"unicode"
)

// hex function convert a string in hexadecimal to decimal
// return the decimal as a string
// if the string is not a valid hexadecimal, return the string as it is
func hex(s string) string {
	decimal, err := strconv.ParseInt(s, 16, 64)

	if err != nil {
		return s
	}

	return strconv.Itoa(int(decimal))
}

// bin function convert a string in binary to decimal
// return the decimal as a string
// if the string is not a valid binary, return the string as it is
func bin(s string) string {
	decimal, err := strconv.ParseInt(s, 2, 64)

	if err != nil {
		return s
	}

	return strconv.Itoa(int(decimal))
}

// up function convert a string to uppercase
// return the uppercase as a string
func up(s string) string {
	return strings.ToUpper(s)
}

// low function convert a string to lowercase
// return the lowercase as a string
func low(s string) string {
	return strings.ToLower(s)
}

// cap function convert a string to capitalize
// return the capitalize as a string
func cap(s string) string {
	var builder strings.Builder
	runes := []rune(s)

	if runes[0] >= 'a' && runes[0] <= 'z' {
		runes[0] -= 32
	}
	for idx, r := range runes {
		if r >= 'A' && r <= 'Z' && idx > 0 {
			builder.WriteRune(r + 32)
			continue
		}
		builder.WriteRune(r)
	}
	return builder.String()
}

// isPunc function check if a string is a punctuation
// return true if the string is a punctuation, false otherwise
func isPunc(s string) bool {
	for _, r := range s {
		if !unicode.IsPunct(r) {
			return false
		}
	}

	return true
}

// join function create a string from a slice of strings
// return the string with spaces between words
// if the string is a punctuation, it is not separated by spaces
// if the string is a quote, it is not separated by spaces
func join(slice []string) string {
	var builder strings.Builder

	qoute := false

	for idx, r := range slice {

		if r == "A" || r == "a" {
			builder.WriteString(" ")
			if strings.ContainsAny(string(slice[idx+1][0]), "aeiouh") {
				builder.WriteString(r + "n")
			} else {
				builder.WriteString(r)
			}
			continue
		}

		if strings.ContainsAny(r, "'") {
			qoute = true
			builder.WriteString(r)
			continue
		}

		if isPunc(r) {
			builder.WriteString(r)
			continue
		}

		if isPunc(string(r[0])) {
			builder.WriteByte(r[0])
			builder.WriteString(" ")
			builder.WriteString(r[1:])
			continue
		}

		if idx > 0 {
			if qoute {
				qoute = false
			} else {
				builder.WriteString(" ")
			}
		}
		builder.WriteString(r)
	}

	return builder.String()
}
