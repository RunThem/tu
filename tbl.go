package tu

import (
	"reflect"
)

type Tbl[K comparable, V any] map[K]V

// NewTbl return Tbl[K, V], There are two calling forms:
//
//	NewTbl[K, V]() return a Tbl[K, V] of length 0.
//	NewTbl[K, V](map[K, V] | Tbl[K, V]) return a shallow copy Tbl[K, V].
func NewTbl[K comparable, V any](args ...any) Tbl[K, V] {
	tbl := make(map[K]V)

	if len(args) == 0 {
		return tbl
	} else if len(args) != 1 {
		panic("Does not conform to the two calling forms of " +
			"NewMap[K, V](), NewMap[K, V](map[K, V] | Tbl[K, V])")
	}

	value := reflect.ValueOf(args[0])

	switch value.Kind() {
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			tbl[iter.Key().Interface().(K)] = iter.Value().Interface().(V)
		}
	default:
		panic("Does not conform to the two calling forms of " +
			"NewMap[K, V](), NewMap[K, V](map[K, V] | Tbl[K, V])")
	}

	return tbl
}

func (self Tbl[K, V]) Len() int {
	return len(self)
}

func (self Tbl[K, V]) IsEmpty() bool {
	return self.Len() == 0
}

func (self Tbl[K, V]) Put(key K, val V) {
	self[key] = val
}

func (self Tbl[K, V]) Pop(key K) V {
	val, ok := self[key]
	if ok {
		delete(self, key)
	}

	return val
}

func (self Tbl[K, V]) Keys() Vec[K] {
	vec := NewVec[K](self.Len())

	for k, _ := range self {
		vec.Put(k)
	}

	return vec
}

func (self Tbl[K, V]) vals() Vec[V] {
	vec := NewVec[V](self.Len())

	for _, v := range self {
		vec.Put(v)
	}

	return vec
}

func (self Tbl[K, V]) Map(fn func(key K, val V) V) Tbl[K, V] {
	tbl := make(map[K]V)

	for k, v := range self {
		tbl[k] = fn(k, v)
	}

	return tbl
}

func (self Tbl[K, V]) Filter(fn func(key K, val V) bool) Tbl[K, V] {
	tbl := make(map[K]V)

	for k, v := range self {
		if fn(k, v) {
			tbl[k] = v
		}
	}

	return tbl
}

func (self Tbl[K, V]) IsAny(fn func(key K, val V) bool) bool {
	for k, v := range self {
		if fn(k, v) {
			return true
		}
	}

	return false
}

func (self Tbl[K, V]) IsAll(fn func(key K, val V) bool) bool {
	for k, v := range self {
		if !fn(k, v) {
			return false
		}
	}

	return true
}
