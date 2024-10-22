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
