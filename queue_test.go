package main

import (
	"fmt"
	"testing"
)

func TestQueueArray(t *testing.T) {
	q := newQueue(5)
	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)
	q.display()

	fmt.Println(q.dequeue())
}

func TestCircularQueueArray(t *testing.T) {
	q := newCircularQueue(5)
	q.circularEnqueue(10)
	q.circularEnqueue(20)
	q.circularEnqueue(30)
	q.circularEnqueue(40)
	q.circularEnqueue(50) // should not get added (full)
	q.circularEnqueue(60) // should not get added (full)
	q.circularDisplay()
	fmt.Println(q.circularDequeue())
	q.circularEnqueue(50)
	q.circularDisplay()
}

func TestQueueLL(t *testing.T) {
	q := newQueueLL()
	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)
	q.enqueue(40)
	q.enqueue(50)
	q.display()

	fmt.Println(q.dequeue())
}

func TestQueueStacks(t *testing.T) {
	var q QueueStacks
	test := []int{1, 3, 5, 7, 9}

	for _, num := range test {
		q.enqueue(num)
	}
	fmt.Printf("Enqueue: %v\n", q.enq)

	fmt.Printf("Dequeue: ")
	for range test {
		fmt.Printf("%d ", q.dequeue())
	}
	fmt.Printf("\n")
}
