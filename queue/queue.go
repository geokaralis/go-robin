package queue

import (
	"sync"
)

type Queue struct {
	Items []string
	lock sync.Mutex
}

func (q *Queue) NewQueue() *Queue {
	q.Items = []string{}
	return q
}

func (q *Queue) Enqueue(s string) {
	q.lock.Lock()
	q.Items = append(q.Items, s)
	q.lock.Unlock()
}

func (q *Queue) Dequeue() *string {
	q.lock.Lock()
	item := q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	q.lock.Unlock()
	return &item
}

func (q *Queue) Front() *string {
	q.lock.Lock()
	item := q.Items[0]
	q.lock.Unlock()
	return &item
}

func (q *Queue) Empty() bool {
	return len(q.Items) == 0
}

func (q *Queue) Size() int {
	return len(q.Items)
}