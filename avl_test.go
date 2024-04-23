package tu

import (
	"cmp"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvl_Put(t *testing.T) {
	a := assert.New(t)
	avl := NewAvl[int, string](cmp.Compare[int])

	a.True(avl.IsEmpty())
	a.Equal(0, avl.Len())

	for i := 1; i <= 7; i++ {
		avl.Put(i, fmt.Sprintf("%d", i))
	}

	a.False(avl.IsEmpty())
	a.Equal(7, avl.Len())

	for i := 1; i <= 7; i++ {
		a.Equal(fmt.Sprintf("%d", i), avl.At(i))
	}
}

func TestAvl_Pop(t *testing.T) {
	a := assert.New(t)
	avl := NewAvl[int, string](cmp.Compare[int])

	a.True(avl.IsEmpty())
	a.Equal(0, avl.Len())

	for i := 1; i <= 7; i++ {
		avl.Put(i, fmt.Sprintf("%d", i))
	}

	a.False(avl.IsEmpty())
	a.Equal(7, avl.Len())

	for i := 1; i <= 7; i++ {
		a.Equal(fmt.Sprintf("%d", i), avl.Pop(i))
	}
}

func TestAvl_At(t *testing.T) {
	a := assert.New(t)
	avl := NewAvl[int, string](cmp.Compare[int])

	a.True(avl.IsEmpty())
	a.Equal(0, avl.Len())

	for i := 1; i <= 7; i++ {
		avl.Put(i, fmt.Sprintf("%d", i))
	}

	a.False(avl.IsEmpty())
	a.Equal(7, avl.Len())

	for i := 1; i <= 7; i++ {
		a.Equal(fmt.Sprintf("%d", i), avl.At(i))
	}
}

func TestAvl_Re(t *testing.T) {
	a := assert.New(t)
	avl := NewAvl[int, string](cmp.Compare[int])

	a.True(avl.IsEmpty())
	a.Equal(0, avl.Len())

	for i := 1; i <= 7; i++ {
		avl.Put(i, fmt.Sprintf("%d", i))
	}

	for i := 1; i <= 7; i++ {
		avl.Re(i, fmt.Sprintf("-%d", i))
	}

	a.False(avl.IsEmpty())
	a.Equal(7, avl.Len())

	for i := 1; i <= 7; i++ {
		a.Equal(fmt.Sprintf("-%d", i), avl.At(i))
	}
}

func TestAvl_Range(t *testing.T) {
	a := assert.New(t)
	avl := NewAvl[int, string](cmp.Compare[int])

	for i := 1; i <= 7; i++ {
		avl.Put(i, fmt.Sprintf("%d", i))
	}

	i := 1
	for k, v := range avl.Range(true) {
		a.Equal(fmt.Sprintf("%d", k), v)
		a.Equal(fmt.Sprintf("%d", i), v)
		i++
	}

	i = 7
	for k, v := range avl.Range(false) {
		a.Equal(fmt.Sprintf("%d", k), v)
		a.Equal(fmt.Sprintf("%d", i), v)
		i--
	}
}

func TestAvl_Map(t *testing.T) {
	a := assert.New(t)
	avl := NewAvl[int, string](cmp.Compare[int])
	avl1 := NewAvl[int, string](cmp.Compare[int])

	for i := 1; i <= 7; i++ {
		avl.Put(i, fmt.Sprintf("%d", i))
	}

	for i := 1; i <= 7; i++ {
		avl1.Put(i, fmt.Sprintf("-%d", i))
	}

	avl2 := avl.Map(func(key int, val string) string {
		return "-" + val
	})

	a.Equal(avl1.Len(), avl2.Len())

	for k, v := range avl1.Range(true) {
		a.Equal(v, avl2.At(k))
	}
}

func TestAvl_Filter(t *testing.T) {
	a := assert.New(t)
	avl := NewAvl[int, string](cmp.Compare[int])
	avl1 := NewAvl[int, string](cmp.Compare[int])

	for i := 1; i <= 7; i++ {
		avl.Put(i, fmt.Sprintf("%d", i))
	}

	for i := 1; i <= 7; i += 2 {
		avl1.Put(i, fmt.Sprintf("%d", i))
	}

	avl2 := avl.Filter(func(key int, val string) bool {
		return key%2 == 1
	})

	a.Equal(avl1.Len(), avl2.Len())

	for k, v := range avl1.Range(true) {
		a.Equal(v, avl2.At(k))
	}
}

func TestAvl_IsAny(t *testing.T) {
	a := assert.New(t)
	avl := NewAvl[int, string](cmp.Compare[int])

	for i := 1; i <= 7; i++ {
		avl.Put(i, fmt.Sprintf("%d", i))
	}

	a.True(avl.IsAny(func(key int, val string) bool {
		return val == "7"
	}))

	a.False(avl.IsAny(func(key int, val string) bool {
		return val == "8"
	}))
}

func TestAvl_IsAll(t *testing.T) {
	a := assert.New(t)
	avl := NewAvl[int, string](cmp.Compare[int])

	for i := 1; i <= 7; i++ {
		avl.Put(i, fmt.Sprintf("%d", i))
	}

	a.True(avl.IsAll(func(key int, val string) bool {
		return len(val) == 1
	}))

	a.False(avl.IsAny(func(key int, val string) bool {
		return len(val) == 2
	}))
}

func TestAvl_Copy(t *testing.T) {
	a := assert.New(t)
	avl := NewAvl[int, string](cmp.Compare[int])

	for i := 1; i <= 7; i++ {
		avl.Put(i, fmt.Sprintf("%d", i))
	}

	avl1 := avl.Copy()

	a.Equal(avl.Len(), avl1.Len())

	for k, v := range avl.Range(true) {
		a.Equal(v, avl1.At(k))
	}
}
