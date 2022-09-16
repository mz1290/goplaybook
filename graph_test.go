package main

import (
	"fmt"
	"testing"
)

func TestDfs(t *testing.T) {
	test := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
		{0, 0, 1, 1, 0, 1, 1},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
	}

	// Recursive
	Rdfs(test, 1)
	fmt.Println()

	// Iterative
	Idfs(test, 1)
	fmt.Println()
}
