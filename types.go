// Package collections provides useful collections such as lists or set
// subpackages:
// list
// set
package collections

import "fmt"

// Iterable represents an object that can iterated through
type Iterable[T any] interface {
	// Iterate calls function for every element of iterable
	Iterate(func(T))
}

// Collection represents an object that stores data
type Collection[T any] interface {
	fmt.Stringer
	Iterable[T]
	// Length returns size of collection
	Length() int
	// ToSlice converts collection of T to []T
	ToSlice() []T
	// IsEmpty checks if collection contains any elements
	IsEmpty() bool
	// FromIterable adds all elements from iterable to collection
	FromIterable(Iterable[T])
	// Add adds elements to collection
	Add(...T)
}
