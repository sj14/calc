package aslice

import (
	"testing"

	"github.com/sj14/calc/tree"
)

func TestPush(t *testing.T) {
	var s Aslice
	want := tree.Tree{Value: "42"}
	s.Push(want)

	if r := s.Slice[0].Value; r != want.Value {
		t.Errorf("Push: got %v, want %v", r, want)
	}
}

func TestPop(t *testing.T) {
	var s Aslice
	want := tree.Tree{Value: "13"}
	s.Push(want)
	r, err := s.Pop()
	if err != nil {
		t.Error(err)
	}

	if r.Value != want.Value {
		t.Errorf("Pop: got %v, want %v", r, want)
	}
	if len(s.Slice) != 0 {
		t.Errorf("Pop: Slice should be empty, got %v", s.Slice[0])
	}
}

func TestLast(t *testing.T) {
	var s Aslice
	want := tree.Tree{Value: "67"}
	s.Push(tree.Tree{Value: "51"})
	s.Push(want)
	r, err := s.Last()
	if err != nil {
		t.Error(err)
	}
	if r.Value != want.Value {
		t.Errorf("Last: got %v, want %v", r, want)
	}
}

func TestFirst(t *testing.T) {
	var s Aslice
	want := tree.Tree{Value: "29"}
	s.Push(want)
	s.Push(tree.Tree{Value: "83"})
	r, err := s.First()
	if err != nil {
		t.Error(err)
	}
	if r.Value != want.Value {
		t.Errorf("Last: got %v, want %v", r, want)
	}
}

// func TestContains(t *testing.T) {
// 	var s Aslice
// 	s.Push(tree.Tree{Value: "58"})
// 	s.Push(tree.Tree{Value: "13"})
// 	s.Push(tree.Tree{Value: "38"})
// 	s.Push(tree.Tree{Value: "98"})
// 	s.Push(tree.Tree{Value: "83"})
//
// 	result1 := s.Contains("38")
//
// 	if result1 != true {
// 		t.Errorf("Contains: should contain 38 but got false")
// 	}
//
// 	result2 := s.Contains(99)
//
// 	if result2 != false {
// 		t.Errorf("Contains: should not contain 99 but got true")
// 	}
// }

// func TestContainsOneOf(t *testing.T) {
// 	var s Aslice
// 	s.Push(58)
// 	s.Push(13)
// 	s.Push(38)
// 	s.Push(98)
// 	s.Push(83)
//
// 	var n Aslice
// 	n.Push(34)
// 	n.Push(98)
// 	n.Push(45)
//
// 	result1 := s.ContainsAny(n.Slice)
//
// 	if result1 != true {
// 		t.Errorf("ContainsOneOf: should contain 98 but got false")
// 	}
//
// 	var m Aslice
// 	m.Push(34)
// 	m.Push(45)
// 	m.Push(56)
//
// 	result2 := s.ContainsAny(m.Slice)
//
// 	if result2 != false {
// 		t.Errorf("ContainsOneOf: should got false")
// 	}
// }
