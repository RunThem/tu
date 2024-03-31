package u

import (
	"fmt"
	"iter"
	"maps"
)

type Map[K comparable, T any] struct {
	maps map[K]T
}

func NewMap[K comparable, T any](other map[K]T) Map[K, T] {
	m := Map[K, T]{}
	if other != nil {
		m.maps = maps.Clone(other)
	} else {
		m.maps = make(map[K]T)
	}

	return m
}

func (m *Map[K, T]) Len() int {
	return len(m.maps)
}

func (m *Map[K, T]) IsEmpty() bool {
	return len(m.maps) == 0
}

func (m *Map[K, T]) Clear() {
}

func (m *Map[K, T]) At(key K) T {
	return m.maps[key]
}

func (m *Map[K, T]) Re(key K, val T) {
	m.maps[key] = val
}

func (m *Map[K, T]) Pop(key K) T {
	v, ok := m.maps[key]
	if ok {
		delete(m.maps, key)
	}

	return v
}

func (m *Map[K, T]) Put(key K, val T) {
	m.maps[key] = val
}

func (m *Map[K, T]) String() string {
	return fmt.Sprintf("+%v", m.maps)
}

func (m *Map[K, T]) Range() iter.Seq2[K, T] {
	return func(yield func(K, T) bool) {
		for k, v := range m.maps {
			if !yield(k, v) {
				return
			}
		}

		return
	}

}
