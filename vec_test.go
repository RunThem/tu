package tu

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestVec_NewVec(t *testing.T) {
	var vec Vec[int]
	a := assert.New(t)
	array := [4]int{1, 2, 3, 4}
	slice := []int{1, 2, 3, 4}

	a.Equal(reflect.Array, reflect.ValueOf(array).Kind())
	a.Equal(reflect.Slice, reflect.ValueOf(slice).Kind())

	{ // NewVec()
		vec = NewVec[int]()

		a.NotNil(vec)
		a.Equal(0, len(vec))
		a.Equal(0, vec.Len())
		a.Equal(8, cap(vec))
		a.Equal(8, vec.Cap())
	}
	{
		vec = NewVec[int](4)

		a.NotNil(vec)
		a.Equal(0, len(vec))
		a.Equal(0, vec.Len())
		a.Equal(4, cap(vec))
		a.Equal(4, vec.Cap())

		vec = NewVec[int](128)

		a.NotNil(vec)
		a.Equal(0, len(vec))
		a.Equal(0, vec.Len())
		a.Equal(128, cap(vec))
		a.Equal(128, vec.Cap())
	}
	{
		// array
		vec = NewVec[int](array)

		a.NotNil(vec)
		a.Equal(len(array), len(vec))
		a.Equal(len(array), vec.Len())
		a.Equal(cap(array), cap(vec))
		a.Equal(cap(array), vec.Cap())

		a.Equal(array, [4]int(vec))

		// slice
		vec = NewVec[int](slice)

		a.NotNil(vec)
		a.Equal(len(slice), len(vec))
		a.Equal(len(slice), vec.Len())
		a.Equal(cap(slice), cap(vec))
		a.Equal(cap(slice), vec.Cap())

		a.Equal(slice, []int(vec))

		// Vec
		vec2 := NewVec[int](vec)

		a.NotNil(vec2)
		a.Equal(len(vec), len(vec2))
		a.Equal(len(vec), vec2.Len())
		a.Equal(cap(vec), cap(vec2))
		a.Equal(cap(vec), vec2.Cap())

		a.Equal(vec, vec2)
	}
}

func TestVec_Len(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		vec Vec[int]
		len int
	}{
		{Vec[int]{}, 0},
		{Vec[int]{1}, 1},
		{Vec[int]{1, 2, 3}, 3},
		{Vec[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
	}

	for _, tt := range tests {
		a.Equal(tt.len, tt.vec.Len())
	}
}

func TestVec_Cap(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		vec Vec[int]
		cap int
	}{
		{Vec[int]{}, 1},
		{Vec[int]{1}, 2},
		{Vec[int]{1, 2, 3}, 6},
		{Vec[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 20},
	}

	for _, tt := range tests {
		tt.vec.Put(0)

		a.Equal(tt.cap, tt.vec.Cap())
	}
}

func TestVec_Put(t *testing.T) {
	a := assert.New(t)

	vec := Vec[int]{}

	a.True(vec.IsEmpty())
	a.Equal(0, vec.Len())
	a.Equal(0, vec.Cap())

	tt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range tt {
		vec.Put(v)
	}

	a.False(vec.IsEmpty())
	a.Equal(len(tt), vec.Len())

	a.Equal(tt, []int(vec))
}

func TestVec_Pop(t *testing.T) {
	a := assert.New(t)

	vec := Vec[int]{1, 2, 3, 4, 5}
	a.Equal(5, vec.Pop())
	a.Equal(4, vec.Len())

	a.Equal(4, vec.Pop())
	a.Equal(3, vec.Len())

	a.Equal(3, vec.Pop())
	a.Equal(2, vec.Len())

	a.Equal(2, vec.Pop())
	a.Equal(1, vec.Len())

	a.Equal(1, vec.Pop())
	a.Equal(0, vec.Len())
	a.True(vec.IsEmpty())

	a.Equal(5, vec.Cap())
}

func TestVec_Ins(t *testing.T) {

}
