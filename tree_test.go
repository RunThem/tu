package tu

import (
	"cmp"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Put(t *testing.T) {
	a := assert.New(t)
	tree := NewTree[int, string](cmp.Compare[int])

	a.True(tree.IsEmpty())
	a.Equal(0, tree.Len())

	for i := 1; i <= 7; i++ {
		tree.Put(i, fmt.Sprintf("%d", i))
	}

	a.False(tree.IsEmpty())
	a.Equal(7, tree.Len())

	for i := 1; i <= 7; i++ {
		a.Equal(fmt.Sprintf("%d", i), tree.At(i))
	}
}

func TestTree_Pop(t *testing.T) {
	a := assert.New(t)
	tree := NewTree[int, string](cmp.Compare[int])

	a.True(tree.IsEmpty())
	a.Equal(0, tree.Len())

	for i := 1; i <= 7; i++ {
		tree.Put(i, fmt.Sprintf("%d", i))
	}

	a.False(tree.IsEmpty())
	a.Equal(7, tree.Len())

	for i := 1; i <= 7; i++ {
		a.Equal(fmt.Sprintf("%d", i), tree.Pop(i))
	}
}

func TestTree_At(t *testing.T) {
	a := assert.New(t)
	tree := NewTree[int, string](cmp.Compare[int])

	a.True(tree.IsEmpty())
	a.Equal(0, tree.Len())

	for i := 1; i <= 7; i++ {
		tree.Put(i, fmt.Sprintf("%d", i))
	}

	a.False(tree.IsEmpty())
	a.Equal(7, tree.Len())

	for i := 1; i <= 7; i++ {
		a.Equal(fmt.Sprintf("%d", i), tree.At(i))
	}
}

func TestTree_Re(t *testing.T) {
	a := assert.New(t)
	tree := NewTree[int, string](cmp.Compare[int])

	a.True(tree.IsEmpty())
	a.Equal(0, tree.Len())

	for i := 1; i <= 7; i++ {
		tree.Put(i, fmt.Sprintf("%d", i))
	}

	for i := 1; i <= 7; i++ {
		tree.Re(i, fmt.Sprintf("-%d", i))
	}

	a.False(tree.IsEmpty())
	a.Equal(7, tree.Len())

	for i := 1; i <= 7; i++ {
		a.Equal(fmt.Sprintf("-%d", i), tree.At(i))
	}
}

func TestTree_Range(t *testing.T) {
	a := assert.New(t)
	tree := NewTree[int, string](cmp.Compare[int])

	for i := 1; i <= 7; i++ {
		tree.Put(i, fmt.Sprintf("%d", i))
	}

	i := 1
	for k, v := range tree.Range(true) {
		a.Equal(fmt.Sprintf("%d", k), v)
		a.Equal(fmt.Sprintf("%d", i), v)
		i++
	}

	i = 7
	for k, v := range tree.Range(false) {
		a.Equal(fmt.Sprintf("%d", k), v)
		a.Equal(fmt.Sprintf("%d", i), v)
		i--
	}
}

func TestTree_Map(t *testing.T) {
	a := assert.New(t)
	tree := NewTree[int, string](cmp.Compare[int])
	tree1 := NewTree[int, string](cmp.Compare[int])

	for i := 1; i <= 7; i++ {
		tree.Put(i, fmt.Sprintf("%d", i))
	}

	for i := 1; i <= 7; i++ {
		tree1.Put(i, fmt.Sprintf("-%d", i))
	}

	tree2 := tree.Map(func(key int, val string) string {
		return "-" + val
	})

	a.Equal(tree1.Len(), tree2.Len())

	for k, v := range tree1.Range(true) {
		a.Equal(v, tree2.At(k))
	}
}

func TestTree_Filter(t *testing.T) {
	a := assert.New(t)
	tree := NewTree[int, string](cmp.Compare[int])
	tree1 := NewTree[int, string](cmp.Compare[int])

	for i := 1; i <= 7; i++ {
		tree.Put(i, fmt.Sprintf("%d", i))
	}

	for i := 1; i <= 7; i += 2 {
		tree1.Put(i, fmt.Sprintf("%d", i))
	}

	tree2 := tree.Filter(func(key int, val string) bool {
		return key%2 == 1
	})

	a.Equal(tree1.Len(), tree2.Len())

	for k, v := range tree1.Range(true) {
		a.Equal(v, tree2.At(k))
	}
}

func TestTree_IsAny(t *testing.T) {
	a := assert.New(t)
	tree := NewTree[int, string](cmp.Compare[int])

	for i := 1; i <= 7; i++ {
		tree.Put(i, fmt.Sprintf("%d", i))
	}

	a.True(tree.IsAny(func(key int, val string) bool {
		return val == "7"
	}))

	a.False(tree.IsAny(func(key int, val string) bool {
		return val == "8"
	}))
}

func TestTree_IsAll(t *testing.T) {
	a := assert.New(t)
	tree := NewTree[int, string](cmp.Compare[int])

	for i := 1; i <= 7; i++ {
		tree.Put(i, fmt.Sprintf("%d", i))
	}

	a.True(tree.IsAll(func(key int, val string) bool {
		return len(val) == 1
	}))

	a.False(tree.IsAny(func(key int, val string) bool {
		return len(val) == 2
	}))
}

func TestTree_Copy(t *testing.T) {
	a := assert.New(t)
	tree := NewTree[int, string](cmp.Compare[int])

	for i := 1; i <= 7; i++ {
		tree.Put(i, fmt.Sprintf("%d", i))
	}

	tree1 := tree.Copy()

	a.Equal(tree.Len(), tree1.Len())

	for k, v := range tree.Range(true) {
		a.Equal(v, tree1.At(k))
	}
}
