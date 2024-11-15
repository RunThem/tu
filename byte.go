package tu

// C libc <ctype.h>: https://en.cppreference.com/w/c/string/byte

// IsAlnum checks if ch character is alphanumeric
func IsAlnum(ch byte) bool {
	return IsLower(ch) || IsUpper(ch) || IsDigit(ch)
}

// IsAlpha checks if ch character is alphabetic
func IsAlpha(ch byte) bool {
	return IsLower(ch) || IsUpper(ch)
}

// IsLower checks if ch character is lowercase character
func IsLower(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

// IsUpper checks if ch character is uppercase character
func IsUpper(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}

// IsDigit checks if ch character is digit
func IsDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// IsXdigit checks if ch character is hexadecimal character
func IsXdigit(ch byte) bool {
	return IsDigit(ch) || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')
}

// IsCntrl checks if ch character is control character
func IsCntrl(ch byte) bool {
	return ch <= 0x1f || ch == 0x7f
}

// IsGraph checks if ch character is graphical character
func IsGraph(ch byte) bool {
	return ch >= '!' && ch <= '~'
}

// IsSpace checks if ch character is space character
func IsSpace(ch byte) bool {
	return (ch >= '\t' && ch <= '\r') || ch == ' '
}

// IsBlank checks if ch character is blank character
func IsBlank(ch byte) bool {
	return ch == '\t' || ch == ' '
}

// IsPrint checks if ch character is printing character
func IsPrint(ch byte) bool {
	return IsGraph(ch) || ch == ' '
}

// IsPunct checks if ch character is punctuation character
func IsPunct(ch byte) bool {
	return (ch >= '!' && ch <= '/') || (ch >= ':' && ch <= '@') || (ch >= '[' && ch <= '`') || (ch >= '{' && ch <= '~')
}

// ToLower converts a character to lowercase
func ToLower(ch byte) byte {
	if IsLower(ch) || !IsAlpha(ch) {
		return ch
	}

	return ch + 0x20
}

// ToUpper converts a character to uppercase
func ToUpper(ch byte) byte {
	if IsUpper(ch) || !IsAlpha(ch) {
		return ch
	}

	return ch - 0x20
}
