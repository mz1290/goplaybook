package main

import "fmt"

type Matrix interface {
	Set(int, int, int)
	Get(int, int) int
	Display()
}

type DiagonalMatrix struct {
	Array []int
	N     int
}

func NewDiagonalMatrix(n int) *DiagonalMatrix {
	return &DiagonalMatrix{
		N:     n,
		Array: make([]int, n),
	}
}

func (m *DiagonalMatrix) Set(row, col, element int) {
	if row == col {
		m.Array[row] = element
	}
}

func (m DiagonalMatrix) Get(row, col int) int {
	if row == col {
		return m.Array[row]
	}

	return 0
}

func (m DiagonalMatrix) Display() {
	for row := 0; row < m.N; row++ {
		for col := 0; col < m.N; col++ {
			if row == col {
				fmt.Printf("%d ", m.Array[row])
			} else {
				fmt.Printf("0 ")
			}
		}

		fmt.Printf("\n")
	}
}

type LowerTriangleMatrix struct {
	Array []int
	N     int
}

func NewLowerTriangleMatrix(n int) *LowerTriangleMatrix {
	return &LowerTriangleMatrix{
		N:     n,
		Array: make([]int, n*(n+1)/2),
	}
}

func (m *LowerTriangleMatrix) Set(row, col, element int) {
	if row >= col {
		// Row major formula for finding flat index
		idx := row*(row)/2 + col
		m.Array[idx] = element
	}
}

func (m LowerTriangleMatrix) Get(row, col int) int {
	if row >= col {
		// Row major
		idx := row*(row)/2 + col

		return m.Array[idx]
	}

	return 0
}

func (m LowerTriangleMatrix) Display() {
	for row := 0; row < m.N; row++ {
		for col := 0; col < m.N; col++ {
			idx := row*(row)/2 + col // Row Major

			if row >= col {
				fmt.Printf("%d ", m.Array[idx])
			} else {
				fmt.Printf("0 ")
			}
		}

		fmt.Printf("\n")
	}
}
