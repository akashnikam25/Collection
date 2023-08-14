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
	choice := 1 //by default 1 value we set it
	q := &queue{}
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
		fmt.Println("choice", choice)
	}
}
