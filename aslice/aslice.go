// Package aslice offers some additional
// methods on slices (advanced slices).
package aslice

import (
	"errors"

	"github.com/sj14/calc/tree"
)

// Aslice are Slices with
// some helping methods.
type Aslice struct {
	Slice []tree.Tree
}

// Push appends an element to the slice.
func (p *Aslice) Push(e tree.Tree) {
	p.Slice = append(p.Slice, e)
}

// Pop returns the last element and
// removes it from the slice.
func (p *Aslice) Pop() (tree.Tree, error) {
	l, err := p.Last()
	if err != nil {
		return tree.Tree{}, err
	}
	p.Slice = p.Slice[:len(p.Slice)-1]
	return l, nil
}

// Last return the last element
// without removing it from the slice.
func (p *Aslice) Last() (tree.Tree, error) {
	if len(p.Slice) > 0 {
		return p.Slice[len(p.Slice)-1], nil
	}
	return tree.Tree{}, errors.New("Slice is empty")
}

// First returns the first element
// without removing it from the slice.
func (p *Aslice) First() (tree.Tree, error) {
	if len(p.Slice) > 0 {
		return p.Slice[0], nil
	}
	return tree.Tree{}, errors.New("Slice is empty")
}

// Len returns the length of the slice.
func (p *Aslice) Len() int {
	return len(p.Slice)
}
