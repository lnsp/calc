package lib

import "errors"

var EmptyStackError = errors.New("Stack is empty.")

type IntStack []int

func (stack *IntStack) Push(item int) {
	*stack = append(*stack, item)
}

func (stack *IntStack) Pop() int {
	size := len(*stack)
	if size < 1 {
		return 0
	}

	value := (*stack)[size-1]
	*stack = (*stack)[:size-1]
	return value
}

func (stack *IntStack) Peek() int {
	return (*stack)[len(*stack)-1]
}

func (stack *IntStack) IsEmpty() bool {
	return len(*stack) == 0
}

func NewIntStack() IntStack {
	return make(IntStack, 0)
}

type ValueStack []string

func (stack *ValueStack) Push(item string) {
	*stack = append(*stack, item)
}

func (stack *ValueStack) Pop() string {
	size := len(*stack)
	if size < 1 {
		return ""
	}

	value := (*stack)[size-1]
	*stack = (*stack)[:size-1]
	return value
}

func (stack *ValueStack) Peek() string {
	return (*stack)[len(*stack)-1]
}

func (stack *ValueStack) IsEmpty() bool {
	return len(*stack) == 0
}

func NewStack() ValueStack {
	return make(ValueStack, 0)
}
