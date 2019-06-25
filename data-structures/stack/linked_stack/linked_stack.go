package linked_stack

import "errors"

type Node struct {
	data interface{}
	next *Node
}

type LinkedStack struct {
	top *Node
}

func NewStack() *LinkedStack {
	stack := new(LinkedStack)

	return stack
}

func (stack *LinkedStack) Push(v interface{}) {
	node := new(Node)

	node.data = v
	node.next = stack.top
	stack.top = node
}

func (stack *LinkedStack) Pop() (interface{}, error) {
	if stack.top == nil {
		return nil, errors.New("下溢")
	}

	v := stack.top.data
	stack.top = stack.top.next

	return v, nil
}

func (stack *LinkedStack) GetTop() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("栈为空")
	}

	return stack.top, nil
}

func (stack *LinkedStack) Empty() bool {
	return stack.top == nil
}
