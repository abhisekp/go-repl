package canbepalindrome

import "testing"

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
	}

	for _, tc := range testCases {
		t.Run("can be palindrome?", func(t *testing.T) {
			actual := CanBePalindrome(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
