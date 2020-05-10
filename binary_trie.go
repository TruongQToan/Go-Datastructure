package main

// BinaryTrieNode defines a binary trie node
type BinaryTrieNode struct {
	child []*BinaryTrieNode
	jump  *BinaryTrieNode
	value uint64
	parent *BinaryTrieNode
}

func NewBinaryTrieNode() *BinaryTrieNode {
	return &BinaryTrieNode{
		child : make([]*BinaryTrieNode, 2),
	}
}

type BinaryTrie struct {
	w    int
	root *BinaryTrieNode
}

func NewBinaryTrie() *BinaryTrie {
	return &BinaryTrie{
		w: 64,
		root: &BinaryTrieNode{
			child: make([]*BinaryTrieNode, 2),
		},
	}
}

func (trie *BinaryTrie) Find(x uint64) int64 {
	var (
		i int
		c uint64
	)

	u := trie.root
	for i = 0; i < trie.w; i++ {
		c = (x >> (trie.w - i - 1)) & 1
		if u.child[c] == nil {
			break
		}

		u = u.child[c]
	}

	if i == trie.w {
		return int64(u.value)
	}

	u = u.jump
	if c == 1 {
		u = u.child[c]
	}

	if u != nil {
		return int64(u.value)
	}

	return -1
}

func (trie *BinaryTrie) Remove(x uint64) bool {
	var (
		i int
		c uint64
	)

	u := trie.root
	for i = 0; i < trie.w; i++ {
		c = (x >> (trie.w-i-1)) & 1
		if u.child[c] == nil {
			return false
		}

		u = u.child[c]
	}

	if u.child[0] != nil {
		u.child[0].child[1] = u.child[1]
	}

	if u.child[1] != nil {
		u.child[1].child[0] = u.child[0]
	}

	v := u
	for i = trie.w-1; i >= 0; i -= 1 {
		c = (x >> (trie.w-i-1)) & 1
		v = v.parent
		v.child[c] = nil
		if v.child[1-c] != nil {
			break
		}
	}

	for ; i >= 0; i -= 1 {
		c = (x >> (trie.w-i-1)) & 1
		if v.jump == u {
			v.jump = u.child[1-c]
		}

		v = v.parent
	}

	u = nil
	return true
}

func (trie *BinaryTrie) Add(x uint64) bool {
	var (
		i int
		c uint64
	)

	u := trie.root
	for i = 0; i < trie.w; i++ {
		c = (x >> (trie.w-i-1)) & 1
		if u.child[c] == nil {
			break
		}

		u = u.child[c]
	}

	if i == trie.w {
		return false
	}

	var pred *BinaryTrieNode
	if c == 1 {
		pred = u.jump
	} else if u.jump != nil {
		pred = u.jump.child[0]
	}

	for ; i < trie.w; i++ {
		c = (x >> (trie.w-i-1))	 & 1
		u.child[c] = NewBinaryTrieNode()
		u.child[c].parent = u
		u = u.child[c]
	}

	u.value = x
	if pred != nil {
		u.child[1] = pred.child[1]
		pred.child[1] = u
	}

	u.child[0] = pred
	if u.child[1] != nil {
		u.child[1].child[0] = u
	}

	v := u.parent
	for v != nil {
		if ((v.child[0] == nil && (v.jump == nil || v.jump.value > x)) ||
			(v.child[1] == nil && (v.jump == nil || v.jump.value < x))) {
			v.jump = u
		}

		v = v.parent
	}

	return true
}
