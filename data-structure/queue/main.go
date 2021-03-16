package main

import "fmt"

// Queue sdfsdfsdf
type Queue struct {
	items []int
}

func (q *Queue) enqueue(i int) {
	q.items = append(q.items, i)
}

func (q Queue) printQueque() {
	fmt.Print("data : ")
	for _, item := range q.items {
		fmt.Printf("%d ", item)
	}
	fmt.Printf("\n")
}

func (q *Queue) dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	queue := Queue{}
	queue.printQueque()
	queue.enqueue(50)
	queue.enqueue(100)
	queue.printQueque()
	queue.dequeue()
	queue.printQueque()
	queue.dequeue()
	queue.printQueque()
}
