// TODO Make use of it
package step

import (
	"fmt"

	"github.com/sj14/calc/tree"
)

var steps []string

// Check if a logic has been applied and if so, add it to the slice
func Check(oldTree, newTree tree.Tree) {
	fmt.Printf("oldtree: %v", oldTree)
	fmt.Printf("newtree: %v", newTree)

	// append step if the tree changed
	if oldTree.InOrder() != newTree.InOrder() {
		steps = append(steps, fmt.Sprintf("%v = %v", oldTree.InOrder(), newTree.InOrder()))
	}
}

// Get return the slice of steps
func Get() []string {
	fmt.Println(steps)
	return steps
}
