package u

import (
	"fmt"
	"iter"
	"slices"
)

type Vec[T any] struct {
	items []T
	cmp   func(T, T) int
}

func NewVec[T any](cmp func(T, T) int, other ...T) Vec[T] {
	return Vec[T]{items: slices.Clone(other), cmp: cmp}
}

func (v *Vec[T]) Len() int {
	return len(v.items)
}

func (v *Vec[T]) IsEmpty() bool {
	return len(v.items) == 0
}

func (v *Vec[T]) Clear() {
	v.items = []T{}
}

func (v *Vec[T]) At(idx int) T {
	if idx < 0 || idx >= v.Len() {
		panic("out off range")
	}

	return v.items[idx]
}

func (v *Vec[T]) AtFront() T {
	return v.At(0)
}

func (v *Vec[T]) AtBack() T {
	return v.At(v.Len() - 1)
}

func (v *Vec[T]) Re(idx int, it T) {
	if idx < 0 || idx >= v.Len() {
		panic("out off range")
	}

	v.items[idx] = it
}

func (v *Vec[T]) ReFront(it T) {
	if v.Len() == 0 {
		panic("out off range")
	}

	v.items[0] = it
}

func (v *Vec[T]) ReBack(it T) {
	if v.Len() == 0 {
		panic("out off range")
	}

	v.items[len(v.items)-1] = it
}

func (v *Vec[T]) Pop(idx int) T {
	if idx < 0 || idx >= v.Len() {
		panic("out off range")
	}

	item := v.items[idx]
	if idx != v.Len()-1 {
		copy(v.items[idx:], v.items[idx+1:])
	}

	v.items = v.items[:len(v.items)-1]

	return item
}

func (v *Vec[T]) PopFront() T {
	return v.Pop(0)
}

func (v *Vec[T]) PopBack() T {
	return v.Pop(v.Len() - 1)
}

func (v *Vec[T]) Put(idx int, it T) {
	if idx < 0 || idx > v.Len() {
		panic("out off range")
	}

	v.items = append(v.items, it)
	if idx != v.Len() {
		copy(v.items[idx+1:], v.items[idx:])
		v.items[idx] = it
	}
}

func (v *Vec[T]) PutFront(it T) {
	v.Put(0, it)
}

func (v *Vec[T]) PutBack(it T) {
	v.Put(v.Len(), it)
}

func (v *Vec[T]) String() string {
	return fmt.Sprintf("+%v", v.items)
}

func (v *Vec[T]) Range(order bool) iter.Seq2[int, T] {
	var fn iter.Seq2[int, T]
	if order {
		fn = func(yield func(int, T) bool) {
			for i := 0; i < v.Len(); i++ {
				if !yield(i, v.items[i]) {
					return
				}
			}

			return
		}
	} else {
		fn = func(yield func(int, T) bool) {
			for i := v.Len() - 1; i >= 0; i-- {
				if !yield(i, v.items[i]) {
					return
				}
			}

			return
		}
	}

	return fn
}

func (v *Vec[T]) Map(fn func(idx int, it T) T) Vec[T] {
	vec := NewVec[T](v.cmp)

	for i, it := range v.Range(true) {
		vec.PutBack(fn(i, it))
	}

	return vec
}

func (v *Vec[T]) Filter(fn func(idx int, it T) bool) Vec[T] {
	vec := NewVec[T](v.cmp)

	for i, it := range v.Range(true) {
		if fn(i, it) {
			vec.PutBack(it)
		}
	}

	return vec
}

func (v *Vec[T]) IsAny(fn func(idx int, it T) bool) bool {
	for i, it := range v.Range(true) {
		if fn(i, it) {
			return true
		}
	}

	return false
}

func (v *Vec[T]) IsAll(fn func(idx int, it T) bool) bool {
	for i, it := range v.Range(true) {
		if !fn(i, it) {
			return false
		}
	}

	return true
}

func (v *Vec[T]) Find(it T) (int, T) {
	if v.cmp != nil {
		for i, _it := range v.Range(true) {
			if v.cmp(it, _it) == 0 {
				return i, it
			}
		}
	}

	var t T
	return -1, t
}

func (v *Vec[T]) FindBy(fn func(idx int, it T) bool) (int, T) {
	for i, it := range v.Range(true) {
		if fn(i, it) {
			return i, it
		}
	}

	var t T
	return -1, t
}

func (v *Vec[T]) Index(it T) int {
	if v.cmp != nil {
		for i, _it := range v.Range(true) {
			if v.cmp(it, _it) == 0 {
				return i
			}
		}
	}

	return -1
}

func (v *Vec[T]) IndexBy(fn func(it T) bool) int {
	for i, it := range v.Range(true) {
		if fn(it) {
			return i
		}
	}

	return -1
}

func (v *Vec[T]) Sort() {
	slices.SortFunc(v.items, v.cmp)
}

func (v *Vec[T]) SortBy(fn func(a, b T) int) {
	slices.SortFunc(v.items, fn)
}

func (v *Vec[T]) IsSorted() bool {
	return slices.IsSortedFunc(v.items, v.cmp)
}

func (v *Vec[T]) IsSortedBy(fn func(a, b T) int) bool {
	return slices.IsSortedFunc(v.items, fn)
}

func (v *Vec[T]) Min() T {
	return slices.MinFunc(v.items, v.cmp)
}

func (v *Vec[T]) MinBy(fn func(a, b T) int) T {
	return slices.MinFunc(v.items, fn)
}

func (v *Vec[T]) Max() T {
	return slices.MaxFunc(v.items, v.cmp)
}

func (v *Vec[T]) MaxBy(fn func(a, b T) int) T {
	return slices.MaxFunc(v.items, fn)
}

func (v *Vec[T]) Items() []T {
	return v.items
}
