package main

import (
	"fmt"
	"testing"
)

func funTail(n int) {
	if n <= 0 {
		return
	}

	fmt.Println(n)
	funTail(n - 1)
}

func funHead(n int) {
	if n <= 0 {
		return
	}

	funHead(n - 1)
	fmt.Println(n)
}

func funClosure(n int) int {
	x := 0
	var f func(int) int

	f = func(num int) int {
		if num <= 0 {
			return 0
		}

		x++
		return f(num-1) + x
	}

	return f(n)
}

func funTree(n int) {
	if n <= 0 {
		return
	}

	fmt.Println(n)
	funTree(n - 1)
	funTree(n - 1)
}

func funIndirectA(n int) {
	if n <= 0 {
		return
	}

	fmt.Println(n)
	funIndirectB(n - 1)
}

func funIndirectB(n int) {
	if n < 1 {
		return
	}

	fmt.Println(n)
	funIndirectA(n / 2)
}

func funNested(n int) int {
	if n > 100 {
		return n - 10
	} else {
		return funNested(funNested(n + 11))
	}
}

func TestRecursion(t *testing.T) {
	fmt.Println("Tail recursion:")
	funTail(3) // 3 2 1
	fmt.Println("Head recursion:")
	funHead(3) // 1 2 3
	fmt.Println("Closure recursion:")
	fmt.Println(funClosure(5)) // 25
	fmt.Println("Tree recursion:")
	funTree(3) // 3 2 1 1 2 1 1
	fmt.Println("Indirect recursion:")
	funIndirectA(20) // 20 19 9 8 4 3 1
	fmt.Println("Nested recursion:")
	fmt.Println(funNested(95)) // 91
}

func TestTaylorSeries(t *testing.T) {
	fmt.Println("Taylor Series Recusrive brute force")
	fmt.Println(TaylorSeries(1, 10)) // 2.7182818011463845
	fmt.Println(TaylorSeries(3, 10)) // 20.079665178571425
	fmt.Println(TaylorSeries(4, 10)) // 54.44310405643739
	fmt.Println(TaylorSeries(4, 15)) // 54.5978829056501

	fmt.Println("Taylor Series Iterative optimized")
	fmt.Println(TaylorSereisIterative(1, 10)) // 2.7182818011463845
	fmt.Println(TaylorSereisIterative(3, 10)) // 20.079665178571425
	fmt.Println(TaylorSereisIterative(4, 10)) // 54.44310405643739
	fmt.Println(TaylorSereisIterative(4, 15)) // 54.5978829056501

	fmt.Println("Taylor Series Recursive optimized")
	fmt.Println(TaylorSeries2(1, 10)) // 2.7182818011463845
	fmt.Println(TaylorSeries2(3, 10)) // 20.079665178571425
	fmt.Println(TaylorSeries2(4, 10)) // 54.44310405643739
	fmt.Println(TaylorSeries2(4, 15)) // 54.5978829056501

}

func TestFibonacci(t *testing.T) {
	fmt.Println("Fib memoization")
	fmt.Println(FibonacciMemo(5)) // 5
	fmt.Println(FibonacciMemo(7)) // 13

	fmt.Println("Fib Recursive")
	fmt.Println(FibonacciRecursive(5)) // 5
	fmt.Println(FibonacciRecursive(7)) // 13

	fmt.Println("Fib Iterative")
	fmt.Println(FibonacciIterative(5)) // 5
	fmt.Println(FibonacciIterative(7)) // 13
}

func TestFactorialR(t *testing.T) {
	fmt.Println(factorialR(-1)) // 0
	fmt.Println(factorialR(0))  // 1
	fmt.Println(factorialR(1))  // 1
	fmt.Println(factorialR(3))  // 6
}

func TestCombinationsR(t *testing.T) {
	fmt.Println(combinationsR(5, 1)) // 5
	fmt.Println(combinationsR(5, 2)) // 10
	fmt.Println(combinationsR(5, 3)) // 10
	fmt.Println(combinationsR(4, 2)) // 6
}

func TestToh(t *testing.T) {
	toh(3, 1, 2, 3)
}
