package tbl

import (
	"maps"
)

func Clone[S ~map[K]V, K comparable, V any](values S) S {
	return maps.Clone(values)
}

func Clear[S ~map[K]V, K comparable, V any](self S) {
	clear(self)
}

func Pop[S ~map[K]V, K comparable, V any](self S, key K) V {
	val := self[key]
	delete(self, key)
	return val
}

func Put[S ~map[K]V, K comparable, V any](self S, key K, val V) {
	self[key] = val
}

func IsEmpty[S ~map[K]V, K comparable, V any](self S) bool {
	return len(self) == 0
}

func Map[S ~map[K]V, K comparable, V any](self S, fn func(k K, v V) V) S {
	tbl := make(S, len(self))

	for k, v := range self {
		tbl[k] = fn(k, v)
	}

	return tbl
}

func Filter[S ~map[K]V, K comparable, V any](self S, fn func(k K, v V) bool) S {
	tbl := make(S, len(self))

	for k, v := range self {
		if fn(k, v) {
			tbl[k] = v
		}
	}

	return tbl
}

func IsAny[S ~map[K]V, K comparable, V any](self S, fn func(k K, v V) bool) bool {
	for k, v := range self {
		if fn(k, v) {
			return true
		}
	}

	return false
}

func IsAll[S ~map[K]V, K comparable, V any](self S, fn func(k K, v V) bool) bool {
	for k, v := range self {
		if !fn(k, v) {
			return false
		}
	}

	return true
}

func Keys[S ~map[K]V, K comparable, V any](self S) []K {
	vec := make([]K, 0, len(self))
	for k, _ := range self {
		vec = append(vec, k)
	}

	return vec
}

func Vals[S ~map[K]V, K comparable, V any](self S) []V {
	vec := make([]V, 0, len(self))
	for _, v := range self {
		vec = append(vec, v)
	}

	return vec
}
