// Package linkedlist ...
package linkedlist

import (
	"fmt"
)

// Node on a double linked list
type Node struct {
	// Val holds the value of the list
	Val   interface{}
	next  *Node
	prev  *Node
	first *Node
	last  *Node
}

// ErrEmptyList is returned for an empty list
var ErrEmptyList = fmt.Errorf("empty list error")

// First returns the first value of a node list
func (e *Node) First() *Node {
	// fmt.Println("Node, First()")
	return e.first
}

// Last returns the last value of a node list
func (e *Node) Last() *Node {
	// fmt.Println("Node, Last()")
	return e.last
}

// Next returns the next node
func (e *Node) Next() *Node {
	return e.next
}

// Prev returns the prev node
func (e *Node) Prev() *Node {
	return e.prev
}

// List of nodes
type List struct {
	head *Node
	tail *Node
}

// PushFront pushes a new value to the front of a list
func (l *List) PushFront(v interface{}) {
	// fmt.Println("List, PushFront", v, l)
	n := Node{Val: v}
	// fmt.Println(n, l.head)
	if l.head != nil {
		n.next = l.head
		l.head.prev = &n
	}
	l.head = &n
	if l.tail == nil {
		l.tail = &n
	}
	if l.tail != l.head {
		l.tail.next = nil
	}
	// fmt.Println(l)
	// fmt.Println()
}

// NewList creates a new list
func NewList(args ...interface{}) *List {
	// fmt.Println("NewList", args)
	l := new(List)
	// fmt.Println(l)
	for i := len(args) - 1; i >= 0; i-- {
		l.PushFront(args[i])
	}
	// fmt.Println(l)
	// fmt.Println()
	return l
}

// First return head of list
func (l List) First() *Node {
	return l.head
}

// Last return head of list
func (l List) Last() *Node {
	return l.tail
}

// PopFront returns an item from the front of a list
func (l *List) PopFront() (v interface{}, err error) {
	n := l.First()
	if n == nil {
		err = ErrEmptyList
		// fmt.Println("PopFront empty list", v, err)
		return
	}

	v = n.Val
	// fmt.Println("PopFront", v)
	// fmt.Printf("Val: %v, Next: %v, Prev: %v, head: %v, tail: %v\n",
	// 	n.Val, n.next, n.prev, l.head, l.tail)
	l.head = n.Next()
	if l.head != nil {
		l.head.prev = nil
	}
	if l.tail == n {
		l.tail = nil
	}
	// fmt.Println()
	return
}

// PushBack pushes a value to the tail of a list
func (l *List) PushBack(v interface{}) {
	// fmt.Println("PushBack", l.tail)
	n := Node{Val: v}
	if l.tail != nil {
		last := l.tail
		last.next = &n
		n.prev = l.tail
	}
	if l.head == nil {
		l.head = &n
	}
	l.tail = &n
	// fmt.Println(l.tail)
	// fmt.Println()
}

// PopBack return the value of the tail of list and removes the node
func (l *List) PopBack() (v interface{}, err error) {
	err = nil
	if l.tail == nil {
		err = ErrEmptyList
		// fmt.Println("Popback empty list", v, l.tail, err)
		return
	}
	last := l.tail
	v = last.Val
	if l.tail == l.head {
		// fmt.Printf("-----------------\n  Last Element\n-----------------\n")
		l.head = nil
		l.tail = nil
		return
	}
	// fmt.Println("Popback", l.tail)
	// fmt.Printf("Val: %v, Next: %v, Prev: %v, head: %v, tail: %v\n",
	//	v, last.next, last.prev, l.head, l.tail)
	if last.prev != nil {
		l.tail = last.prev
		l.tail.next = nil
	}
	// fmt.Println("PopBack", v)
	// fmt.Println()
	return
}

// Reverse a linked list
func (l *List) Reverse() {
	fmt.Println("Reverse", l)
	temp := make([]interface{}, 0)
	for l.tail != nil {
		v, _ := l.PopFront()
		temp = append(temp, v)
	}
	for _, v := range temp {
		l.PushFront(v)
	}
	fmt.Println()
	return
}
