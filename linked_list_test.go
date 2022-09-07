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

func TestDeleteNodeLL(t *testing.T) {
	test := []int{10, 20, 30, 40, 50}
	head := CreateLL(test)

	head = DeleteNodeLL(head, 1)
	DisplayLL(head)

	head = DeleteNodeLL(head, 0)
	DisplayLL(head)
}

func TestIsLinkedListSorted(t *testing.T) {
	test := []int{10, 20, 30, 40, 50}
	head := CreateLL(test)
	fmt.Println(IsLinkedListSorted(head)) // true

	head = InsertLL(head, 2, 100)
	fmt.Println(IsLinkedListSorted(head)) // false
}

func TestRemoveDuplicatesLL(t *testing.T) {
	test := []int{10, 20, 20, 20, 30, 40, 40, 40, 50}
	head := CreateLL(test)

	head = RemoveDuplicatesLL(head)
	DisplayLL(head)
}

func TestReverseLL(t *testing.T) {
	test := []int{10, 20, 30, 40, 50}
	head := CreateLL(test)

	fmt.Println("reversing links...")
	head = reverseLL(head)
	DisplayLL(head)

	// Reset LL
	head = CreateLL(test)
	fmt.Println("reversing links recursively...")
	head = reverseLLR(head)
	DisplayLL(head)
}

func TestMergeLL(t *testing.T) {
	test := []int{10, 20, 30, 40, 50}
	test2 := []int{5, 15, 25, 35, 45}
	first := CreateLL(test)
	second := CreateLL(test2)

	head := mergeLL(first, second)
	DisplayLL(head)
}

func TestHasCycleLL(t *testing.T) {
	test := []int{10, 20, 30, 40, 50}
	head := CreateLL(test)

	t1 := head.Next.Next
	t2 := head.Next.Next.Next.Next
	t2.Next = t1

	fmt.Println(hasCycleLL(head)) // true
}

func TestInsertCycleLL(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	head := CreateLL(test)

	t1 := head.Next.Next.Next.Next
	t1.Next = head

	head = InsertCycleLL(head, 5, 10)
	if head == nil {
		t.Fatal("got nil head")
	}
	DisplayCycleLL(head)
}

func TestDeleteCycleLL(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	head := CreateLL(test)

	t1 := head.Next.Next.Next.Next
	t1.Next = head

	head = DeleteNodeCycleLL(head, 4)
	if head == nil {
		t.Fatal("received nil head")
	}
	DisplayCycleLL(head)
}
