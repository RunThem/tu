package u

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAlnum(t *testing.T) {
	assert.True(t, IsAlnum('a'))
	assert.True(t, IsAlnum('z'))
	assert.True(t, IsAlnum('A'))
	assert.True(t, IsAlnum('Z'))
	assert.True(t, IsAlnum('0'))
	assert.True(t, IsAlnum('9'))

	assert.False(t, IsAlnum('!'))
}

func TestIsAlpha(t *testing.T) {
	assert.True(t, IsAlpha('a'))
	assert.True(t, IsAlpha('z'))
	assert.True(t, IsAlpha('A'))
	assert.True(t, IsAlpha('Z'))

	assert.False(t, IsAlpha('0'))
}

func TestIsLower(t *testing.T) {
	assert.True(t, IsLower('a'))
	assert.True(t, IsLower('z'))

	assert.False(t, IsLower('0'))
}

func TestIsUpper(t *testing.T) {
	assert.True(t, IsUpper('A'))
	assert.True(t, IsUpper('Z'))

	assert.False(t, IsUpper('0'))
}

func TestIsDigit(t *testing.T) {
	assert.True(t, IsDigit('0'))
	assert.True(t, IsDigit('9'))

	assert.False(t, IsDigit('a'))
}

func TestIsXdigit(t *testing.T) {
	assert.True(t, IsXdigit('0'))
	assert.True(t, IsXdigit('9'))

	assert.True(t, IsXdigit('a'))
	assert.True(t, IsXdigit('f'))
	assert.True(t, IsXdigit('A'))
	assert.True(t, IsXdigit('F'))

	assert.False(t, IsXdigit('g'))
}

func TestIsCntrl(t *testing.T) {
	assert.True(t, IsCntrl(0x1f))
	assert.True(t, IsCntrl(0x7f))

	assert.False(t, IsCntrl('!'))
}

func TestIsGraph(t *testing.T) {
	assert.True(t, IsGraph('!'))
	assert.True(t, IsGraph('~'))

	assert.False(t, IsGraph('\t'))
}

func TestIsSpace(t *testing.T) {
	assert.True(t, IsSpace('\t'))
	assert.True(t, IsSpace('\r'))

	assert.True(t, IsSpace(' '))

	assert.False(t, IsSpace('1'))
}

func TestIsBlank(t *testing.T) {
	assert.True(t, IsBlank('\t'))
	assert.True(t, IsBlank(' '))

	assert.False(t, IsBlank('1'))
}

func TestIsPrint(t *testing.T) {
	assert.True(t, IsPrint(' '))
	assert.True(t, IsPrint('~'))

	assert.False(t, IsPrint('\t'))
}

func TestIsPunct(t *testing.T) {
	assert.True(t, IsPunct('!'))
	assert.True(t, IsPunct('/'))

	assert.True(t, IsPunct(':'))
	assert.True(t, IsPunct('@'))

	assert.True(t, IsPunct('['))
	assert.True(t, IsPunct('`'))

	assert.True(t, IsPunct('{'))
	assert.True(t, IsPunct('~'))

	assert.False(t, IsPunct('1'))
}

func TestToLower(t *testing.T) {

	assert.Equal(t, byte('a'), ToLower('A'))
	assert.Equal(t, byte('z'), ToLower('Z'))

	assert.Equal(t, byte('a'), ToLower('a'))
	assert.Equal(t, byte('z'), ToLower('z'))

	assert.Equal(t, byte('1'), ToLower('1'))
}

func TestToUpper(t *testing.T) {

	assert.Equal(t, byte('A'), ToUpper('a'))
	assert.Equal(t, byte('Z'), ToUpper('z'))

	assert.Equal(t, byte('A'), ToUpper('A'))
	assert.Equal(t, byte('Z'), ToUpper('Z'))

	assert.Equal(t, byte('1'), ToUpper('1'))
}
