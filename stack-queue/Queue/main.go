package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) EnQueue(i int) {
	q.data = append(q.data, i)
}

func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		return 0
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Display() {
	fmt.Println("Queue elements:", q.data)
}

func main() {
	queue := Queue{}
	queue.EnQueue(30)
	queue.EnQueue(23)
	queue.EnQueue(43)
	queue.EnQueue(54)
	queue.EnQueue(34)
	queue.Display()
	queue.Dequeue()
	queue.Display()

}
