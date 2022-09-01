package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func CreateLL(arr []int) *Node {
	var first *Node
	var last *Node
	var temp *Node

	// Init list by starting first node
	first = &Node{arr[0], nil}
	last = first

	for i := 1; i < len(arr); i++ {
		// Create new node and use temp to point at it
		temp = &Node{arr[i], nil}

		// Append the new node to the LL by setting current last next to new
		// node
		last.Next = temp

		// Update our true new last node with this node
		last = temp
	}

	return first
}

func DisplayLL(ptr *Node) {
	for ptr != nil {
		fmt.Printf("%d ", ptr.Data)
		ptr = ptr.Next
	}
	fmt.Printf("\n")
}

func LenLL(ptr *Node) int {
	count := 0

	for ptr != nil {
		count++
		ptr = ptr.Next
	}

	return count
}

// Returns the original ptr as head
func InsertLL(ptr *Node, index, x int) *Node {
	if index < 0 || index > LenLL(ptr) {
		return nil
	}

	head := ptr
	temp := &Node{x, nil}

	// Check if we are inserting new head of list
	if index == 0 {
		// Assumes that ptr arg is the head of the list
		temp.Next = head
		head = temp
	} else {
		// Stop at the node just before our desired index
		for i := 0; i < index-1; i++ {
			ptr = ptr.Next
		}

		// Copy node before position's next to new node about to be inserted
		temp.Next = ptr.Next

		// Update that same node's next to now be this inserted node
		ptr.Next = temp
	}

	return head
}

func InsertSortedLL(ptr *Node, x int) *Node {
	head := ptr
	temp := &Node{x, nil}
	var prev *Node

	if head == nil {
		// Handles scenario where list is empty
		head = temp
	} else {
		// Iterate list and find spot where prev < x and ptr >= x
		for ptr != nil && ptr.Data < x {
			prev = ptr
			ptr = ptr.Next
		}

		if ptr == head {
			// If x is smallest element in list, we must update the head
			temp.Next = head
			head = temp
		} else {
			// Otherwise, insert sorted node
			temp.Next = prev.Next
			prev.Next = temp
		}
	}

	return head
}
