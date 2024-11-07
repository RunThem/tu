package tu

import (
	"reflect"
)

// Tbl is a generic map-like structure for keys of any comparable type K and values of any type V.
type Tbl[K comparable, V any] map[K]V

// NewTbl returns Tbl[K, V], There are two calling forms:
//
//	NewTbl[K, V]() returns a Tbl[K, V] of length 0.
//	NewTbl[K, V](map[K, V] | Tbl[K, V]) returns a shallow copy of the given map or Tbl.
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

// Len returns the number of key-value pairs in the table.
func (self Tbl[K, V]) Len() int {
	return len(self)
}

// IsEmpty returns true if the table has no key-value pairs.
func (self Tbl[K, V]) IsEmpty() bool {
	return self.Len() == 0
}

// Put sets the value for a specified key in the table.
func (self Tbl[K, V]) Put(key K, val V) {
	self[key] = val
}

// Pop removes a key-value pair from the table and returns the removed value.
func (self Tbl[K, V]) Pop(key K) V {
	val, ok := self[key]
	if ok {
		delete(self, key)
	}

	return val
}

// Keys returns a Vec containing all the keys from the table.
func (self Tbl[K, V]) Keys() Vec[K] {
	vec := NewVec[K](self.Len())

	for k := range self {
		vec.Put(k)
	}

	return vec
}

// Vals returns a Vec containing all the values from the table.
func (self Tbl[K, V]) vals() Vec[V] {
	vec := NewVec[V](self.Len())

	for _, v := range self {
		vec.Put(v)
	}

	return vec
}

// Map applies a function to each key-value pair and returns a new Tbl with the results.
func (self Tbl[K, V]) Map(fn func(key K, val V) V) Tbl[K, V] {
	tbl := NewTbl[K, V]()

	for k, v := range self {
		tbl[k] = fn(k, v)
	}

	return tbl
}

// Filter creates a new Tbl with only the key-value pairs that satisfy the provided function.
func (self Tbl[K, V]) Filter(fn func(key K, val V) bool) Tbl[K, V] {
	tbl := NewTbl[K, V]()

	for k, v := range self {
		if fn(k, v) {
			tbl[k] = v
		}
	}

	return tbl
}

// FilterMap applies a function to each key-value pair, filtering and transforming the pairs in the
// process.
func (self Tbl[K, V]) FilterMap(fn func(key K, val V) (bool, V)) Tbl[K, V] {
	tbl := NewTbl[K, V]()

	for k, v := range self {
		if ok, value := fn(k, v); ok {
			tbl[k] = value
		}
	}

	return tbl
}

// IsAny returns true if any key-value pair satisfies the provided function.
func (self Tbl[K, V]) IsAny(fn func(key K, val V) bool) bool {
	for k, v := range self {
		if fn(k, v) {
			return true
		}
	}

	return false
}

// IsAll returns true if all key-value pairs satisfy the provided function.
func (self Tbl[K, V]) IsAll(fn func(key K, val V) bool) bool {
	for k, v := range self {
		if !fn(k, v) {
			return false
		}
	}

	return true
}
