package tree

import (
	"testing"
)

func TestHeight(t *testing.T) {
	root := Tree{Value: "1"}
	node1 := Tree{Value: "2"}
	node2 := Tree{Value: "3"}

	node1.Children = append(node1.Children, node2)
	root.Children = append(root.Children, node1)

	h := root.Height()

	if h != 3 {
		t.Errorf("TestHeight failed. wanted 3, got: %v", h)
	}
}

func TestInOrder(t *testing.T) {
	root := Tree{Value: "1"}
	node1 := Tree{Value: "2"}
	node2 := Tree{Value: "3"}

	node1.Children = append(node1.Children, node2)
	root.Children = append(root.Children, node1)
}
