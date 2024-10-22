package tu

import (
	"reflect"
	"slices"
	"strings"
)

type Vec[T any] []T

// NewVec return Vec[T]. There are three calling forms:
//
//	NewVec[T]() returns a Vec[T] of length 0 and capacity 8.
//	NewVec[T](cap) returns a Vec[T] of length 0 and capacity parameter cap.
//	NewVec[T](array[T] | slice[T] | Vec[T]) returns a shallow copy Vec[T].
func NewVec[T any](args ...any) Vec[T] {
	vec := make([]T, 0, 8)

	if len(args) == 0 {
		return vec
	} else if len(args) != 1 {
		panic("Does not conform to the three calling forms of " +
			"NewVec[T](), NewVec[T](integer), NewVec[T](array[T] | slice[T] | Vec[T]).")
	}

	value := reflect.ValueOf(args[0])

	switch value.Kind() {
	case reflect.Slice:
		if strings.HasPrefix(value.Type().Name(), "Vec") {
			vec = append(args[0].(Vec[T]))
		} else {
			vec = append(args[0].([]T))
		}

	case reflect.Array:
		vec = make([]T, 0, value.Len())
		for i := 0; i < value.Len(); i++ {
			vec = append(vec, value.Index(i).Interface().(T))
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vec = make(Vec[T], 0, value.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vec = make(Vec[T], 0, value.Uint())

	default:
		panic("Does not conform to the three calling forms of " +
			"NewVec[T](), NewVec[T](integer), NewVec[T](array[T] | slice[T] | Vec[T]).")
	}

	return vec
}

func (self *Vec[T]) Len() int {
	return len(*self)
}

func (self *Vec[T]) Cap() int {
	return cap(*self)
}

func (self *Vec[T]) IsEmpty() bool {
	return self.Len() == 0
}

func (self *Vec[T]) Put(it T) {
	*self = slices.Insert(*self, self.Len(), it)
}

func (self *Vec[T]) Pop() T {
	it := (*self)[self.Len()-1]
	*self = (*self)[:self.Len()-1]

	return it
}

func (self *Vec[T]) Ins(idx int, it T) {
	*self = slices.Insert(*self, idx, it)
}

func (self *Vec[T]) Del(idx int) T {
	it := (*self)[idx]
	*self = slices.Delete(*self, idx, idx+1)

	return it
}

func (self *Vec[T]) IsExist(fn func(it T) bool) bool {
	return slices.IndexFunc(*self, fn) > 0
}

func (self *Vec[T]) Index(fn func(it T) bool) int {
	return slices.IndexFunc(*self, fn)
}

func (self *Vec[T]) Sort(fn func(a, b T) int) {
	slices.SortFunc(*self, fn)
}

func (self *Vec[T]) IsSort(fn func(a, b T) int) bool {
	return slices.IsSortedFunc(*self, fn)
}

func (self *Vec[T]) Map(fn func(i int, it T) T) Vec[T] {
	vec := NewVec[T](self.Len())

	for i, v := range *self {
		vec.Put(fn(i, v))
	}

	return vec
}

func (self *Vec[T]) Filter(fn func(i int, it T) bool) Vec[T] {
	vec := NewVec[T](self.Len())

	for i, v := range *self {
		if fn(i, v) {
			vec.Put(v)
		}
	}

	return vec
}

func (self *Vec[T]) IsAny(fn func(i int, it T) bool) bool {
	for i, v := range *self {
		if fn(i, v) {
			return true
		}
	}

	return false
}

func (self *Vec[T]) IsAll(fn func(i int, it T) bool) bool {
	for i, v := range *self {
		if !fn(i, v) {
			return false
		}
	}

	return true
}
