package linked_queue

import "errors"

type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	rear *Node
	head *Node
}

func NewQueue() *Queue {
	queue := new(Queue)

	return queue
}

func (queue *Queue) EnQueue(v interface{}) {
	node := new(Node)
	node.data = v

	if queue.Empty() {
		queue.head = node
		queue.rear = node
	} else {
		p := queue.head

		for p.next != nil {
			p = p.next
		}

		p.next = node
		queue.rear = node
	}

	// node.next = queue.rear

	// queue.rear = node
}

func (queue *Queue) DeQueue() (interface{}, error) {
	if queue.head == nil {
		return nil, errors.New("队列为空")
	}

	v := queue.head.data
	queue.head = queue.head.next

	return v, nil
}

func (queue *Queue) GetQueue() interface{} {
	if queue.head == nil {
		return nil
	}

	return queue.head.data
}

func (queue *Queue) Empty() bool {
	return queue.head == nil
}
