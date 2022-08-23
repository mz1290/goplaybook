package main

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	fmt.Println(factorial(-1)) // 0
	fmt.Println(factorial(0))  // 0
	fmt.Println(factorial(1))  // 1
	fmt.Println(factorial(3))  // 6
}

func TestPermutations1(t *testing.T) {
	fmt.Println(permutations1("ABC")) // ABC ACB BAC BCA CAB CBA
}

func TestPermutations2(t *testing.T) {
	fmt.Printf("%s\n", permutations2([]byte("ABC"))) // ABC ACB BAC BCA CBA CAB
}
