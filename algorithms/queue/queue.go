package main

import "fmt"

type MyItem int
type Queue struct {
	list []MyItem
}

func NewQueue() *Queue {
	return &Queue{
		list: []MyItem{},
	}
}

func (q *Queue) Enqueue(val MyItem) {
	q.list = append(q.list, val)
}

func (q *Queue) Dequeue() (MyItem, error) {
	if len(q.list) == 0 {
		return -1, fmt.Errorf("the queue is empty")
	}

	item := q.list[0]

	q.list = q.list[1:]
	return item, nil
}

func main() {
	q := NewQueue()
	q.Enqueue(MyItem(1))
	q.Enqueue(MyItem(2))
	q.Enqueue(MyItem(3))
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

}
