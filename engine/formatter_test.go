package engine

import "testing"

func TestModiefiers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "hex",
			input:    "1E",
			expected: "30",
		},
		{
			name:     "bin",
			input:    "11",
			expected: "3",
		},
		{
			name:     "up",
			input:    "daniel",
			expected: "DANIEL",
		},
		{
			name:     "low",
			input:    "JOSHUA",
			expected: "joshua",
		},
		{
			name:     "cap",
			input:    "bolA",
			expected: "Bola",
		},
	}

	checker := func(input, expected string, fn func(s string) string) {
		if fn(input) != expected {
			t.Errorf("Tried %v expected %v got %v", input, expected, fn(input))
		}
	}

	for _, r := range tests {
		switch r.name {
		case "hex":
			checker(r.input, r.expected, hex)
		case "bin":
			checker(r.input, r.expected, bin)
		case "up":
			checker(r.input, r.expected, up)
		case "low":
			checker(r.input, r.expected, low)
		case "cap":
			checker(r.input, r.expected, cap)
		}
	}
}

func TestIsPunc(t *testing.T) {
	input := []string{".", "...", "?!", ":", "name.."}
	expected := []bool{true, true, true, true, false}

	for i := 0; i < len(input); i++ {
		if isPunc(input[i]) != expected[i] {
			t.Errorf("Test '%v' expect %v got %v", input[i], expected[i], isPunc(input[i]))
		}
	}
}
