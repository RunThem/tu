package tu

import (
	"fmt"
	"iter"
	"slices"
)

type Vec[T any] struct {
	items []T
	cmp   func(T, T) int
}

func NewVec[T any](cmp func(T, T) int, other ...T) *Vec[T] {
	return &Vec[T]{items: slices.Clone(other), cmp: cmp}
}

func (mod *Vec[T]) Len() int {
	return len(mod.items)
}

func (mod *Vec[T]) IsEmpty() bool {
	return len(mod.items) == 0
}

func (mod *Vec[T]) Clear() {
	mod.items = []T{}
}

func (mod *Vec[T]) At(idx int) T {
	if idx < 0 || idx >= mod.Len() {
		panic("out off range")
	}

	return mod.items[idx]
}

func (mod *Vec[T]) AtFront() T {
	return mod.At(0)
}

func (mod *Vec[T]) AtBack() T {
	return mod.At(mod.Len() - 1)
}

func (mod *Vec[T]) Re(idx int, it T) {
	if idx < 0 || idx >= mod.Len() {
		panic("out off range")
	}

	mod.items[idx] = it
}

func (mod *Vec[T]) ReFront(it T) {
	if mod.Len() == 0 {
		panic("out off range")
	}

	mod.items[0] = it
}

func (mod *Vec[T]) ReBack(it T) {
	if mod.Len() == 0 {
		panic("out off range")
	}

	mod.items[len(mod.items)-1] = it
}

func (mod *Vec[T]) Pop(idx int) T {
	if idx < 0 || idx >= mod.Len() {
		panic("out off range")
	}

	item := mod.items[idx]
	if idx != mod.Len()-1 {
		copy(mod.items[idx:], mod.items[idx+1:])
	}

	mod.items = mod.items[:len(mod.items)-1]

	return item
}

func (mod *Vec[T]) PopFront() T {
	return mod.Pop(0)
}

func (mod *Vec[T]) PopBack() T {
	return mod.Pop(mod.Len() - 1)
}

func (mod *Vec[T]) Put(idx int, it T) {
	if idx < 0 || idx > mod.Len() {
		panic("out off range")
	}

	mod.items = append(mod.items, it)
	if idx != mod.Len() {
		copy(mod.items[idx+1:], mod.items[idx:])
		mod.items[idx] = it
	}
}

func (mod *Vec[T]) PutFront(it T) {
	mod.Put(0, it)
}

func (mod *Vec[T]) PutBack(it T) {
	mod.Put(mod.Len(), it)
}

func (mod *Vec[T]) String() string {
	return fmt.Sprintf("%v", mod.items)
}

func (mod *Vec[T]) Range(order bool) iter.Seq2[int, T] {
	var fn iter.Seq2[int, T]
	if order {
		fn = func(yield func(int, T) bool) {
			for i := 0; i < mod.Len(); i++ {
				if !yield(i, mod.items[i]) {
					return
				}
			}

			return
		}
	} else {
		fn = func(yield func(int, T) bool) {
			for i := mod.Len() - 1; i >= 0; i-- {
				if !yield(i, mod.items[i]) {
					return
				}
			}

			return
		}
	}

	return fn
}

func (mod *Vec[T]) Map(fn func(idx int, it T) T) *Vec[T] {
	vec := NewVec[T](mod.cmp)

	for i, it := range mod.Range(true) {
		vec.PutBack(fn(i, it))
	}

	return vec
}

func (mod *Vec[T]) Filter(fn func(idx int, it T) bool) *Vec[T] {
	vec := NewVec[T](mod.cmp)

	for i, it := range mod.Range(true) {
		if fn(i, it) {
			vec.PutBack(it)
		}
	}

	return vec
}

func (mod *Vec[T]) IsAny(fn func(idx int, it T) bool) bool {
	for i, it := range mod.Range(true) {
		if fn(i, it) {
			return true
		}
	}

	return false
}

func (mod *Vec[T]) IsAll(fn func(idx int, it T) bool) bool {
	for i, it := range mod.Range(true) {
		if !fn(i, it) {
			return false
		}
	}

	return true
}

func (mod *Vec[T]) Find(it T) (int, T) {
	if mod.cmp != nil {
		for i, _it := range mod.Range(true) {
			if mod.cmp(it, _it) == 0 {
				return i, it
			}
		}
	}

	var t T
	return -1, t
}

func (mod *Vec[T]) FindBy(fn func(idx int, it T) bool) (int, T) {
	for i, it := range mod.Range(true) {
		if fn(i, it) {
			return i, it
		}
	}

	var t T
	return -1, t
}

func (mod *Vec[T]) Index(it T) int {
	if mod.cmp != nil {
		for i, _it := range mod.Range(true) {
			if mod.cmp(it, _it) == 0 {
				return i
			}
		}
	}

	return -1
}

func (mod *Vec[T]) IndexBy(fn func(it T) bool) int {
	for i, it := range mod.Range(true) {
		if fn(it) {
			return i
		}
	}

	return -1
}

func (mod *Vec[T]) Sort() {
	slices.SortFunc(mod.items, mod.cmp)
}

func (mod *Vec[T]) SortBy(fn func(a, b T) int) {
	slices.SortFunc(mod.items, fn)
}

func (mod *Vec[T]) IsSorted() bool {
	return slices.IsSortedFunc(mod.items, mod.cmp)
}

func (mod *Vec[T]) IsSortedBy(fn func(a, b T) int) bool {
	return slices.IsSortedFunc(mod.items, fn)
}

func (mod *Vec[T]) Min() T {
	return slices.MinFunc(mod.items, mod.cmp)
}

func (mod *Vec[T]) MinBy(fn func(a, b T) int) T {
	return slices.MinFunc(mod.items, fn)
}

func (mod *Vec[T]) Max() T {
	return slices.MaxFunc(mod.items, mod.cmp)
}

func (mod *Vec[T]) MaxBy(fn func(a, b T) int) T {
	return slices.MaxFunc(mod.items, fn)
}

func (mod *Vec[T]) Items() []T {
	return mod.items
}

func (mod *Vec[T]) Copy() *Vec[T] {
	vec := NewVec[T](mod.cmp)

	for _, t := range mod.Range(true) {
		vec.PutBack(t)
	}

	return vec
}
