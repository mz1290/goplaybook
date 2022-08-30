package main

import "fmt"

func factorialR(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	return factorialR(n-1) * n
}

// Where n is the number of terms used to calculate a more precise value
func TaylorSeries(x, n float64) float64 {
	var power float64 = 1
	var factorial float64 = 1
	var e func(float64, float64) float64

	e = func(x1, n1 float64) float64 {
		if n1 == 0 {
			return 1
		}

		r := e(x1, n1-1)
		power *= x1
		factorial *= n1
		return r + power/factorial
	}

	return e(x, n)
}

func TaylorSereisIterative(x, n float64) float64 {
	var s float64 = 1

	for ; n > 0; n-- {
		s = 1 + x/n*s
	}

	return s
}

func TaylorSeries2(x, n float64) float64 {
	var s float64 = 1
	var e func(float64, float64) float64

	e = func(x1, n1 float64) float64 {
		if n1 == 0 {
			return s
		}

		s = 1 + x1/n1*s
		return e(x1, n1-1)
	}

	return e(x, n)
}

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return FibonacciRecursive(n-2) + FibonacciRecursive(n-1)
}

func FibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	t0 := 0
	t1 := 1
	sum := 0

	// Start at 2 since fib terms intialized above (t0 t1)
	for i := 2; i <= n; i++ {
		sum = t0 + t1
		t0 = t1
		t1 = sum
	}

	return sum
}

func FibonacciMemo(n int) int {
	// We don't need to have +1 because we're never going to store the final
	// result of 'n' in the memoization table. It will just be returned.
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}

	var f func(int) int
	f = func(n1 int) int {
		if n1 <= 1 {
			memo[n1] = n1
			return n1
		}

		if memo[n1-2] == -1 {
			memo[n1-2] = f(n1 - 2)
		}

		if memo[n1-1] == -1 {
			memo[n1-1] = f(n1 - 1)
		}

		return memo[n1-2] + memo[n1-1]
	}

	return f(n)
}

func combinationsR(n, r int) int {
	if r == 0 || n == r {
		return 1
	}

	return combinationsR(n-1, r-1) + combinationsR(n-1, r)
}

func toh(n, a, b, c int) {
	if n <= 0 {
		return
	}

	toh(n-1, a, c, b)
	fmt.Printf("from %d to %d\n", a, c)
	toh(n-1, b, a, c)
}
