package u

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var map_expected = map[int]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5"}

func TestMap_Put(t *testing.T) {
	a := assert.New(t)
	v := NewMap[int, string](nil)

	a.True(v.IsEmpty())

	v.Put(1, "1")
	v.Put(2, "2")
	v.Put(3, "3")
	v.Put(4, "4")
	v.Put(5, "5")

	t.Logf("+%v", v.maps)

	a.False(v.IsEmpty())

	a.Equal(map_expected, v.maps)
}
