package sequential_stack

import "errors"

const MaxSize = 10

type SequentialStack struct {
	data [MaxSize]interface{}
	top  int
}

func NewStack() *SequentialStack {
	stack := new(SequentialStack)

	stack.top = -1

	return stack
}

func (stack *SequentialStack) Push(v interface{}) error {
	if stack.top+1 >= MaxSize {
		return errors.New("上溢")
	}

	stack.top++
	stack.data[stack.top] = v

	return nil
}

func (stack *SequentialStack) Pop() (interface{}, error) {
	if stack.top == -1 {
		return nil, errors.New("下溢")
	}

	v := stack.data[stack.top]
	stack.top--

	return v, nil
}

func (stack *SequentialStack) GetTop() (interface{}, error) {
	if stack.top == -1 {
		return nil, errors.New("栈为空")
	}

	v := stack.data[stack.top]

	return v, nil
}

func (stack *SequentialStack) Empty() bool {
	return stack.top == -1
}
