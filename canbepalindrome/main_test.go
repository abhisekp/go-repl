package canbepalindrome

import (
	"fmt"
	"testing"
)

func TestCanBePalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{
			input:    "madam",
			expected: true,
		},
		{
			input:    "mmaad",
			expected: true,
		},
		{
			input:    "abccbb",
			expected: false,
		},
		{
			input:    "abccab",
			expected: true,
		},
		{
			input:    "!@#$%",
			expected: false,
		},
		{
			input:    "()[]{}",
			expected: false,
		},
		{
			input:    "<>?/\\",
			expected: false,
		},
		{
			input:    "abbaaa",
			expected: true,
		},
		{
			input:    "racecarrr",
			expected: true,
		},
		{
			input:    "levelaa",
			expected: true,
		},
		{
			input:    "aabbc",
			expected: true,
		},
		{
			input:    "aabbcc",
			expected: true,
		},
		{
			input:    "aabbccd",
			expected: true,
		},
		{
			input:    "aabbcdd",
			expected: true,
		},
		{
			input:    "aabbccddee",
			expected: true,
		},
		{
			input:    "aabbccddeeff",
			expected: true,
		},
		{
			input:    "racecar",
			expected: true,
		},
		{
			input:    "hello",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("can be palindrome? %s", tc.input), func(t *testing.T) {
			actual := CanBePalindrome(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
