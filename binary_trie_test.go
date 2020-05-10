package main

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkXFastTrie(b *testing.B) {
	start := time.Now()
	trie := NewBinaryTrie()
	var i uint64
	var numSeeds uint64
	numSeeds = 10000
	for i = 1; i <= numSeeds; i++ {
		_ = trie.Add(i)
	}

	for i = numSeeds; i > 0; i-- {
		_ = trie.Find(i)
		_ = trie.Remove(i)
	}

	end := time.Now()
	fmt.Println("xfast trie duration", end.Sub(start).Seconds())
}

func BenchmarkSkiplist(b *testing.B) {
	start := time.Now()

	sklist := CreateSkipList()
	var i uint64
	var numSeeds uint64
	numSeeds = 10000
	for i = 1; i <= numSeeds; i++ {
		sklist.Add(i, int(i))
	}

	for i = numSeeds; i > 0; i-- {
		_, _ = sklist.Get(int(i))
		sklist.Delete(int(i))
	}

	end := time.Now()
	fmt.Println("skiplist duration", end.Sub(start).Seconds())
}
