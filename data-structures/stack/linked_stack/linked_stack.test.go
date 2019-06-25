package linked_stack

import "fmt"

func LinkedStackTest() {
	stack := NewStack()

	fmt.Println(stack.Empty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.GetTop())
	stack.Push(6)            // 6 | 0
	stack.Push(2)            // 6 2 | 1
	stack.Push(7)            // 6 2 7 | 2
	stack.Push(8)            // 6 2 7 8 | 3
	stack.Push(3)            // 6 2 7 8 3 | 4
	stack.Push(8)            // 6 2 7 8 3 8 | 5
	stack.Push(1)            // 6 2 7 8 3 8 1 | 6
	stack.Push(3)            // 6 2 7 8 3 8 1 3 | 7
	stack.Push(4)            // 6 2 7 8 3 8 1 3 4 | 8
	stack.Push(0)            // 6 2 7 8 3 8 1 3 4 0 | 9
	stack.Push(0)            // 6 2 7 8 3 8 1 3 4 0 | 9
	fmt.Println(stack.Pop()) // 6 2 7 8 3 8 1 3 4 0 | 8
	fmt.Println(stack.Pop()) // 6 2 7 8 3 8 1 3 4 0 | 7
	fmt.Println(stack.Empty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}
