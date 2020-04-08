// Package linkedlist defines a linked list
package linkedlist

import (
	"fmt"
)

// Element of a linked list
type Element struct {
	data int
	next *Element
}

// List is the list proper
type List struct {
	head *Element
	size int
}

// New creates a new linked list from a list of integers
func New(list []int) *List {
	// fmt.Println("New:", list)
	l := List{nil, 0}
	if len(list) == 0 {
		return &l
	}
	linked := make([]Element, len(list))
	for x, d := range list {
		element := Element{d, nil}
		linked[x] = element
		if x > 0 {
			linked[x-1].next = &linked[x]
		}
	}
	// fmt.Println(linked)
	l.size = len(list)
	l.head = &linked[0]
	// fmt.Println("Done new:", list)
	return &l
}

// Size returns the size of a list
func (l *List) Size() int {
	return l.size
}

// Push pushes a new element into a list
func (l *List) Push(i int) {
	if l == nil {
		l = &List{nil, 0}
	}
	l.size++
	// fmt.Println("Push", l, i)
	var p = &Element{}
	if l.head == nil {
		l.head = &Element{i, nil}
	} else {
		p = l.head
		for p.next != nil {
			p = p.next
		}
	}
	p.next = &Element{i, nil}
}

// Pop pops an element from a list
func (l *List) Pop() (i int, err error) {
	// fmt.Println("Pop", l)
	if l.head == nil {
		err = fmt.Errorf("empty list")
		return
	}
	p := l.head
	var last *Element
	for p.next != nil {
		last = p
		p = p.next
	}

	i = p.data
	last.next = nil
	// fmt.Println(p)
	l.size--

	return
}

// Array returns and array with all elements in the list
func (l *List) Array() (a []int) {
	// fmt.Println("Array")
	if l.head == nil {
		return
	}
	tmp := l.head
	for {
		// fmt.Println(a)
		a = append(a, tmp.data)
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
	return
}

// Reverse reverses the list
func (l *List) Reverse() *List {
	// fmt.Println("Array")
	if l.size == 0 {
		return l
	}
	a := l.Array()
	l = &List{}
	for x := len(a) - 1; x >= 0; x-- {
		l.Push(a[x])
	}
	return l
}
