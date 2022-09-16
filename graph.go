package main

import "fmt"

func Rdfs(matrix [][]int, start int) {
	visited := make([]int, len(matrix[0]))
	var traverse func(int)

	traverse = func(start int) {
		if visited[start] == 1 {
			return
		}

		// Print that we are at vertex
		fmt.Printf("%d ", start)

		// Mark vertex as visited
		visited[start] = 1

		for v := 1; v < len(matrix[0]); v++ {
			if matrix[start][v] == 1 && visited[v] == 0 {
				traverse(v)
			}
		}
	}

	traverse(start)
}

func Idfs(matrix [][]int, start int) {
	visited := make([]int, len(matrix[0]))
	stack := []int{start}

	for len(stack) != 0 {
		stackTop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Acknowledge vertex if it has not been visited and mark as visited
		if visited[stackTop] == 0 {
			// Print that we have visited vertex
			fmt.Printf("%d ", stackTop)

			// Mark node as visited
			visited[stackTop] = 1
		}

		for v := 1; v < len(matrix[0]); v++ {
			if matrix[stackTop][v] == 1 && visited[v] == 0 {
				stack = append(stack, v)
			}
		}
	}
}
