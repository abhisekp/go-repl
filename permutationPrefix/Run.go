package permutationPrefix

import (
	"fmt"
	"slices"
)

func Run(A, B []int) []int {
	C := make([]int, len(A))

	seen := map[int]bool{}

	for i := range A {
		// fmt.Println("Seen", seen)
		// fmt.Println("A = ", A[i], " B = ", B[i])

		if i > 0 {
			C[i] = C[i-1]
		}

		if seen[A[i]] {
			C[i]++
		}

		if seen[B[i]] {
			C[i]++
		}

		if A[i] == B[i] {
			C[i]++
		}

		seen[A[i]] = true
		seen[B[i]] = true

		// fmt.Println("C = ", C[i])
	}

	return C
}

func Setup() {
	cases := []struct {
		A, B, C []int
	}{
		{A: []int{2, 3, 1}, B: []int{3, 1, 2}, C: []int{0, 1, 3}},
		{A: []int{1, 3, 2, 4}, B: []int{3, 1, 2, 4}, C: []int{0, 2, 3, 4}},
		{A: []int{2, 3, 1, 4}, B: []int{3, 1, 2, 4}, C: []int{0, 1, 3, 4}},
		{A: []int{1, 2, 3, 4}, B: []int{1, 2, 3, 4}, C: []int{1, 2, 3, 4}},
	}

	for _, _case := range cases {
		fmt.Println("----")
		actual := Run(_case.A, _case.B)

		same := slices.Equal(_case.C, actual)
		msg := "Not Equal"
		if same {
			msg = "Equal"
		}
		fmt.Printf("\nA = %v\nB = %v\nC = %v\nGot: %v (%s)\n", _case.A, _case.B, _case.C, actual, msg)
	}
}
