package tree

import (
	"fmt"
	"regexp"
	"strconv"
)

var variable = regexp.MustCompile(`[a-z]`)

// Tree Structure with variable amount of children
type Tree struct {
	Value    string
	Children []Tree
}

// Height returns the number of levels of the tree. If empty, it returns 0.
func (t *Tree) Height() int {
	i := 1
	if t == nil {
		return 0
	}
	for _, o := range t.Children {
		i = o.Height()
		i++
	}
	return i
}

// Contains checks if the tree has an Element with value s.
func (t *Tree) Contains(s string) bool {
	if t.Value == s {
		return true
	}
	for _, o := range t.Children {
		o.Contains(s)
	}
	return false
}

// ContainsVariable checks if the tree has a Variable as Element Value
func (t *Tree) ContainsVariable() bool {
	if variable.MatchString(t.Value) {
		return true
	}
	for _, o := range t.Children {
		o.ContainsVariable()
	}
	return false
}

// PostOrder prints the root last
func (t *Tree) PostOrder() {
	for _, o := range t.Children {
		fmt.Printf("Going down %v Children: %v\n", t.Value, t.Children)
		o.PostOrder()
	}
	fmt.Printf("Post-Order: %v Children: %v\n", t.Value, t.Children)
}

// InOrder returns a string of the tree. Works only for max. 2 children.
func (t *Tree) InOrder() string {
	s := ""
	if t == nil {
		return "0"
	}
	if t.Children != nil && len(t.Children) > 1 && t.Children[1].Value != "" {
		s += t.Children[1].InOrder()
	}
	s += t.Value
	if t.Children != nil && t.Children[0].Value != "" {
		s += t.Children[0].InOrder()
	}
	return s
}

// PreOrder prints the root first
func (t *Tree) PreOrder() {
	fmt.Printf("Pre-Order: %v Children: %v\n", t.Value, t.Children)

	for _, o := range t.Children {
		fmt.Printf("Going down %v\n", t.Value)
		o.PreOrder()
	}
}

// ChildrenToString Converts this: [{1 []} {2 []} {3 []}] to: 1, 2, 3
// Needed for pretty step output of functions, e.g. min(1,2,3).
func (t *Tree) ChildrenToString() string {
	var s string
	for i, o := range t.Children {
		if i < len(t.Children)-1 {
			s += o.Value + ", "
		} else {
			s += o.Value
		}
	}
	return s
}

// ChildrenToFloat returns all children as []float64
func (t *Tree) ChildrenToFloat() ([]float64, error) {
	var numbers []float64

	for _, c := range t.Children {
		v, err := strconv.ParseFloat(c.Value, 64)
		if err != nil {
			return []float64{0}, err
		}
		numbers = append(numbers, v)
	}
	return numbers, nil
}
