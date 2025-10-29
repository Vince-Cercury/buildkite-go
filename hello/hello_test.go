package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "greet with single name",
			input:    "Alice",
			expected: "Hello, Alice!",
		},
		{
			name:     "greet with full name",
			input:    "Alice Smith",
			expected: "Hello, Alice Smith!",
		},
		{
			name:     "greet with empty string",
			input:    "",
			expected: "Hello, World!",
		},
		{
			name:     "greet with multiple spaces",
			input:    "Alice   Bob",
			expected: "Hello, Alice   Bob!",
		},
		{
			name:     "greet with special characters",
			input:    "Alice-Bob O'Connor",
			expected: "Hello, Alice-Bob O'Connor!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Greet(tt.input)
			if result != tt.expected {
				t.Errorf("Greet(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
