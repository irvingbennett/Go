// Package binarysearchtree is...
package binarysearchtree

import (
	"fmt"
)

// SearchTreeData is the tree structure
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// Bst returns the item being search
func Bst(i int) *SearchTreeData {
	// fmt.Println("Bst", i)
	b := SearchTreeData{data: i}
	return &b
}

// Insert inserts data into a tree
func (t *SearchTreeData) Insert(i int) {
	// fmt.Println("Insert", i)
	if t.data == i {
		n := SearchTreeData{data: i}
		t.left = &n
		return
	} else if t.data > i {
		if t.left == nil {
			n := SearchTreeData{data: i}
			t.left = &n
		} else {
			t.left.Insert(i)
		}
	} else if t.data < i {
		if t.right == nil {
			n := SearchTreeData{data: i}
			t.right = &n
		} else {
			t.right.Insert(i)
		}
	}
	// t.print()
	// fmt.Println()
	return
}

func (t *SearchTreeData) print() {
	if t.left != nil {
		t.left.print()
	}
	fmt.Print(" ", t.data)
	if t.right != nil {
		t.right.print()
	}
}

// MapString applies a map to a btree
func (t *SearchTreeData) MapString(func(i int) string) []string {
	x := t.string()
	fmt.Println("MapStr", x)
	return x
}

func (t *SearchTreeData) string() []string {
	x := []string{}
	if t.left != nil {
		x = append(x, t.left.string()...)
	}
	x = append(x, fmt.Sprintf("%d", t.data))
	if t.right != nil {
		x = append(x, t.right.string()...)
	}
	return x
}

// MapInt returns a list of integers
func (t *SearchTreeData) MapInt(func(i int) int) []int {
	x := t.int()
	fmt.Println("MapInt", x)

	return x
}

func (t *SearchTreeData) int() []int {
	x := []int{}
	if t.left != nil {
		x = append(x, t.left.int()...)
	}
	x = append(x, t.data)
	if t.right != nil {
		x = append(x, t.right.int()...)
	}
	return x
}
