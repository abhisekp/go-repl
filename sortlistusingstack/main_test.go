package sortlistusingstack

import (
	"slices"
	"testing"
)

func TestSortListUsingStack(t *testing.T) {

	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "test case 1",
			input:    []int{5, 2, 8, 1, 9},
			expected: []int{1, 2, 5, 8, 9},
		},
		{
			name:     "test case 2",
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Already sorted",
			input:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "test case 4",
			input:    []int{1, 9, 3, 6, 5, 2, 0},
			expected: []int{0, 1, 2, 3, 5, 6, 9},
		},
		{
			name:     "test case 4",
			input:    []int{3, 1, 2, 3, 2, 1},
			expected: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name:     "test case 4",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "test case 4",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "test case 4",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "test case 4",
			input:    []int{1, 1, 1, 1, 1},
			expected: []int{1, 1, 1, 1, 1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := SortListUsingStack(testCase.input)
			if !slices.Equal(actual, testCase.expected) {
				t.Errorf("got %v, want %v", actual, testCase.expected)
			}
		})
	}
}
