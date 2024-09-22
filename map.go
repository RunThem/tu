package tu

import (
	"fmt"
	"maps"
)

type Map[K comparable, V any] struct {
	items map[K]V
}

func NewMap[K comparable, V any](other map[K]V) *Map[K, V] {
	m := &Map[K, V]{}
	if other != nil {
		m.items = maps.Clone(other)
	} else {
		m.items = make(map[K]V)
	}

	return m
}

func (mod *Map[K, V]) Len() int {
	return len(mod.items)
}

func (mod *Map[K, V]) IsEmpty() bool {
	return len(mod.items) == 0
}

func (mod *Map[K, V]) Clear() {
}

func (mod *Map[K, V]) IsExist(key K) bool {
	_, ok := mod.items[key]
	return ok
}

func (mod *Map[K, V]) At(key K) V {
	return mod.items[key]
}

func (mod *Map[K, V]) Re(key K, val V) {
	mod.items[key] = val
}

func (mod *Map[K, V]) Pop(key K) V {
	v, ok := mod.items[key]
	if ok {
		delete(mod.items, key)
	}

	return v
}

func (mod *Map[K, V]) Put(key K, val V) {
	mod.items[key] = val
}

func (mod *Map[K, V]) String() string {
	return fmt.Sprintf("+%v", mod.items)
}

func (mod *Map[K, V]) L(yield func(K, V) bool) {
	for k, v := range mod.items {
		if !yield(k, v) {
			return
		}
	}
}

func (mod *Map[K, V]) Map(fn func(key K, val V) V) *Map[K, V] {
	m := NewMap[K, V](nil)

	for k, v := range mod.L {
		m.Put(k, fn(k, v))
	}

	return m
}

func (mod *Map[K, V]) Filter(fn func(key K, val V) bool) *Map[K, V] {
	m := NewMap[K, V](nil)

	for k, v := range mod.L {
		if fn(k, v) {
			m.Put(k, v)
		}
	}

	return m
}

func (mod *Map[K, V]) IsAny(fn func(key K, val V) bool) bool {
	for k, v := range mod.L {
		if fn(k, v) {
			return true
		}
	}

	return false
}

func (mod *Map[K, V]) IsAll(fn func(key K, val V) bool) bool {
	for k, v := range mod.L {
		if !fn(k, v) {
			return false
		}
	}

	return true
}

func (mod *Map[K, V]) Keys() *Vec[K] {
	vec := NewVec[K]()
	for k, _ := range mod.L {
		vec.Put(-1, k)
	}

	return vec
}

func (mod *Map[K, V]) Vals() *Vec[V] {
	vec := NewVec[V]()
	for _, v := range mod.L {
		vec.Put(-1, v)
	}

	return vec
}

func (mod *Map[K, V]) Clone() *Map[K, V] {
	return &Map[K, V]{items: maps.Clone(mod.items)}
}
