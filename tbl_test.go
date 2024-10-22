package tu

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
)

func TestNewTbl(t *testing.T) {
	var tbl Tbl[int, string]
	a := assert.New(t)
	map_ := map[int]string{1: "1", 2: "2", 3: "3", 4: "4"}

	a.Equal(reflect.Map, reflect.ValueOf(map_).Kind())

	{
		// NewTbl()
		tbl = NewTbl[int, string]()

		a.NotNil(tbl)
		a.Equal(0, len(tbl))
	}
	{
		// NewTbl()
		tbl = NewTbl[int, string](map_)

		a.NotNil(tbl)
		a.Equal(len(map_), len(tbl))

		// NewTbl()
		tbl2 := NewTbl[int, string](tbl)

		a.NotNil(tbl2)
		a.Equal(len(tbl), len(tbl2))

		for i := 0; i < 100; i++ {
			tbl.Put(i, strconv.Itoa(i))
		}

		t.Log(tbl)
	}
}
