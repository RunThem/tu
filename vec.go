package u

import (
	"fmt"
	"iter"
	"slices"
)

// Vec is a linear data structure, the internal is a slice.
type Vec[T any] struct {
	items []T
}

// NewVec creates a new vector.
func NewVec[T any](items ...T) *Vec[T] {
	v := &Vec[T]{items: make([]T, 0)}

	v.items = append(v.items, items...)

	return v
}

// Size returns the length of the vector.
func (v *Vec[T]) Size() int {
	return len(v.items)
}

// Empty returns true if the vector is empty, otherwise returns false.
func (v *Vec[T]) Empty() bool {
	return len(v.items) == 0
}

// Clear removes all items from the vector.
func (v *Vec[T]) Clear() {
	v.items = []T{}
}

// At returns the value at position pos, returns nil if pos is out off range.
func (v *Vec[T]) At(idx int) T {
	if idx < 0 || idx >= v.Size() {
		panic("out off range")
	}

	return v.items[idx]
}

// AtFront returns the first value in the vector, returns nil if the vector is empty.
func (v *Vec[T]) AtFront() T {
	return v.At(0)
}

// AtBack returns the last value in the vector, returns nil if the vector is empty.
func (v *Vec[T]) AtBack() T {
	return v.At(v.Size() - 1)
}

// Re sets the value it to the vector to position idx.
func (v *Vec[T]) Re(idx int, it T) {
	if idx < 0 || idx >= v.Size() {
		panic("out off range")
	}

	v.items[idx] = it
}

// ReFront sets the value it to the vector at position first.
func (v *Vec[T]) ReFront(it T) {
	if v.Size() == 0 {
		panic("out off range")
	}

	v.items[0] = it
}

// ReBack sets the value it to the vector at position back.
func (v *Vec[T]) ReBack(it T) {
	if v.Size() == 0 {
		panic("out off range")
	}

	v.items[len(v.items)-1] = it
}

// Pop returns the position idx value of the vector and erase ti, returns nil if the vector is empty.
func (v *Vec[T]) Pop(idx int) T {
	if idx < 0 || idx >= v.Size() {
		panic("out off range")
	}

	item := v.items[idx]
	if idx != v.Size()-1 {
		copy(v.items[idx:], v.items[idx+1:])
	}

	v.items = v.items[:len(v.items)-1]

	return item
}

// PopFront returns the first value of the vector and erase ti, returns nil if the vector is empty.
func (v *Vec[T]) PopFront() T {
	return v.Pop(0)
}

// PopBack returns the last value of the vector and erase ti, returns nil if the vector is empty.
func (v *Vec[T]) PopBack() T {
	return v.Pop(v.Size() - 1)
}

// Put inserts the value it to the vector at position pos.
func (v *Vec[T]) Put(idx int, it T) {
	if idx < 0 || idx > v.Size() {
		panic("out off range")
	}

	v.items = append(v.items, it)
	if idx != v.Size() {
		copy(v.items[idx+1:], v.items[idx:])
		v.items[idx] = it
	}
}

// PutFront inserts the value it to the vector at position pos.
func (v *Vec[T]) PutFront(it T) {
	v.Put(0, it)
}

// PutBack inserts the value it to the vector at position pos.
func (v *Vec[T]) PutBack(it T) {
	v.Put(v.Size(), it)
}

// String returns a string representation of the vector.
func (v *Vec[T]) String() string {
	return fmt.Sprintf("+%v", v.items)
}

// Range traversal method based on `rangefunc` feature.
func (v *Vec[T]) Range(order bool) iter.Seq2[int, T] {
	var fn iter.Seq2[int, T]
	if order {
		fn = func(yield func(int, T) bool) {
			for i := 0; i < v.Size(); i++ {
				if !yield(i, v.items[i]) {
					return
				}
			}

			return
		}
	} else {
		fn = func(yield func(int, T) bool) {
			for i := v.Size() - 1; i >= 0; i-- {
				if !yield(i, v.items[i]) {
					return
				}
			}

			return
		}
	}

	return fn
}

// Map invokes the given function once for each element and returns a container containing the values returned by the
// given function.
func (v *Vec[T]) Map(fn func(idx int, it T) T) *Vec[T] {
	vec := NewVec[T]()

	for i, it := range v.Range(true) {
		vec.PutBack(fn(i, it))
	}

	return vec
}

// Filter returns a new container containing all elements for which the given function returns a true value.
func (v *Vec[T]) Filter(fn func(idx int, it T) bool) *Vec[T] {
	vec := NewVec[T]()

	for i, it := range v.Range(true) {
		if fn(i, it) {
			vec.PutBack(it)
		}
	}

	return vec
}

// Any passes each element of the collection to the given function and returns true if the function ever returns true
// for any element.
func (v *Vec[T]) Any(fn func(idx int, it T) bool) bool {
	for i, it := range v.Range(true) {
		if fn(i, it) {
			return true
		}
	}

	return false
}

// All passes each element of the collection to the given function and returns true if the function returns true for
// all elements.
func (v *Vec[T]) All(fn func(idx int, it T) bool) bool {
	for i, it := range v.Range(true) {
		if !fn(i, it) {
			return false
		}
	}

	return true
}

// Find passes each element of the container to the given function and returns the first (index,value) for which the
// function is true or -1,nil otherwise if no element matches the criteria.
func (v *Vec[T]) Find(fn func(idx int, it T) bool) (int, T) {
	for i, it := range v.Range(true) {
		if fn(i, it) {
			return i, it
		}
	}

	var t T
	return -1, t
}

// Index returns the first index i satisfying f(s[i]),
// or -1 if none do.
func (v *Vec[T]) Index(fn func(it T) bool) int {
	for i, it := range v.Range(true) {
		if fn(it) {
			return i
		}
	}

	return -1
}

// Sort sorts the slice x in ascending order as determined by the cmp
// function. This sort is not guaranteed to be stable.
// cmp(a, b) should return a negative number when a < b, a positive number when
// a > b and zero when a == b.
//
// SortFunc requires that cmp is a strict weak ordering.
// See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.
func (v *Vec[T]) Sort(fn func(a, b T) int) {
	slices.SortFunc(v.items, fn)
}

// IsSorted reports whether x is sorted in ascending order, with cmp as the
// comparison function as defined by [Sort].
func (v *Vec[T]) IsSorted(fn func(a, b T) int) bool {
	return slices.IsSortedFunc(v.items, fn)
}

// Min returns the minimal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one minimal element
// according to the cmp function, Min returns the first one.
func (v *Vec[T]) Min(fn func(a, b T) int) T {
	return slices.MinFunc(v.items, fn)
}

// Max returns the maximal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one maximal element
// according to the cmp function, Max returns the first one.
func (v *Vec[T]) Max(fn func(a, b T) int) T {
	return slices.MaxFunc(v.items, fn)
}

// Items returns original data.
func (v *Vec[T]) Items() []T {
	return v.items
}
