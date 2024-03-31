package u

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var rexpected = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

func TestPut(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int]()

	a.True(v.Empty())
	a.Equal(0, v.Size())

	// [7, 8, 9]
	v.PutBack(7)
	v.PutBack(8)
	v.PutBack(9)

	// [1, 2, 3, 7, 8, 9]
	v.PutFront(3)
	v.PutFront(2)
	v.PutFront(1)

	// [1, 2, 3, 4, 5, 6, 7, 8, 9]
	v.Put(3, 4)
	v.Put(4, 5)
	v.Put(5, 6)

	a.False(v.Empty())
	a.Equal(9, v.Size())

	a.Equal(expected, v.items)
}

func TestPop(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int](expected...)

	a.False(v.Empty())
	a.Equal(9, v.Size())

	a.Equal(expected, v.items)

	for i := 1; i < 8; i++ {
		a.Equal(i+1, v.Pop(1))
	}

	a.Equal([]int{1, 9}, v.items)
}

func TestAt(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int](expected...)

	a.False(v.Empty())
	a.Equal(9, v.Size())

	a.Equal(expected, v.items)

	a.Equal(1, v.AtFront())
	a.Equal(9, v.AtBack())

	for i := 1; i < 8; i++ {
		a.Equal(i+1, v.At(i))
	}
}

func TestRe(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int](expected...)

	a.False(v.Empty())
	a.Equal(9, v.Size())

	a.Equal(expected, v.items)

	v.ReFront(9)
	v.ReBack(1)

	for i := 1; i < 8; i++ {
		v.Re(i, 9-i)
	}

	a.Equal(rexpected, v.items)
}

func TestRange(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int](expected...)

	for i, it := range v.Range(true) {
		a.Equal(i+1, it)
	}

	for i, it := range v.Range(false) {
		a.Equal(i+1, it)
	}
}

func TestMap(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int](expected...)

	a.Equal([]int{2, 4, 6, 8, 10, 12, 14, 16, 18},
		v.Map(func(idx int, it int) int { return it * 2 }).items)
}

func TestFilter(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int](expected...)

	a.Equal([]int{2, 4, 6, 8},
		v.Filter(func(idx int, it int) bool { return it%2 == 0 }).items)
}

func TestAny(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int](expected...)

	a.True(v.Any(func(idx int, it int) bool { return it == 9 }))

	a.False(v.Any(func(idx int, it int) bool { return it == 0 }))
}

func TestAll(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int](expected...)

	a.True(v.All(func(idx int, it int) bool { return it > 0 }))

	a.False(v.All(func(idx int, it int) bool { return it > 1 }))
}

func TestFind(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int](expected...)

	i, val := v.Find(func(idx int, it int) bool { return it == 9 })

	a.Equal(i, 8)
	a.Equal(val, 9)
}

func TestIndex(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int](expected...)

	i := v.Index(func(it int) bool { return it == 9 })

	a.Equal(i, 8)
}

func TestSort(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int](expected...)

	cmp := func(a, b int) int {
		if a > b {
			return -1
		} else if a < b {
			return 1
		}

		return 0
	}

	a.False(v.IsSorted(cmp))

	v.Sort(cmp)

	a.Equal(rexpected, v.items)

	a.True(v.IsSorted(cmp))
}
