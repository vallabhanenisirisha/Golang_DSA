package main

import "fmt"

//node
type node struct {
	data int
	next *node
}

//linkedlist
type circularlist struct {
	head *node
	tail *node
}

//push
func (l *circularlist) push(n int) {
	newnode := &node{n, nil}
	if l.head == nil {
		l.head = newnode
		newnode.next = newnode
	} else {
		newnode.next = l.head
		l.tail.next = newnode
	}
	l.tail = newnode
}

//display
func (l *circularlist) display() {
	p := l.head
	for p != l.tail {
		fmt.Print(p.data, "-->")
		p = p.next
	}
	fmt.Print(p.data)
}

//length
func (l *circularlist) len() int {