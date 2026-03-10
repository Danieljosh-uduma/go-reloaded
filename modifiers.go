package formatter

import (
	"strconv"
	"strings"
	"unicode"
)

func hex(s string) string {
	decimal, err := strconv.ParseInt(s, 16, 64)

	if err != nil {
		return s
	}

	return strconv.Itoa(int(decimal))
}

func bin(s string) string {
	decimal, err := strconv.ParseInt(s, 2, 64)

	if err != nil {
		return s
	}

	return strconv.Itoa(int(decimal))
}

func up(s string) string {
	return strings.ToUpper(s)
}

func low(s string) string {
	return strings.ToLower(s)
}

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

func isPunc(s string) bool {
	for _, r := range s {
		if !unicode.IsPunct(r) {
			return false
		}
	}

	return true
}

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
