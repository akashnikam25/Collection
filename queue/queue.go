package main

import (
	"fmt"
	"os"
)

type queue struct {
	items []int
}

func (q *queue) enqueue(val int) {
	q.items = append(q.items, val)
}
func (q *queue) dequeue() {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
	} else {
		fmt.Println("dequeue element is", q.items[0])
		q.items = q.items[1:]
	}
}
func (q *queue) print() {
	if len(q.items) != 0 {
		for _, val := range q.items {
			fmt.Print(val, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("queue is empty")
	}

}
func enterChoice() (choice int) {
	fmt.Println("0.Exit")
	fmt.Println("1.Add element")
	fmt.Println("2.Remove element")
	fmt.Println("3.Print")
	fmt.Println("Enter choice")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return
}
func enterElement() (element int) {
	_, err := fmt.Scan(&element)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return
}
func main() {
	// choice := 1 //by default 1 value we set it
	// q := &queue{}
	// for choice != 0 {
	// 	switch choice {
	// 	case 1:
	// 		fmt.Println("Enter element to add")
	// 		q.enqueue(enterElement())
	// 	case 2:
	// 		q.dequeue()
	// 	case 3:
	// 		fmt.Println("Print all queue")
	// 		q.print()
	// 	}
	// 	choice = enterChoice()
	// 	fmt.Println("choice", choice)
	// }
	choice := 1 //by default 1 value we set it
	q := newCircularQueue(5)
	for choice != 0 {
		switch choice {
		case 1:
			fmt.Println("Enter element to add")
			q.enqueue(enterElement())
		case 2:
			q.dequeue()
		case 3:
			fmt.Println("Print all queue")
			q.print()
		}
		choice = enterChoice()
		//fmt.Println("choice", choice)
	}
}

type circularQueue struct {
	rear  int
	front int
	size  int
	arr   []int
}

func newCircularQueue(size int) *circularQueue {

	if size == 0 {
		return nil
	}
	return &circularQueue{
		rear:  -1,
		front: -1,
		size:  size,
		arr:   make([]int, size),
	}
}

// Enqueue puts a element in the rear of queue
func (q *circularQueue) enqueue(val int) {
	if q.front == -1 {
		q.rear = 0
		q.front = 0
		q.arr[q.rear] = val
	} else if q.front != 0 && q.rear == q.size-1 {
		q.rear = 0
		q.arr[q.rear] = val
	} else if (q.front == 0 && q.rear == q.size-1) || ((q.rear + 1 %q.size) == q.front) {
		fmt.Println("queue is full")
	} else {
		q.rear++
		q.arr[q.rear] = val
	}
}

// dequeue removes element from front 
func (q *circularQueue) dequeue() {
	if q.front == -1 {
		fmt.Println("queue is empty")
		return
	}
	fmt.Println("deleted element is", q.arr[q.front])
	if q.front == q.rear {
		q.front, q.rear = -1, -1
	} else if q.front == q.size-1 {
		q.front = 0
	} else {
		q.front = q.front + 1
	}

}

// prints the queue
func (q *circularQueue) print() {
	if q.front <= q.rear {
		for i := q.front; i <= q.rear; i++ {
			fmt.Print(q.arr[i], " ")
		}
	} else {
		for i := q.front; i < q.size; i++ {
			fmt.Print(q.arr[i], " ")
		}
		for i := 0; i <= q.rear; i++ {
			fmt.Print(q.arr[i], " ")
		}

	}
	fmt.Println()
}
