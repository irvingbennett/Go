// Package tree builds a tree from records
package tree

import (
	"fmt"
	"sort"
)

// Record is a stuct of posts
type Record struct {
	ID     int
	Parent int
}

// Node is an empty struct
type Node struct {
	ID       int
	Children []*Node
}

type ByID []Record

func (this ByID) Len() int {
	return len(this)
}
func (this ByID) Less(i, j int) bool {
	return this[i].ID < this[j].ID
}
func (this ByID) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// Build a tree from records
func Build(in ByID) (*Node, error) {
	if len(in) == 0 {
		fmt.Println("Empty input")
		return nil, nil
	}
	sort.Sort(in)
	nodes := make(map[int]*Node)
	for i, r := range in {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, fmt.Errorf("Bad record %v", i)
		}

		nodes[r.ID] = &Node{ID: r.ID}
		if i > 0 {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
		}
	}
	return nodes[0], nil
}
