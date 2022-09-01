package main

import (
	"fmt"
	"testing"
)

func TestCreateLL(t *testing.T) {
	test := []int{10, 20, 30, 40, 50}
	head := CreateLL(test)
	DisplayLL(head)
	fmt.Printf("len=%d\n", LenLL(head))
}

func TestInsertLL(t *testing.T) {
	test := []int{10, 20, 30, 40, 50}
	head := CreateLL(test)

	fmt.Println("BEFORE")
	DisplayLL(head)
	fmt.Printf("len=%d\n", LenLL(head))

	// Insert node in middle
	fmt.Println("MIDDLE")
	head = InsertLL(head, 3, 35)
	DisplayLL(head)
	fmt.Printf("len=%d\n", LenLL(head))

	// Insert new head
	fmt.Println("NEW HEAD")
	head = InsertLL(head, 0, 0)
	DisplayLL(head)
	fmt.Printf("len=%d\n", LenLL(head))
}

func TestInsertLLCreate(t *testing.T) {
	head := InsertLL(nil, 0, 1)
	head = InsertLL(head, 1, 2)
	head = InsertLL(head, 2, 3)
	head = InsertLL(head, 3, 4)
	DisplayLL(head)
}

func TestInsertSortedLL(t *testing.T) {
	test := []int{10, 20, 30, 40, 50}
	head := CreateLL(test)
	head = InsertSortedLL(head, 5)
	DisplayLL(head)

	head = InsertSortedLL(nil, 5)
	DisplayLL(head)
}
