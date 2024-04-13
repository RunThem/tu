package u

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAlnum(t *testing.T) {
	a := assert.New(t)

	a.True(IsAlnum('a'))
	a.True(IsAlnum('z'))
	a.True(IsAlnum('A'))
	a.True(IsAlnum('Z'))
	a.True(IsAlnum('0'))
	a.True(IsAlnum('9'))

	a.False(IsAlnum('!'))
}

func TestIsAlpha(t *testing.T) {
	a := assert.New(t)

	a.True(IsAlpha('a'))
	a.True(IsAlpha('z'))
	a.True(IsAlpha('A'))
	a.True(IsAlpha('Z'))

	a.False(IsAlpha('0'))
}

func TestIsLower(t *testing.T) {
	a := assert.New(t)

	a.True(IsLower('a'))
	a.True(IsLower('z'))

	a.False(IsLower('0'))
}

func TestIsUpper(t *testing.T) {
	a := assert.New(t)

	a.True(IsUpper('A'))
	a.True(IsUpper('Z'))

	a.False(IsUpper('0'))
}

func TestIsDigit(t *testing.T) {
	a := assert.New(t)

	a.True(IsDigit('0'))
	a.True(IsDigit('9'))

	a.False(IsDigit('a'))
}

func TestIsXdigit(t *testing.T) {
	a := assert.New(t)

	a.True(IsXdigit('0'))
	a.True(IsXdigit('9'))

	a.True(IsXdigit('a'))
	a.True(IsXdigit('f'))
	a.True(IsXdigit('A'))
	a.True(IsXdigit('F'))

	a.False(IsXdigit('g'))
}

func TestIsCntrl(t *testing.T) {
	a := assert.New(t)

	a.True(IsCntrl(0x1f))
	a.True(IsCntrl(0x7f))

	a.False(IsCntrl('!'))
}

func TestIsGraph(t *testing.T) {
	a := assert.New(t)

	a.True(IsGraph('!'))
	a.True(IsGraph('~'))

	a.False(IsGraph('\t'))
}

func TestIsSpace(t *testing.T) {
	a := assert.New(t)

	a.True(IsSpace('\t'))
	a.True(IsSpace('\r'))

	a.True(IsSpace(' '))

	a.False(IsSpace('1'))
}

func TestIsBlank(t *testing.T) {
	a := assert.New(t)

	a.True(IsBlank('\t'))
	a.True(IsBlank(' '))

	a.False(IsBlank('1'))
}

func TestIsPrint(t *testing.T) {
	a := assert.New(t)

	a.True(IsPrint(' '))
	a.True(IsPrint('~'))

	a.False(IsPrint('\t'))
}

func TestIsPunct(t *testing.T) {
	a := assert.New(t)

	a.True(IsPunct('!'))
	a.True(IsPunct('/'))

	a.True(IsPunct(':'))
	a.True(IsPunct('@'))

	a.True(IsPunct('['))
	a.True(IsPunct('`'))

	a.True(IsPunct('{'))
	a.True(IsPunct('~'))

	a.False(IsPunct('1'))
}

func TestToLower(t *testing.T) {
	a := assert.New(t)

	a.Equal(byte('a'), ToLower('A'))
	a.Equal(byte('z'), ToLower('Z'))

	a.Equal(byte('a'), ToLower('a'))
	a.Equal(byte('z'), ToLower('z'))

	a.Equal(byte('1'), ToLower('1'))
}

func TestToUpper(t *testing.T) {
	a := assert.New(t)

	a.Equal(byte('A'), ToUpper('a'))
	a.Equal(byte('Z'), ToUpper('z'))

	a.Equal(byte('A'), ToUpper('A'))
	a.Equal(byte('Z'), ToUpper('Z'))

	a.Equal(byte('1'), ToUpper('1'))
}
