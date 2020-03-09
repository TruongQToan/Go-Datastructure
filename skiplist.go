package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Node represents a node in skiplist
type Node struct {
	Data   interface{}
	Length []int
	Next   []*Node
}

// Skiplist ...
type Skiplist struct {
	Sentinel *Node
	Height   int
}

func createNewNode(data interface{}, height int) *Node {
	return &Node{
		Data:   data,
		Length: make([]int, height),
		Next:   make([]*Node, height),
	}
}

// CreateSkipList creates the skiplist, starting index is 1
func CreateSkipList() *Skiplist {
	return &Skiplist{
		Sentinel: createNewNode(nil, 0),
		Height:   0,
	}
}

// Add adds an node to position i
func (s *Skiplist) Add(data interface{}, i int) {
	newHeight := s.getRandomHeight()
	newNode := createNewNode(data, newHeight)
	node := s.Sentinel
	h := s.Height
	l := 0
	for h > 0 {
		if next, length := node.Next[h-1], node.Length[h-1]; next != nil && l+length < i {
			node = next
			l += length
		} else {
			if h > newHeight {
				if node.Next[h-1] != nil {
					node.Length[h-1]++
				}
			} else {
				newNode.Next[h-1] = node.Next[h-1]
				node.Next[h-1] = newNode
				if node.Length[h-1] != 0 {
					newNode.Length[h-1] = node.Length[h-1] - i + l
				}
				node.Length[h-1] = i - l
			}

			h--
		}
	}

	for h := s.Height; h < newHeight; h++ {
		s.Sentinel.Next = append(s.Sentinel.Next, newNode)
		s.Sentinel.Length = append(s.Sentinel.Length, i)
	}

	if newHeight > s.Height {
		s.Height = newHeight
	}
}

// Get return the node of position i
func (s *Skiplist) Get(i int) (interface{}, error) {
	if prev := s.findPrevNode(i); prev != nil {
		return prev.Next[0].Data, nil
	}

	return nil, fmt.Errorf("cannot get the node of position i")
}

// Delete deletes the node of position i
func (s *Skiplist) Delete(i int) {
	prev := s.findPrevNode(i)
	if prev.Next[0] == nil {
		return
	}

	deleteNode := prev.Next[0]
	deleteNodeHeight := len(deleteNode.Length)

	node := s.Sentinel
	h := s.Height
	l := 0
	for h > 0 {
		if next, length := node.Next[h-1], node.Length[h-1]; next != nil && l+length < i {
			node = next
			l += length
		} else {
			if h > deleteNodeHeight {
				if node.Length[h-1] > 0 {
					node.Length[h-1]--
				}
			} else {
				node.Length[h-1] += node.Next[h-1].Length[h-1] - 1
				node.Next[h-1] = node.Next[h-1].Next[h-1]
				if node.Next[h-1] == nil {
					node.Length[h-1] = 0
				}
			}

			h--
		}
	}

	for s.Height > 0 && s.Sentinel.Next[s.Height-1] == nil {
		s.Height--
		s.Sentinel.Next = s.Sentinel.Next[:s.Height]
		s.Sentinel.Length = s.Sentinel.Length[:s.Height]
	}
}

func (s *Skiplist) getRandomHeight() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	z := r.Uint64()
	h := 1
	for z&1 != 0 {
		z <<= 1
		h++
	}

	return h
}

func (s *Skiplist) findPrevNode(i int) *Node {
	node := s.Sentinel
	h := s.Height
	l := 0
	for h > 0 {
		if node.Next[h-1] != nil && l+node.Length[h-1] < i {
			l = l + node.Length[h-1]
			node = node.Next[h-1]
		} else {
			h--
		}
	}

	return node
}
