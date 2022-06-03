// Package set provides Set struct and functions to work with it
package set

import (
	"fmt"
	root "github.com/TimurTimergalin/collections.git"
)

type Set[T comparable] map[T]bool

func NewSet[T comparable](els ...T) (res Set[T]) {
	res = make(Set[T])
	res.Add(els...)
	return
}

func (h Set[T]) Add(t ...T) {
	for _, val := range t {
		h[val] = true
	}
}

func (h Set[T]) Remove(t ...T) {
	for _, val := range t {
		delete(h, val)
	}
}

func (h Set[T]) Contains(t T) bool {
	_, ok := h[t]
	return ok
}

func (h Set[T]) Clear() {
	for key := range h {
		delete(h, key)
	}
}

func (h Set[T]) String() (res string) {
	res += "{"

	h.Iterate(func(el T) {
		res += fmt.Sprint(el) + " "
	})

	if len(res) > 1 {
		res = res[:len(res)-1]
	}
	res += "}"
	return
}

func (h Set[T]) Iterate(f func(T)) {
	for key := range h {
		f(key)
	}
}

func (h Set[T]) Length() int {
	return len(h)
}

func (h Set[T]) ToSlice() (res []T) {
	res = make([]T, len(h))

	h.Iterate(func(el T) {
		res = append(res, el)
	})
	return
}

func (h Set[T]) IsEmpty() bool {
	return len(h) == 0
}

func (h Set[T]) FromIterable(i root.Iterable[T]) {
	i.Iterate(func(el T) {
		h.Add(el)
	})
}

func (h Set[T]) Copy() (res Set[T]) {
	res = make(Set[T])
	for el := range h {
		res.Add(el)
	}
	return
}
