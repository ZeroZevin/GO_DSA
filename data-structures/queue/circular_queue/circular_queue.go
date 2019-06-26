package circular_queue

import "errors"

const MaxSize = 10

type Queue struct {
	data [MaxSize]interface{}
	head int
	rear int
}

func NewQueue() *Queue {
	queue := new(Queue)

	return queue
}

func (queue *Queue) EnQueue(v interface{}) error {
	if (queue.rear+1)%MaxSize == queue.head {
		return errors.New("队列已满")
	}

	queue.rear = (queue.rear + 1) % MaxSize

	queue.data[queue.rear] = v

	return nil
}

func (queue *Queue) DeQueue() (interface{}, error) {
	if queue.rear == queue.head {
		return nil, errors.New("队列为空")
	}

	queue.head = (queue.head + 1) % MaxSize

	return queue.data[queue.head], nil
}

func (queue *Queue) GetQueue() interface{} {
	if queue.rear == queue.head {
		return nil
	}

	i := (queue.head + 1) % MaxSize

	return queue.data[i]
}

func (queue *Queue) Empty() bool {
	return queue.head == queue.rear
}
