package tu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var map_expected = map[int]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7"}

func TestMap_Put(t *testing.T) {
	a := assert.New(t)
	m := NewMap[int, string](nil)

	a.True(m.IsEmpty())
	a.Equal(0, m.Len())

	for i := 1; i <= 7; i++ {
		m.Put(i, fmt.Sprintf("%d", i))
	}

	a.False(m.IsEmpty())
	a.Equal(7, m.Len())

	a.Equal(map_expected, m.items)
}

func TestMap_Pop(t *testing.T) {
	a := assert.New(t)
	m := NewMap[int, string](nil)

	for i := 1; i <= 7; i++ {
		m.Put(i, fmt.Sprintf("%d", i))
	}

	for k, v := range m.Range() {
		a.Equal(v, map_expected[k])
	}
}
