package formatter

import (
	"strconv"
	"strings"
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
	for _, r := range runes {
		builder.WriteRune(r)
	}
	return builder.String()
}
