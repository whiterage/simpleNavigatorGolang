package queue

import "errors"

var ErrEmptyQueue = errors.New("queue is empty")

type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]int, 0),
	}
}

func (q *Queue) Push(v int) {
	q.items = append(q.items, v)
}

func (q *Queue) Pop() (int, error) {
	front, err := q.Front()
	if err != nil {
		return front, err
	}
	q.items = q.items[1:]
	return front, nil
}

func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, ErrEmptyQueue
	}

	return q.items[0], nil
}

func (q *Queue) Back() (int, error) {
	if q.IsEmpty() {
		return 0, ErrEmptyQueue
	}

	return q.items[q.Size()-1], nil
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
