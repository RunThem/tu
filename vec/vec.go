package vec

import (
	"slices"
)

func Clone[S ~[]T, T any](values S) S {
	return slices.Clone(values)
}

func Clear[S ~*[]T, T any](self S) {
	clear(*self)
	*self = (*self)[:0]
}

func Put[S ~*[]T, T any](self S, it T) {
	*self = slices.Insert(*self, len(*self), it)
}

func Pop[S ~*[]T, T any](self S) T {
	l := len(*self)
	it := (*self)[l-1]
	*self = slices.Delete(*self, l-1, l)
	return it
}

func Ins[S ~*[]T, T any](self S, idx int, it T) {
	*self = slices.Insert(*self, idx, it)
}

func Del[S ~*[]T, T any](self S, idx int) T {
	it := (*self)[idx]
	*self = slices.Delete(*self, idx, idx+1)
	return it
}

func IsEmpty[S ~[]T, T any](self S) bool {
	return len(self) == 0
}

func Map[S ~[]T, T any](self S, fn func(i int, it T) T) S {
	vec := make(S, 0, len(self))

	for i, v := range self {
		vec[i] = fn(i, v)
	}

	return vec
}

func Filter[S ~[]T, T any](self S, fn func(i int, it T) bool) S {
	vec := make(S, 0, len(self))

	for i, v := range self {
		if fn(i, v) {
			vec = append(vec, v)
		}
	}

	return vec
}

func IsAny[S ~[]T, T any](self S, fn func(i int, it T) bool) bool {
	for i, v := range self {
		if fn(i, v) {
			return true
		}
	}

	return false
}

func IsAll[S ~[]T, T any](self S, fn func(i int, it T) bool) bool {
	for i, v := range self {
		if !fn(i, v) {
			return false
		}
	}

	return true
}

func Find[S ~[]T, T any](self S, fn func(it T) bool) bool {
	return slices.IndexFunc(self, fn) >= 0
}

func Index[S ~[]T, T any](self S, fn func(it T) bool) int {
	return slices.IndexFunc(self, fn)
}

func Sort[S ~[]T, T any](self S, fn func(a, b T) int) {
	slices.SortFunc(self, fn)
}

func IsSorted[S ~[]T, T any](self S, fn func(a, b T) int) {
	slices.IsSortedFunc(self, fn)
}

func Min[S ~[]T, T any](self S, fn func(a, b T) int) {
	slices.MinFunc(self, fn)
}

func Max[S ~[]T, T any](self S, fn func(a, b T) int) {
	slices.MaxFunc(self, fn)
}

func Eq[S1 ~[]T1, S2 ~[]T2, T1, T2 any](a S1, b S2, fn func(a T1, b T2) bool) bool {
	return slices.EqualFunc(a, b, fn)
}
