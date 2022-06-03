package list

import (
	"fmt"
	root "github.com/TimurTimergalin/collections"
)

type ArrayList[T any] struct {
	slice []T
}

func NewArrayList[T any](els ...T) *ArrayList[T] {
	return &ArrayList[T]{els}
}

func (a *ArrayList[T]) String() string {
	return fmt.Sprint(a.slice)
}

func (a *ArrayList[T]) negativeIndex(i int) int {
	if i < 0 {
		return a.Length() + i
	}
	return i
}

func (a *ArrayList[T]) assertIndex(i int) {
	if i < 0 || i >= a.Length() {
		panic("list index out of range")
	}
}

func (a *ArrayList[T]) Iterate(f func(T)) {
	for _, val := range a.slice {
		f(val)
	}
}

func (a *ArrayList[T]) Length() int {
	return len(a.slice)
}

func (a *ArrayList[T]) ToSlice() []T {
	return append(make([]T, 0), a.slice...)
}

func (a *ArrayList[T]) IsEmpty() bool {
	return a.Length() == 0
}

func (a *ArrayList[T]) Insert(i int, t T) {
	i = a.negativeIndex(i)
	if i != a.Length() {
		a.assertIndex(i)
	}
	a.slice = append(a.slice, *new(T))

	for j := a.Length() - 2; j >= i; j-- {
		a.slice[j+1] = a.slice[j]
	}
	a.slice[i] = t
	return
}

func (a *ArrayList[T]) Extract(i int) (res T) {
	i = a.negativeIndex(i)
	res = a.Get(i)
	a.slice = append(a.slice[:i], a.slice[i+1:]...)
	return
}

func (a *ArrayList[T]) Get(i int) (res T) {
	i = a.negativeIndex(i)
	a.assertIndex(i)
	res = a.slice[i]
	return
}

func (a *ArrayList[T]) Put(i int, t T) {
	i = a.negativeIndex(i)
	a.assertIndex(i)
	a.slice[i] = t
	return
}

func (a *ArrayList[T]) Clear() {
	a.slice = make([]T, 0)
}

func (a *ArrayList[T]) FromIterable(r root.Iterable[T]) {
	r.Iterate(func(el T) {
		a.slice = append(a.slice, el)
	})
}

func (a *ArrayList[T]) Add(els ...T) {
	a.slice = append(a.slice, els...)
}

func (a *ArrayList[T]) Copy() (res List[T]) {
	res = NewArrayList[T]()
	res.FromIterable(a)
	return
}
