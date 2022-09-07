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

func DisplayCycleLL(ptr *Node) {
	head := ptr

	for {
		fmt.Printf("%d ", ptr.Data)
		ptr = ptr.Next

		if ptr == head {
			break
		}
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

func InsertCycleLL(p *Node, index, x int) *Node {
	// Validate index
	if index < 0 || index > FindCycleLength(p) {
		return nil
	}

	head := p
	temp := &Node{x, nil}

	// Check inserting at head position
	if index == 0 {
		// Check if this is the first node
		if head == nil {
			head = temp

			// Create cycle
			head.Next = head
		} else {
			for p.Next != head {
				p = p.Next
			}

			p.Next = temp
			temp.Next = head
			head = temp // optional
		}
	} else { // Thist matches inserting in linear linked list
		// Stop at the node just before our desired index
		for i := 0; i < index-1; i++ {
			p = p.Next
		}

		// Copy node before position's next to new node about to be inserted
		temp.Next = p.Next

		// Update that same node's next to now be this inserted node
		p.Next = temp
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

func DeleteNodeLL(ptr *Node, index int) *Node {
	head := ptr
	prev := head

	if index < 0 || index > LenLL(head) {
		return head
	}

	if index == 0 {
		return head.Next
	}

	for i := 0; i < index; i++ {
		prev = ptr
		ptr = ptr.Next
	}

	prev.Next = ptr.Next
	// Node at ptr has now been deleted from list GC will clean up

	return head
}

func DeleteNodeCycleLL(p *Node, idx int) *Node {
	head := p
	prev := head

	// Validate index
	if idx < 0 || idx > FindCycleLength(p)-1 {
		return nil
	}

	if idx == 0 {
		for p.Next != head {
			p = p.Next
		}

		if p == head {
			// We only have one node left, delete
			return nil
		} else {
			// Update tail to point to new head
			p.Next = head.Next

			// Delete head by simply returning it's next value
			return head.Next
		}
	}

	// Otherwise treat as you normally would linear LL delete
	for i := 0; i < idx; i++ {
		prev = p
		p = p.Next
	}

	prev.Next = p.Next

	return head
}

func IsLinkedListSorted(ptr *Node) bool {
	prev := ptr.Data
	ptr = ptr.Next

	for ptr != nil {
		if ptr.Data < prev {
			return false
		}

		prev = ptr.Data
		ptr = ptr.Next
	}

	return true
}

// Linked list must be sorted
func RemoveDuplicatesLL(ptr *Node) *Node {
	head := ptr
	next := ptr.Next

	for next != nil {
		if ptr.Data != next.Data {
			ptr = next
			next = next.Next
		} else {
			// Remove duplicate
			ptr.Next = next.Next
			// node pointed to by 'next' is now removed

			next = ptr.Next
		}
	}

	return head
}

func reverseLL(head *Node) *Node {
	// Sliding pointers
	current := head
	var prev *Node = nil

	for current != nil {
		// Store current's next for later assignment
		next := current.Next

		// Peform the link reversal
		current.Next = prev

		// Current is the next rounds previous
		prev = current

		// Current is the previously saved next from above
		current = next
	}

	// Previous will hold the new head when complete
	return prev
}

func reverseLLR(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	r := reverseLLR(head.Next)

	// Reverse the links on the recursive return by updating the updated list
	// with the new node and bubbling up the nil until reaching the original
	// first node
	head.Next.Next = head
	head.Next = nil
	return r
}

// Assumes both lists are sorted
func mergeLL(first, second *Node) *Node {
	var dummy = new(Node)

	// p serves as our LL iterator. It will take the lower node from the two
	// lists and update the previously appeneded node's next to point to new
	// node.
	p := dummy

	for first != nil && second != nil {
		if first.Data < second.Data {
			// The first value is smaller so attach to dummy list ahead
			p.Next = first

			// Advance to the following node in first list
			first = first.Next
		} else {
			p.Next = second
			second = second.Next
		}

		// Now that we've added a new node, advance to it so that we may update
		// it's "next" pointer
		p = p.Next
	}

	// Check if there are any trailing nodes and append
	if first != nil {
		p.Next = first
	} else {
		p.Next = second
	}

	return dummy.Next
}

func hasCycleLL(head *Node) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		/* Advance fast 2 nodes */
		fast = fast.Next.Next

		/* Advance slow 1 node */
		slow = slow.Next

		/* Check if pointers have met */
		if slow == fast {
			return true
		}
	}

	return false
}

// This also determines if list has cycle
func FindCycleLength(head *Node) int {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		/* Advance fast 2 nodes */
		fast = fast.Next.Next

		/* Advance slow 1 node */
		slow = slow.Next

		/* Check if pointers have met */
		if slow == fast {
			return CalculateCycleLength(slow)
		}
	}

	return 0
}

func CalculateCycleLength(slow *Node) int {
	current := slow
	cycleLen := 0

	for {
		/* Get next node */
		current = current.Next

		/* Increment length */
		cycleLen++

		/* Check if we've circled back */
		if current == slow {
			break
		}
	}

	return cycleLen
}
