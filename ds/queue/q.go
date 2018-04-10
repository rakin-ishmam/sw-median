package queue

import "errors"

type node struct {
	value interface{}
	next  *node
}

type queue struct {
	size  int
	front *node
	tail  *node
}

func (q queue) Size() int {
	return q.size
}

func (q *queue) Push(v interface{}) {
	n := node{
		value: v,
		next:  nil,
	}

	q.size++

	if q.tail == nil {
		q.front = &n
		q.tail = &n
		return
	}

	q.tail.next = &n
	q.tail = &n
}

func (q *queue) Front() (interface{}, error) {
	if q.front == nil {
		return nil, errors.New("empty")
	}

	return q.front.value, nil
}

func (q *queue) Pop() error {
	if q.front == nil {
		return errors.New("empty")
	}

	q.front = q.front.next
	if q.front == nil {
		q.tail = nil
	}

	q.size--

	return nil
}

// Queue wraps the queue functinality
type Queue interface {
	Push(interface{})
	Pop() error
	Front() (interface{}, error)
	Size() int
}

// New returns instance of queue
func New() Queue {
	return &queue{}
}
