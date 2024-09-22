package tu

import (
	"cmp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var vec_expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var vec_rexpected = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

func TestVec_Put(t *testing.T) {
	a := assert.New(t)
	v := NewVec[int]()

	a.True(v.IsEmpty())
	a.Equal(0, v.Len())

	// [7, 8, 9]
	v.Put(-1, 7)
	v.Put(-1, 8)
	v.Put(-1, 9)

	// [1, 2, 3, 7, 8, 9]
	v.Put(0, 3)
	v.Put(0, 2)
	v.Put(0, 1)

	// [1, 2, 3, 4, 5, 6, 7, 8, 9]
	v.Put(3, 4)
	v.Put(4, 5)
	v.Put(5, 6)

	a.False(v.IsEmpty())
	a.Equal(9, v.Len())

	a.Equal(vec_expected, v.items)
}

func TestVec_Pop(t *testing.T) {
	a := assert.New(t)
	v := NewVecFrom[int](vec_expected)

	a.False(v.IsEmpty())
	a.Equal(9, v.Len())

	a.Equal(vec_expected, v.items)

	for i := 1; i < 8; i++ {
		a.Equal(i+1, v.Pop(1))
	}

	a.Equal([]int{1, 9}, v.items)
}

func TestVec_At(t *testing.T) {
	a := assert.New(t)
	v := NewVecFrom[int](vec_expected)

	a.False(v.IsEmpty())
	a.Equal(9, v.Len())

	a.Equal(vec_expected, v.items)

	a.Equal(1, v.At(0))
	a.Equal(9, v.At(-1))

	for i := 1; i < 8; i++ {
		a.Equal(i+1, v.At(i))
	}
}

func TestVec_Re(t *testing.T) {
	a := assert.New(t)
	v := NewVecFrom[int](vec_expected)

	a.False(v.IsEmpty())
	a.Equal(9, v.Len())

	a.Equal(vec_expected, v.items)

	v.Re(0, 9)
	v.Re(-1, 1)

	for i := 1; i < 8; i++ {
		v.Re(i, 9-i)
	}

	a.Equal(vec_rexpected, v.items)
}

func TestVec_Range(t *testing.T) {
	a := assert.New(t)
	v := NewVecFrom[int](vec_expected)

	for i, it := range v.L {
		a.Equal(i+1, it)
	}

	for i, it := range v.R {
		a.Equal(i+1, it)
	}
}

func TestVec_Map(t *testing.T) {
	a := assert.New(t)
	v := NewVecFrom[int](vec_expected)

	a.Equal([]int{2, 4, 6, 8, 10, 12, 14, 16, 18},
		v.Map(func(idx int, it int) int { return it * 2 }).items)
}

func TestVec_Filter(t *testing.T) {
	a := assert.New(t)
	v := NewVecFrom[int](vec_expected)

	a.Equal([]int{2, 4, 6, 8},
		v.Filter(func(idx int, it int) bool { return it%2 == 0 }).items)
}

func TestVec_Any(t *testing.T) {
	a := assert.New(t)
	v := NewVecFrom[int](vec_expected)

	a.True(v.IsAny(func(idx int, it int) bool { return it == 9 }))

	a.False(v.IsAny(func(idx int, it int) bool { return it == 0 }))
}

func TestVec_All(t *testing.T) {
	a := assert.New(t)
	v := NewVecFrom[int](vec_expected)

	a.True(v.IsAll(func(idx int, it int) bool { return it > 0 }))

	a.False(v.IsAll(func(idx int, it int) bool { return it > 1 }))
}

func TestVec_Find(t *testing.T) {
	a := assert.New(t)
	v := NewVecFrom[int](vec_expected)

	i, val := v.Find(func(idx int, it int) bool { return it == 9 })

	a.Equal(i, 8)
	a.Equal(val, 9)
}

func TestVec_Index(t *testing.T) {
	a := assert.New(t)
	v := NewVecFrom[int](vec_expected)

	i := v.Index(func(it int) bool { return it == 9 })

	a.Equal(i, 8)
}

func TestVec_Sort(t *testing.T) {
	a := assert.New(t)
	v := NewVecFrom[int](vec_expected)

	v.Put(-1, v.Pop(0))

	a.False(v.IsSorted(cmp.Compare[int]))
	v.Sort(cmp.Compare[int])
	a.Equal(vec_expected, v.items)
	a.True(v.IsSorted(cmp.Compare[int]))
}

func TestVec_Copy(t *testing.T) {
	a := assert.New(t)
	v := NewVecFrom[int](vec_expected)

	v1 := v.Clone()

	a.Equal(v.Len(), v1.Len())

	a.Equal(v.items, v1.items)
}
