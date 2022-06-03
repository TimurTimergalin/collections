// Package list provides List interface and its implementations: ArrayList and LinkedList
package list

import (
	root "github.com/TimurTimergalin/collections"
)

// List provides convenient interface to work with a sequence of elements
type List[T any] interface {
	root.Collection[T]
	// Insert inserts element at given position
	Insert(int, T)
	// Extract removes an element of the list and return it
	Extract(int) T
	// Get  returns element at position
	Get(int) T
	// Put sets the value of element at index
	Put(int, T)
	// Clear removes every element from the list
	Clear()
	// Copy create a copy of the list
	Copy() List[T]
}
