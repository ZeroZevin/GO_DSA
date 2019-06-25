package sequential_stack

import "fmt"

func SequentialStackTest() {
	stack := NewStack()

	fmt.Println(stack.Empty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.GetTop())
	fmt.Println(stack.Push(6)) // 6 | 0
	fmt.Println(stack.Push(2)) // 6 2 | 1
	fmt.Println(stack.Push(7)) // 6 2 7 | 2
	fmt.Println(stack.Push(8)) // 6 2 7 8 | 3
	fmt.Println(stack.Push(3)) // 6 2 7 8 3 | 4
	fmt.Println(stack.Push(8)) // 6 2 7 8 3 8 | 5
	fmt.Println(stack.Push(1)) // 6 2 7 8 3 8 1 | 6
	fmt.Println(stack.Push(3)) // 6 2 7 8 3 8 1 3 | 7
	fmt.Println(stack.Push(4)) // 6 2 7 8 3 8 1 3 4 | 8
	fmt.Println(stack.Push(0)) // 6 2 7 8 3 8 1 3 4 0 | 9
	fmt.Println(stack.Push(0)) // 6 2 7 8 3 8 1 3 4 0 | 9
	fmt.Println(stack.Pop())   // 6 2 7 8 3 8 1 3 4 0 | 8
	fmt.Println(stack.Pop())   // 6 2 7 8 3 8 1 3 4 0 | 7
	fmt.Println(stack.Empty())
	fmt.Println(stack)
}
