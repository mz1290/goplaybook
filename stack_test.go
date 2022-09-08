package main

import (
	"fmt"
	"testing"
)

func TestStackPush(t *testing.T) {
	st := CreateStack(5)

	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Push(40)
	st.Push(50)
	st.Push(60) // stack overflow should not show
	st.Display()
}

func TestStackPop(t *testing.T) {
	st := CreateStack(5)

	st.Push(10)
	st.Push(20)
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop()) // stack underflow (pop on empty stack)
	st.Display()
}

func TestStackPeek(t *testing.T) {
	st := CreateStack(5)

	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Push(40)
	st.Display()

	fmt.Println(st.Peek(0))
	fmt.Println(st.Peek(1))
	fmt.Println(st.Peek(2))
	fmt.Println(st.Peek(3))
	fmt.Println(st.Peek(4)) // invalid index
}

func TestStackLLPush(t *testing.T) {
	st := StackLL{nil}
	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Display()
}

func TestStackLLPop(t *testing.T) {
	st := StackLL{nil}
	st.Push(10)
	st.Push(20)
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop()) // stack underflow (pop on empty stack)
	st.Display()
}

func TestStackLLPeek(t *testing.T) {
	st := StackLL{nil}

	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Push(40)
	st.Display()

	fmt.Println(st.Peek(0))
	fmt.Println(st.Peek(1))
	fmt.Println(st.Peek(2))
	fmt.Println(st.Peek(3))
	fmt.Println(st.Peek(4)) // invalid index
}

func TestIsValid(t *testing.T) {
	test := "{([a+b] * [c - d]) / e}"
	fmt.Println(isValid(test)) // true

	test = "{([a+b] * [c - d]) / e}}"
	fmt.Println(isValid(test)) // false

	test = "{[a+b] * [c - d]) / e}"
	fmt.Println(isValid(test)) // false
}

func TestConvert(t *testing.T) {
	fmt.Println(convert("a+b*c-d/e")) // abc*+de/-
}

func TestEval(t *testing.T) {
	fmt.Println(eval("234*+82/-"))
}
