package main

import "fmt"

type Stack struct {
	Size int
	Top  int
	S    []int
}

func CreateStack(size int) Stack {
	return Stack{
		Size: size,
		Top:  -1,
		S:    make([]int, size),
	}
}

func (s Stack) Display() {
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("%d ", s.S[i])
	}
	fmt.Printf("\n")
}

func (s *Stack) Push(x int) {
	// Check if stack is full
	if s.Top == s.Size-1 {
		return
	}

	s.Top++
	s.S[s.Top] = x
}

func (s *Stack) Pop() int {
	// Check if stack is empty
	if s.Top == -1 {
		return -1
	}

	x := s.S[s.Top]
	s.Top--

	return x
}

func (s Stack) Peek(index int) int {
	// Validate index
	if s.Top-index < 0 {
		return -1
	}

	return s.S[s.Top-index]
}

func (s Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s Stack) IsFull() bool {
	return s.Top == s.Size-1
}

func (s Stack) StackTop() int {
	if s.IsEmpty() {
		return -1
	}

	return s.S[s.Top]
}

type StackLL struct {
	Top *Node
}

func (s StackLL) Display() {
	p := s.Top
	for p != nil {
		fmt.Printf("%d ", p.Data)
		p = p.Next
	}
	fmt.Printf("\n")
}

func (s *StackLL) Push(x int) {
	// Create node, set next to current stack top
	temp := &Node{x, s.Top}

	// Update stack top to point to new node
	s.Top = temp
}

func (s *StackLL) Pop() int {
	if s.Top == nil {
		return -1
	}

	// Store current top value
	x := s.Top.Data
	// Move top pointer to current top's next
	s.Top = s.Top.Next
	// Top is now deleted
	return x
}

func (s *StackLL) Peek(pos int) int {
	p := s.Top

	for i := 0; p != nil && i < pos; i++ {
		p = p.Next
	}

	if p == nil {
		return -1
	}

	return p.Data
}

func (s *StackLL) IsEmpty() bool {
	return s.Top == nil
}

func (s *StackLL) StackTop() int {
	if s.Top == nil {
		return -1
	}

	return s.Top.Data
}

func isValid(s string) bool {
	if s == "" {
		return true
	}

	// Used to verify the corresponding stack element is correct
	symbols := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	// Stack of runes
	st := make([]rune, 0, len(s))

	for _, ch := range s {
		switch ch {
		case '{', '(', '[':
			st = append(st, ch)
		case '}', ')', ']':
			// Handle invalid closing symbol (empty stack)
			if len(st) == 0 {
				return false
			}

			// Handle invalid closing symbol (out of order)
			if st[len(st)-1] != symbols[ch] {
				return false
			}

			// Pop top of stack for match
			st = st[:len(st)-1]
		}
	}

	// If stack is empty -> true
	return len(st) == 0
}

func convert(infix string) string {
	prec := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	getPrec := func(ch byte) int {
		if val, ok := prec[ch]; ok {
			return val
		}

		return 0
	}

	isOperand := func(ch byte) bool {
		if _, ok := prec[ch]; ok {
			return false
		}

		return true
	}

	st := make([]byte, 0, len(infix))

	stackTop := func(stack []byte) byte {
		if len(stack) == 0 {
			return 0
		}

		return stack[len(stack)-1]
	}

	postfix := make([]byte, 0, len(infix))

	i := 0
	for i < len(infix) {
		ch := infix[i]

		// We have an operand, add to postfix, advance i
		if isOperand(ch) {
			postfix = append(postfix, ch)
			i++
			continue
		}

		if getPrec(ch) > getPrec(stackTop(st)) {
			// We have an operator with higher prec, add to stack, advance i
			st = append(st, ch)
			i++
		} else {
			// We have an operator with lower prec, pop stack, retry i
			postfix = append(postfix, stackTop(st))
			st = st[:len(st)-1]
		}
	}

	// Pop any remaining values from stack and append to postfix
	for len(st) != 0 {
		postfix = append(postfix, stackTop(st))
		st = st[:len(st)-1]
	}

	return string(postfix)
}

func eval(postfix string) int {
	// Create stack of ints
	st := StackLL{nil}
	//st := CreateStack(len(postfix))

	isOperand := func(ch byte) bool {
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			return false
		}

		return true
	}

	for i := 0; i < len(postfix); i++ {
		ch := postfix[i]

		if isOperand(ch) {
			// Convert operand byte to int value and push onto stack
			st.Push(int(ch - '0'))
		} else { // Operator
			// Get top 2 operands
			x2 := st.Pop() // Note the order of terms, 2nd is top so we pop first
			x1 := st.Pop()

			res := 0
			switch ch {
			case '+':
				res = x1 + x2
			case '-':
				res = x1 - x2
			case '*':
				res = x1 * x2
			case '/':
				res = x1 / x2
			}

			// Push the resulting operand back onto stack
			st.Push(res)
		}
	}

	return st.Pop()
}
