package main

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	test := []int{3, 7, 9, 10, 6, 5, 12, 4, 11, 2}

	fmt.Println("Before:", test)
	BubbleSort(test)
	fmt.Println("After:", test)
}

func TestInsertionSort(t *testing.T) {
	test := []int{11, 13, 7, 2, 6, 9, 4, 5, 10, 3}

	fmt.Println("Before:", test)
	InsertionSort(test)
	fmt.Println("After:", test)
}

func TestSelectionSort(t *testing.T) {
	test := []int{11, 13, 7, 2, 6, 9, 4, 5, 10, 3}

	fmt.Println("Before:", test)
	SelectionSort(test)
	fmt.Println("After:", test)
}

func TestQuickSort(t *testing.T) {
	test := []int{11, 13, 7, 12, 16, 9, 24, 5, 10, 3}

	fmt.Println("Before:", test)
	QuickSort(test, 0, len(test)-1)
	fmt.Println("After:", test)
}

func TestMergeSort(t *testing.T) {
	test := []int{11, 13, 7, 12, 16, 9, 24, 5, 10, 3}

	fmt.Println("Before:", test)
	MergeSort(test)
	fmt.Println("After:", test)
}

func TestMergeSortR(t *testing.T) {
	test := []int{11, 13, 7, 12, 16, 9, 24, 5, 10, 3}

	fmt.Println("Before:", test)
	MergeSortR(test, 0, len(test)-1)
	fmt.Println("After:", test)
}

func TestCountSort(t *testing.T) {
	test := []int{11, 13, 7, 12, 16, 9, 24, 5, 10, 3}

	fmt.Println("Before:", test)
	CountSort(test)
	fmt.Println("After:", test)
}
