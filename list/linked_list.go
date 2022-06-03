package list

import (
	"fmt"
	root "github.com/TimurTimergalin/collections.git"
)

type element[T any] struct {
	val        *T
	prev, next *element[T]
}

func (e *element[T]) Set(el T) {
	e.val = &el
}

func (e *element[T]) Get() T {
	return *e.val
}

func newElement[T any](el T) *element[T] {
	return &element[T]{&el, nil, nil}
}

type LinkedList[T any] struct {
	first, last *element[T]
}

func NewLinkedList[T any](els ...T) (res *LinkedList[T]) {
	res = &LinkedList[T]{nil, nil}
	res.Add(els...)
	return
}

func (l *LinkedList[T]) indexError() {
	panic("list index out of range")
}

func (l *LinkedList[T]) String() (res string) {
	res = "["
	l.iterateForward(func(el T) {
		res += fmt.Sprint(el) + " "
	})
	if len(res) > 1 {
		res = res[:len(res)-1]
	}
	res += "]"
	return
}

func (l *LinkedList[T]) iterateForward(f func(T)) {
	for el := l.first; el != nil; el = el.next {
		f(el.Get())
	}
}

func (l *LinkedList[T]) iterateBackwards(f func(T)) {
	for el := l.last; el != nil; el = el.prev {
		f(el.Get())
	}
}

func (l *LinkedList[T]) Iterate(f func(T)) {
	l.iterateForward(f)
}

func (l *LinkedList[T]) Length() (res int) {
	l.iterateForward(func(el T) {
		res++
	})
	return
}

func (l *LinkedList[T]) ToSlice() []T {
	res := make([]T, 0)
	l.iterateForward(func(el T) {
		res = append(res, el)
	})
	return res
}

func (l *LinkedList[T]) IsEmpty() bool {
	if el := l.first; el != nil {
		return false
	}
	return true
}

func (l *LinkedList[T]) findForward(i int) *element[T] {
	counter := 0
	el := l.first

	if el == nil {
		l.indexError()
	}

	for counter < i {
		if el == nil {
			l.indexError()
		}
		counter++
		el = el.next
	}
	return el
}

func (l *LinkedList[T]) findBackwards(i int) *element[T] {
	counter := -1
	el := l.last
	if el == nil {
		l.indexError()
	}

	for counter > i {
		if el == nil {
			l.indexError()
		}
		counter--
		el = el.prev
	}
	return el
}

func (l *LinkedList[T]) insertForward(i int, t T) {
	if l.IsEmpty() {
		if i != 0 {
			l.indexError()
		}
		new_ := newElement(t)
		l.first = new_
		l.last = new_
		return
	}

	el := l.findForward(i)
	if el == nil {
		new_ := newElement(t)
		new_.prev = l.last
		l.last.next = new_
		l.last = l.last.next
	} else {
		new_ := newElement(t)
		new_.prev = el.prev
		new_.next = el
		el.prev.next = new_
		el.prev = new_
	}
}

func (l *LinkedList[T]) insertBackwards(i int, t T) {
	if l.IsEmpty() {
		if i != -1 {
			l.indexError()
		}
		new_ := newElement(t)
		l.first = new_
		l.last = new_
		return
	}

	el := l.findBackwards(i)

	if el == nil {
		new_ := newElement(t)
		new_.next = l.first
		l.first.prev = new_
		l.first = l.first.prev
	} else {
		new_ := newElement(t)
		new_.prev = el
		new_.next = el.next
		el.next.prev = new_
		el.next = new_
	}

}

func (l *LinkedList[T]) Insert(i int, t T) {
	if i >= 0 {
		l.insertForward(i, t)
	} else {
		l.insertBackwards(i, t)
	}
}

func (l *LinkedList[T]) extractForward(i int) (res T) {
	el := l.findForward(i)
	if el == nil {
		l.indexError()
	}
	res = el.Get()
	if el.prev != nil {
		el.prev.next = el.next
	}
	if el.next != nil {
		el.next.prev = el.prev
	}
	return
}

func (l *LinkedList[T]) extractBackwards(i int) (res T) {
	el := l.findBackwards(i)

	if el == nil {
		l.indexError()
	}
	res = el.Get()
	if el.prev != nil {
		el.prev.next = el.next
	}
	if el.next != nil {
		el.next.prev = el.prev
	}
	return
}

func (l *LinkedList[T]) Extract(i int) T {
	if i >= 0 {
		return l.extractForward(i)
	} else {
		return l.extractBackwards(i)
	}
}

func (l *LinkedList[T]) getForward(i int) T {
	el := l.findForward(i)
	if el == nil {
		l.indexError()
	}
	return el.Get()
}

func (l *LinkedList[T]) getBackwards(i int) T {
	el := l.findBackwards(i)
	if el == nil {
		l.indexError()
	}
	return el.Get()
}

func (l LinkedList[T]) Get(i int) T {
	if i >= 0 {
		return l.getForward(i)
	} else {
		return l.getBackwards(i)
	}
}

func (l *LinkedList[T]) putForward(i int, t T) {
	el := l.findForward(i)
	if el == nil {
		l.indexError()
	}
	el.Set(t)
}

func (l *LinkedList[T]) putBackwards(i int, t T) {
	el := l.findBackwards(i)
	if el == nil {
		l.indexError()
	}
	el.Set(t)
}

func (l *LinkedList[T]) Put(i int, t T) {
	if i >= 0 {
		l.putForward(i, t)
	} else {
		l.putBackwards(i, t)
	}
}

func (l *LinkedList[T]) Clear() {
	if l.IsEmpty() {
		return
	}
	el := l.first.next

	for ; el != nil; el = el.next {
		if el.prev != nil {
			el.prev.next = nil
		}
		el.prev = nil
	}
	l.first = nil
	l.last = nil
}

func (l *LinkedList[T]) FromIterable(r root.Iterable[T]) {
	r.Iterate(func(el T) {
		l.Add(el)
	})
}

func (l *LinkedList[T]) Add(ts ...T) {
	for _, el := range ts {
		new_ := newElement(el)
		if l.IsEmpty() {
			l.first = new_
			l.last = new_
		} else {
			new_.prev = l.last
			l.last.next = new_
			l.last = new_
		}
	}
}

func (l *LinkedList[T]) Copy() (res List[T]) {
	res = NewLinkedList[T]()
	res.FromIterable(l)
	return
}
