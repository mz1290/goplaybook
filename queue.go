package main

import "fmt"

type Queue struct {
	Size  int
	Front int
	Back  int
	Q     []int
}

func newQueue(size int) Queue {
	return Queue{
		Size:  size,
		Front: -1,
		Back:  -1,
		Q:     make([]int, size),
	}
}

func (q *Queue) enqueue(x int) {
	// Check if queue has capacity
	if q.Back == q.Size-1 {
		return
	}

	q.Back++
	q.Q[q.Back] = x
}

func (q *Queue) dequeue() int {
	// Check if queue is empty
	if q.Front == q.Back {
		return -1
	}

	// By moving the front pointer we "delete" the front queue value
	q.Front++
	return q.Q[q.Front]
}

func (q Queue) display() {
	for i := q.Front + 1; i <= q.Back; i++ {
		fmt.Printf("%d ", q.Q[i])
	}
	fmt.Printf("\n")
}

func newCircularQueue(size int) Queue {
	return Queue{
		Size:  size,
		Front: 0,
		Back:  0,
		Q:     make([]int, size),
	}
}

func (q *Queue) circularEnqueue(x int) {
	// Check if circular queue has capacity
	if (q.Back+1)%q.Size == q.Front {
		return
	}

	q.Back = (q.Back + 1) % q.Size
	q.Q[q.Back] = x
}

func (q *Queue) circularDequeue() int {
	// Check if queue is empty
	if q.Front == q.Back {
		return -1
	}

	q.Front = (q.Front + 1) % q.Size
	return q.Q[q.Front]
}

func (q Queue) circularDisplay() {
	i := q.Front + 1
	for {
		fmt.Printf("%d ", q.Q[i])
		i = (i + 1) % q.Size

		if i == (q.Back+1)%q.Size {
			break
		}
	}

	fmt.Printf("\n")
}

type QueueLL struct {
	Front *Node
	Back  *Node
}

func newQueueLL() QueueLL {
	return QueueLL{nil, nil}
}

func (q QueueLL) Len() int {
	count := 0

	p := q.Front
	for p != nil {
		count++
		p = p.Next
	}

	return count
}

func (q *QueueLL) isEmpty() bool {
	return q.Front == nil
}

func (q *QueueLL) enqueue(x int) {
	// Create new node
	temp := &Node{x, nil}

	// Check if this is the first node
	if q.Front == nil {
		q.Front = temp
		q.Back = temp
	} else {
		// Append the new node to the queue
		q.Back.Next = temp

		// We now have a new end of the queue, update accordingly
		q.Back = temp
	}
}

func (q *QueueLL) dequeue() int {
	// Check if queue is empty
	if q.Front == nil {
		return -1
	}

	// Store the value being dequeued for return
	x := q.Front.Data

	// Update new front of queue by deleting old
	q.Front = q.Front.Next

	return x
}

func (q *QueueLL) display() {
	p := q.Front
	for p != nil {
		fmt.Printf("%d ", p.Data)
		p = p.Next
	}
	fmt.Printf("\n")
}

// Basic idea is to push all value into one stack (enqueue)
// Then when ready to dequeue, we transfer enqueue over to other
// stack and pop accordingly.
type QueueStacks struct {
	enq []int
	deq []int
}

func (q *QueueStacks) enqueue(x int) {
	q.enq = append(q.enq, x)
}

func (q *QueueStacks) dequeue() int {
	if len(q.deq) == 0 {
		if len(q.enq) == 0 {
			return -1
		}

		for len(q.enq) != 0 {
			// Pop item from top of enqueue stack
			val := q.enq[len(q.enq)-1]
			q.enq = q.enq[:len(q.enq)-1]

			// Push item into dequeue stack
			q.deq = append(q.deq, val)
		}
	}

	// Pop item from top of dequeue stack
	val := q.deq[len(q.deq)-1]
	q.deq = q.deq[:len(q.deq)-1]

	return val
}
