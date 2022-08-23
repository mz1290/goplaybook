package main

func factorial(n int) int {
	if n < 0 || n == 0 {
		return 0
	}

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	return factorial
}

// Time O(n * n!)
// Space O(n * n!)
func permutations1(str string) []string {
	permutations := make([]string, 0, factorial(len(str)))
	visited := make([]int, len(str))
	result := make([]byte, len(str))

	var stateSpaceTreePermutations func(string, int)
	stateSpaceTreePermutations = func(s string, k int) {
		// Base case
		if k >= len(s) {
			permutations = append(permutations, string(result))
			return
		}

		// Generate all permutations of string
		for i := 0; i < len(s); i++ {
			// If flag is unset, traverse
			if visited[i] == 0 {
				// Take the ith char from string and put into our result buffer
				result[k] = s[i]

				// Set the flag for this level as visited
				visited[i] = 1

				// Recursive case:
				// Advance to next level kth sport for result buffer
				stateSpaceTreePermutations(s, k+1)

				// We've completed all permutations for ith char, reset level
				visited[i] = 0
			}
		}
	}

	stateSpaceTreePermutations(str, 0)
	return permutations
}

func permutations2(str []byte) [][]byte {
	permutations := make([][]byte, factorial(len(str)))
	for i := range permutations {
		permutations[i] = make([]byte, len(str))
	}

	var swapPermutations func([]byte, int)
	current := 0
	swapPermutations = func(s []byte, nSwaps int) {
		// Base case
		if nSwaps == len(s)-1 {
			copy(permutations[current], s)
			current++
			return
		}

		// Generate all permutations for string
		for i := nSwaps; i < len(s); i++ {
			s[nSwaps], s[i] = s[i], s[nSwaps] // swap
			swapPermutations(s, nSwaps+1)
			s[nSwaps], s[i] = s[i], s[nSwaps] // undo swap
		}
	}

	swapPermutations(str, 0)
	return permutations
}
