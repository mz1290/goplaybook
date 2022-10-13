package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Left  *TreeNode
	Data  int
	Right *TreeNode
}

type TreeStackNode struct {
	Data *TreeNode
	Next *TreeStackNode
}

type TreeStackLL struct {
	Top *TreeStackNode
}

func (s *TreeStackLL) Push(x *TreeNode) {
	// Create node, set next to current stack top
	temp := &TreeStackNode{x, s.Top}

	// Update stack top to point to new node
	s.Top = temp
}

func (s *TreeStackLL) Pop() *TreeNode {
	if s.Top == nil {
		return nil
	}

	// Store current top value
	x := s.Top.Data
	// Move top pointer to current top's next
	s.Top = s.Top.Next
	// Top is now deleted
	return x
}

func (s *TreeStackLL) IsEmpty() bool {
	return s.Top == nil
}

type TreeQueueNode struct {
	Data *TreeNode
	Next *TreeQueueNode
}

type TreeQueueLL struct {
	Front *TreeQueueNode
	Back  *TreeQueueNode
}

func (q *TreeQueueLL) enqueue(x *TreeNode) {
	// Create new node
	temp := &TreeQueueNode{x, nil}

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

func (q *TreeQueueLL) dequeue() *TreeNode {
	// Check if queue is empty
	if q.Front == nil {
		return nil
	}

	// Store the value being dequeued for return
	x := q.Front.Data

	// Update new front of queue by deleting old
	q.Front = q.Front.Next

	return x
}

func (q *TreeQueueLL) isEmpty() bool {
	return q.Front == nil
}

func (q TreeQueueLL) Len() int {
	count := 0

	p := q.Front
	for p != nil {
		count++
		p = p.Next
	}

	return count
}

func createTree(elements []int) *TreeNode {
	var p *TreeNode

	q := &TreeQueueLL{nil, nil}

	root := &TreeNode{nil, elements[0], nil}

	q.enqueue(root)

	i := 1
	//for !q.isEmpty() {
	for i < len(elements) {

		p = q.dequeue()

		// Add left child
		if elements[i] != -1 {
			temp := &TreeNode{nil, elements[i], nil}
			p.Left = temp
			q.enqueue(temp)
			i++
		}

		// Add right child
		if elements[i] != -1 {
			temp := &TreeNode{nil, elements[i], nil}
			p.Right = temp
			q.enqueue(temp)
			i++
		}
	}

	return root
}

func iterativePreorder(p *TreeNode) {
	st := TreeStackLL{nil}

	for p != nil || !st.IsEmpty() {
		if p != nil {
			fmt.Printf("%d ", p.Data)
			st.Push(p)
			p = p.Left
		} else {
			p = st.Pop()
			p = p.Right
		}
	}
}

func preorder(p *TreeNode) {
	if p == nil {
		return
	}

	fmt.Printf("%d ", p.Data)
	preorder(p.Left)
	preorder(p.Right)
}

func iterativeInorder(p *TreeNode) {
	st := TreeStackLL{nil}

	for p != nil || !st.IsEmpty() {
		if p != nil {
			st.Push(p)
			p = p.Left
		} else {
			p = st.Pop()
			fmt.Printf("%d ", p.Data)
			p = p.Right
		}
	}
}

func inorder(p *TreeNode) {
	if p == nil {
		return
	}

	inorder(p.Left)
	fmt.Printf("%d ", p.Data)
	inorder(p.Right)
}

func postorder(p *TreeNode) {
	if p == nil {
		return
	}

	postorder(p.Left)
	postorder(p.Right)
	fmt.Printf("%d ", p.Data)
}

// BFS
func levelorder(root *TreeNode) {
	q := &TreeQueueLL{nil, nil}

	fmt.Printf("%d ", root.Data)
	q.enqueue(root)

	for !q.isEmpty() {
		root = q.dequeue()

		if root.Left != nil {
			fmt.Printf("%d ", root.Left.Data)
			q.enqueue(root.Left)
		}

		if root.Right != nil {
			fmt.Printf("%d ", root.Right.Data)
			q.enqueue(root.Right)
		}
	}
}

func countNodes(p *TreeNode) int {
	// recursive case
	for p != nil {
		// Postorder traversal
		x := countNodes(p.Left)
		y := countNodes(p.Right)

		// left count + right count + THIS node
		return x + y + 1
	}

	// base case
	return 0
}

// DFS
func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	x := treeHeight(root.Left)
	y := treeHeight(root.Right)

	if x > y {
		return x + 1
	} else {
		return y + 1
	}
}

func insertBST(root *TreeNode, key int) *TreeNode {
	t := root
	var r, p *TreeNode

	// Must create first node
	if root == nil {
		return &TreeNode{nil, key, nil}
	}

	// Find the noe we must set 'r' for insert
	for t != nil {
		r = t

		if key < t.Data {
			t = t.Left
		} else if key > t.Data {
			t = t.Right
		} else {
			return root
		}
	}

	// Create new node for insert
	p = &TreeNode{nil, key, nil}
	if key < r.Data {
		r.Left = p
	} else {
		r.Right = p
	}

	return root
}

func searchBST(root *TreeNode, key int) *TreeNode {
	t := root

	for t != nil {
		if key == t.Data {
			return t
		} else if key < t.Data {
			t = t.Left
		} else {
			t = t.Right
		}
	}

	return nil
}

func deleteBSTR(p *TreeNode, key int) *TreeNode {
	if p == nil {
		return nil
	}

	if p.Left == nil && p.Right == nil {
		p = nil
		return p
	}

	if key < p.Data {
		p.Left = deleteBSTR(p.Left, key)
	} else if key > p.Data {
		p.Right = deleteBSTR(p.Right, key)
	} else {
		if treeHeight(p.Left) > treeHeight(p.Right) {
			temp := inorderPrecessor(p.Left)
			p.Data = temp.Data
			p.Left = deleteBSTR(p.Left, temp.Data)
		} else {
			temp := inorderSuccessor(p.Right)
			p.Data = temp.Data
			p.Right = deleteBSTR(p.Right, temp.Data)
		}
	}

	return p
}

func inorderPrecessor(p *TreeNode) *TreeNode {
	for p != nil && p.Right != nil {
		p = p.Right
	}

	return p
}

func inorderSuccessor(p *TreeNode) *TreeNode {
	for p != nil && p.Left != nil {
		p = p.Left
	}

	return p
}

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(*TreeNode) int

	// Variable to track max diameter seen
	res := 0

	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}

		// Recursively find the heights of left and right subtrees
		left := dfs(node.Left)
		right := dfs(node.Right)

		// Calculate diameter from left and right heights and update max
		res = max(res, 2+left+right)

		// Return the height from this node
		return 1 + max(left, right)
	}

	dfs(root)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)

	level := 0
	for queue.Len() > 0 {
		n := queue.Len()

		for i := 0; i < n; i++ {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		level++
	}
	return level
}

type TrieNode struct {
	Children  map[byte]*TrieNode
	EndOfWord bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children:  make(map[byte]*TrieNode),
		EndOfWord: false,
	}
}

type Trie struct {
	Root *TrieNode
}

func Constructor() Trie {
	return Trie{
		Root: NewTrieNode(),
	}
}

func (this *Trie) Insert(word string) {
	cur := this.Root

	// Check every char in word
	for i := 0; i < len(word); i++ {
		if _, ok := cur.Children[word[i]]; !ok {
			// Character not seen yet, add it to our children
			cur.Children[word[i]] = NewTrieNode()
		}

		// Character seen, advance to the childs node
		cur = cur.Children[word[i]]
	}

	// Cur is now set to last char of word
	cur.EndOfWord = true
}

func (this *Trie) Search(word string) bool {
	cur := this.Root

	for i := 0; i < len(word); i++ {
		if _, ok := cur.Children[word[i]]; !ok {
			// Character does not exist, therefore word does not exist
			return false
		}

		// Character seen, advance to the childs node
		cur = cur.Children[word[i]]
	}

	// Cur is on last char of word
	// Check if node is marked as end of word or if it is just prefix
	return cur.EndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.Root

	for i := 0; i < len(prefix); i++ {
		if _, ok := cur.Children[prefix[i]]; !ok {
			// Character does not exist, therefore prefix does not exist
			return false
		}

		// Character seen, advance to the childs node
		cur = cur.Children[prefix[i]]
	}

	// Cur is on last char of prefix
	// We don't care if prefix or prefix so can just return true
	return true
}
