package main

// https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {
	// Init two pointer for bottoms-up DP
	oneStep := 1
	twoStep := 1

	// The idea is that we are doing memoization array via iteration. Stopping
	// when i == 1 means our two pointers are at index 0,1 and any further
	// processing would be underflow.
	for i := n; i > 1; i-- {
		// Store current value of oneStep before updating
		temp := oneStep

		// Update oneStep by adding the oneStep and twoStep paths. This also
		// serves as the bottoms-up result value because it is the sum og the
		// two previous options.
		oneStep = oneStep + twoStep

		// Advance the two step pointer to the pervious one step for the next
		// iteration
		twoStep = temp
	}

	return oneStep
}
