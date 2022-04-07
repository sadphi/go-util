package deque

import "errors"

// A double-ended queue.
type Deque[T any] struct {
	elements []T
}

// Appends the element i to the end of the deque.
func (d *Deque[T]) append(i T) {
	d.elements = append(d.elements, i)
}

// Returns the first element of the deque.
func (d *Deque[T]) getFirst() (T, error) {
	if len(d.elements) > 0 {
		return d.elements[0], nil
	}
	var invalid T
	return invalid, errors.New("Could not get first element; deque is empty.")
}

// Pops the first element of the deque and returns the element.
//
// Returns an error if the deque is empty.
func (d *Deque[T]) popLeft() (T, error) {
	var res T
	if len(d.elements) > 0 {
		res = d.elements[0]
		d.elements = d.elements[1:]
		return res, nil
	}
	return res, errors.New("Could not pop first element; deque is empty.")
}

// Returns the last element of the deque.
//
// Returns an error if the deque is empty.
func (d *Deque[T]) getLast() (T, error) {
	length := len(d.elements)
	if length > 0 {
		return d.elements[length-1], nil
	}
	var invalid T
	return invalid, errors.New("Could not get first element; deque is empty.")
}

// Removes the last element of the deque and returns the element.
func (d *Deque[T]) popRight() (T, error) {
	var res T
	length := len(d.elements)
	if length > 0 {
		res = d.elements[length-1]
		d.elements = d.elements[:length-1]
		return res, nil
	}
	return res, errors.New("Could not pop last element; deque is empty.")
}

// Returns true if the deque is empty, false otherwise.
func (d *Deque[T]) isEmpty() bool {
	return len(d.elements) == 0
}
