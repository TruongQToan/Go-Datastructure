package main

import "math/rand"

// Treap defines a treap
type Treap struct {
	root *TreapNode
	size int
}

// TreapNode defines a treap node
type TreapNode struct {
	key      int
	value    int
	left     *TreapNode
	right    *TreapNode
	parent   *TreapNode
	priority int
}

// Add adds a node to the treap
func (t *Treap) Add(key, value int) {
	u := &TreapNode{key: key, value: value, priority: rand.Int()}
	t.bstAdd(t.root, u)
	t.bubbleUp(u)
}

// Remove removes a node from the treap
func (t *Treap) Remove(x int) bool {
	u := t.FindLast(x)	
	if u != nil && u.key == x {
		t.trickleDown(u)
		t.splice(u)
		return true
	}
	
	return false
}

// FindLast finds the last node on the path from root to the node that contain x (or last node if x is not in the treap)
func (t *Treap) FindLast(x int) *TreapNode {
	w := t.root
	var prev *TreapNode	
	for w != nil {
		prev = w
		if x < w.key {
			w = w.left
		} else if x > w.key {
			w = w.left
		} else {
			return w
		}
	}
	
	return prev
}

func (t *Treap) trickleDown(u *TreapNode) {
	for u.left != nil || u.right != nil {
		if u.left == nil {
			t.rotateLeft(u)
		} else if u.right != nil {
			t.rotateRight(u)
		} else if u.left.priority < u.right.priority {
			t.rotateRight(u)
		} else {
			t.rotateLeft(u)
		}
	}	
	
	if t.root == u {
		t.root = u.parent
	}
}

// splice removes the nodes with at most one child from the treap
func (t *Treap) splice(u *TreapNode) {
	var s, p *TreapNode
	if u.left != nil {
		s = u.left
	} else {
		s = u.right
	}
	
	if u == t.root {
		t.root = s
		p = nil
	} else {
		p = u.parent
		if p.left == u {
			p.left = s
		} else {
			p.right = s
		}
	}
	
	if s != nil {
		s.parent = p
	}
	
	t.size--
}

func (t *Treap) bubbleUp(u *TreapNode) {
	for u.parent != nil && u.parent.priority > u.priority {
		if u.parent.left == u {
			t.rotateRight(u.parent)
		} else {
			t.rotateLeft(u.parent)
		}
	}
	
	if u.parent == nil {
		t.root = u
	}
}

func (t *Treap) bstAdd(v, u *TreapNode) *TreapNode {
	if v == nil {
		v = u
		return v
	}	
	
	if v.key == u.key {
		v.value = u.value
		return v
	}
	
	if u.key < v.key {
		v.left = t.bstAdd(v.left, u)	
	} else {
		v.right = t.bstAdd(v.right, u)
	}
	
	v.left.parent = v
	t.size++
	return v
}

func (t *Treap) rotateRight(u *TreapNode) {
	w := u.left
	if w == nil {
		return
	}
	
	w.parent = u.parent
	if w.parent != nil {
		if w.parent.left == u {
			w.parent.left = w
		} else {
			w.parent.right = w
		}
	}
	
	u.left = w.right
	if u.left != nil {
		u.left.parent = u
	}
	
	u.parent = w
	w.right = u
	if t.root == u {
		t.root = w
		w.parent = nil
	}
}

func (t *Treap) rotateLeft(u *TreapNode) {
	w := u.right
	if w == nil {
		return
	}
	
	w.parent = u.parent
	if w.parent != nil {
		if w.parent.left == u {
			w.parent.left = w
		} else {
			w.parent.right = w
		}
	}
	
	u.right = w.left
	if u.right != nil {
		u.right.parent = u
	}
	
	u.parent = w
	w.left = u
	if t.root == u {
		t.root = w	
		w.parent = nil
	}
}
