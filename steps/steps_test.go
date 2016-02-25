package step

import (
	"testing"

	"github.com/sj14/calc/tree"
)

func TestCheck(t *testing.T) {
	want := []string{"1", "2"}
	oldTree := tree.Tree{Value: "1"}
	newTree := tree.Tree{Value: "2"}

	Check(oldTree, newTree)
	if Get() == nil {
		t.Errorf("TestCheck failed. wanted %v, got: %v", want, Get())
	}

}
