package tu

import (
	"fmt"
	"slices"
)

type Vec[T any] struct {
	items []T
}

func NewVec[T any]() *Vec[T] {
	return &Vec[T]{items: make([]T, 0)}
}

func NewVecFrom[T any](values []T) *Vec[T] {
	return &Vec[T]{items: slices.Clone(values)}
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
	if idx < -mod.Len() || idx >= mod.Len() {
		panic(fmt.Sprintf("Out of range { %d, %d }", mod.Len(), idx))
	}

	if idx < 0 {
		idx += mod.Len()
	}

	return mod.items[idx]
}

func (mod *Vec[T]) Re(idx int, it T) {
	if idx < -mod.Len() || idx >= mod.Len() {
		panic(fmt.Sprintf("Out of range { %d, %d }", mod.Len(), idx))
	}

	if idx < 0 {
		idx += mod.Len()
	}

	mod.items[idx] = it
}

func (mod *Vec[T]) Pop(idx int) T {
	if idx < -mod.Len() || idx >= mod.Len() {
		panic(fmt.Sprintf("Out of range { %d, %d }", mod.Len(), idx))
	}

	if idx < 0 {
		idx += mod.Len()
	}

	item := mod.items[idx]

	if idx != mod.Len()-1 {
		copy(mod.items[idx:], mod.items[idx+1:])
	}

	mod.items = mod.items[:mod.Len()-1]

	return item
}

func (mod *Vec[T]) Put(idx int, it T) {
	if idx < -(mod.Len()+1) || idx > mod.Len() {
		panic(fmt.Sprintf("Out of range { %d, %d }", mod.Len(), idx))
	}

	if idx < 0 {
		idx += mod.Len() + 1
	}

	mod.items = append(mod.items, it)

	if idx != mod.Len() {
		copy(mod.items[idx+1:], mod.items[idx:])
		mod.items[idx] = it
	}
}

func (mod *Vec[T]) String() string {
	return fmt.Sprintf("%v", mod.items)
}

func (mod *Vec[T]) L(yield func(int, T) bool) {
	for i := 0; i < mod.Len(); i++ {
		if !yield(i, mod.items[i]) {
			return
		}
	}
}

func (mod *Vec[T]) R(yield func(int, T) bool) {
	for i := mod.Len() - 1; i >= 0; i-- {
		if !yield(i, mod.items[i]) {
			return
		}
	}
}

func (mod *Vec[T]) Map(fn func(idx int, it T) T) *Vec[T] {
	vec := NewVec[T]()

	for i, it := range mod.L {
		vec.Put(-1, fn(i, it))
	}

	return vec
}

func (mod *Vec[T]) Filter(fn func(idx int, it T) bool) *Vec[T] {
	vec := NewVec[T]()

	for i, it := range mod.L {
		if fn(i, it) {
			vec.Put(-1, it)
		}
	}

	return vec
}

func (mod *Vec[T]) IsAny(fn func(idx int, it T) bool) bool {
	for i, it := range mod.L {
		if fn(i, it) {
			return true
		}
	}

	return false
}

func (mod *Vec[T]) IsAll(fn func(idx int, it T) bool) bool {
	for i, it := range mod.L {
		if !fn(i, it) {
			return false
		}
	}

	return true
}

func (mod *Vec[T]) Find(fn func(idx int, it T) bool) (int, T) {
	for i, it := range mod.L {
		if fn(i, it) {
			return i, it
		}
	}

	var t T
	return -1, t
}

func (mod *Vec[T]) Index(fn func(it T) bool) int {
	for i, it := range mod.L {
		if fn(it) {
			return i
		}
	}

	return -1
}

func (mod *Vec[T]) Sort(fn func(a, b T) int) {
	slices.SortFunc(mod.items, fn)
}

func (mod *Vec[T]) IsSorted(fn func(a, b T) int) bool {
	return slices.IsSortedFunc(mod.items, fn)
}

func (mod *Vec[T]) Min(fn func(a, b T) int) T {
	return slices.MinFunc(mod.items, fn)
}

func (mod *Vec[T]) Max(fn func(a, b T) int) T {
	return slices.MaxFunc(mod.items, fn)
}

func (mod *Vec[T]) Items() []T {
	return mod.items
}

func (mod *Vec[T]) Clone() *Vec[T] {
	return NewVecFrom(mod.items)
}
