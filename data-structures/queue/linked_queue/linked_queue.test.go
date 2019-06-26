package linked_queue

import "fmt"

func LinkedQueueTest() {
	queue := NewQueue()

	fmt.Println(queue.Empty())
	fmt.Println(queue.DeQueue())
	queue.EnQueue(6)
	queue.EnQueue(2)
	queue.EnQueue(7)
	queue.EnQueue(8)
	queue.EnQueue(3)
	fmt.Println(queue.GetQueue())
	queue.EnQueue(8)
	queue.EnQueue(1)
	fmt.Println(queue.GetQueue())
	queue.EnQueue(3)
	queue.EnQueue(4)
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.Empty())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.Empty())
}
