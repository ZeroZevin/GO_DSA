package circular_queue

import "fmt"

func CircularQueueTest() {
	queue := NewQueue()

	fmt.Println(queue.Empty())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.EnQueue(0))
	fmt.Println(queue.EnQueue(6))
	fmt.Println(queue.EnQueue(2))
	fmt.Println(queue.EnQueue(7))
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.GetQueue())
	fmt.Println(queue.EnQueue(8))
	fmt.Println(queue.EnQueue(3))
	fmt.Println(queue.EnQueue(8))
	fmt.Println(queue.GetQueue())
	fmt.Println(queue.EnQueue(1))
	fmt.Println(queue.EnQueue(3))
	fmt.Println(queue.EnQueue(4))
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.Empty())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.Empty())

	fmt.Println(queue)

}
