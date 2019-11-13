package skiplist

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func genMockSkiplist() *Skiplist {
	skiplist := CreateSkipList()
	node1 := createNewNode(1, 1)
	node2 := createNewNode(2, 3)
	node3 := createNewNode(3, 2)
	skiplist.Sentinel.Next = append(skiplist.Sentinel.Next, node1, node2, node2)
	skiplist.Sentinel.Length = append(skiplist.Sentinel.Length, 1, 2, 3)
	skiplist.Height = 3

	node1.Next[0] = node2
	node1.Length[0] = 1

	node2.Next[0] = node3
	node2.Length[0] = 1
	node2.Next[1] = node3
	node2.Length[1] = 1
	node2.Next[2] = nil
	node2.Length[2] = 0

	node3.Next[1] = nil
	node3.Length[1] = 0
	node3.Next[0] = nil
	node3.Length[0] = 0

	return skiplist
}

func TestCreateSkiplist(t *testing.T) {
	skiplist := CreateSkipList()
	require.NotNil(t, skiplist)
	require.NotNil(t, skiplist.Sentinel)
	require.Equal(t, 0, skiplist.Height)
}

func TestGetRandomHeight(t *testing.T) {
	skiplist := CreateSkipList()
	newHeight := skiplist.getRandomHeight()
	require.True(t, newHeight > 0)
}

func TestAddElement(t *testing.T) {
	for n := 0; n < 1000; n++ {
		skiplist := CreateSkipList()
		mockElements := make([]int, 100)
		for i := 1; i <= 100; i++ {
			mockElements[i-1] = rand.Intn(100)
			skiplist.Add(mockElements[i-1], i)
		}

		height := skiplist.Height
		for i := 0; i < height; i++ {
			require.NotNil(t, skiplist.Sentinel.Next[i])
		}
		require.True(t, skiplist.Height > 0)

		for i := 0; i < 100; i++ {
			element, err := skiplist.Get(i + 1)
			require.NoError(t, err)
			require.Equal(t, mockElements[i], element)
		}
	}
}

func TestFindPrevNOde(t *testing.T) {
	skiplist := genMockSkiplist()
	prev1 := skiplist.findPrevNode(1)
	prev2 := skiplist.findPrevNode(2)
	prev3 := skiplist.findPrevNode(3)
	require.Equal(t, skiplist.Sentinel.Next[0], prev2)
	require.Equal(t, skiplist.Sentinel.Next[0].Next[0], prev3)
	require.Equal(t, skiplist.Sentinel, prev1)
}

func TestDeleteElement(t *testing.T) {
	for n := 0; n < 1000; n++ {
		skiplist := CreateSkipList()
		mockElements := make([]int, 100)
		for i := 1; i <= 100; i++ {
			mockElements[i-1] = rand.Intn(100)
			skiplist.Add(mockElements[i-1], i)
		}

		for i := 0; i < 100; i++ {
			r := rand.Intn(100-i) + 1
			skiplist.Delete(r)
			mockElements := append(mockElements[:r-1], mockElements[r:]...)
			if r < 100-i {
				element, err := skiplist.Get(r)
				require.NoError(t, err)
				require.Equal(t, mockElements[r-1], element)
			}
		}
	}
}
