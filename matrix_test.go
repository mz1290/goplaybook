package main

import (
	"fmt"
	"testing"
)

func TestDiagonalMatrix(t *testing.T) {
	m := NewDiagonalMatrix(4)

	m.Set(0, 0, 5)
	m.Set(1, 1, 8)
	m.Set(2, 2, 9)
	m.Set(3, 3, 12)

	m.Display()

	fmt.Printf("\nGet(2,2)=%d\n", m.Get(2, 2)) // 9
}

func TestLowerTriangleMatrix(t *testing.T) {
	m := NewLowerTriangleMatrix(5)

	val := 1
	for row := 0; row < m.N; row++ {
		for col := 0; col < m.N; col++ {
			m.Set(row, col, val)
			val++
		}
	}

	m.Display()
}
