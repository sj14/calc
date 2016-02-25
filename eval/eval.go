package eval

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/sj14/calc/math2"
	"github.com/sj14/calc/operation"
	"github.com/sj14/calc/tree"
)

var variable = regexp.MustCompile(`[a-z]`)
var number = regexp.MustCompile(`[0-9]+`)

// AST calculates the result of the input tree.
// Returns the result as string, the steps as slice of strings and an error.
func AST(t *tree.Tree) (string, []string, error) {
	steps, err := astRec(t, &[]string{})
	if err != nil {
		return "", []string{}, err
	}
	if t.Height() == 1 {
		return t.Value, steps, nil
	}
	return t.InOrder(), steps, nil
}

func astRec(t *tree.Tree, steps *[]string) ([]string, error) {
	fmt.Printf("Step %v: %v\n", 0, t)

	for i := range t.Children {
		_, err := astRec(&t.Children[i], steps)
		if err != nil {
			return []string{""}, err
		}
	}

	if operation.IsOperation(t.Value) {
		op, err := operation.Get(t.Value)
		if err != nil {
			return []string{""}, err
		}

		// Binary Operation
		if op.Operands == 2 {
			if len(t.Children) != 2 {
				return []string{""}, fmt.Errorf("Binary Operation with %v operands", len(t.Children))
			}
			if operation.IsOperation(t.Children[0].Value) && operation.IsOperation(t.Children[1].Value) {
				// left + right = operation
				// booth sides have a variable
				// e.g. a*5+a*2 = a*7
				// or   a*b+a*b = 2*a*b
				// or   a*b+c*d = a*b+c*d

				fmt.Printf("OPERATION %v AND %v ON %v\n", t.Children[0].Value, t.Children[1].Value, t.Value)
				fmt.Println(t.InOrder())

				err := operationAndOperation(t, &t.Children[0], &t.Children[1], steps)

				if err != nil {
					return []string{""}, err
				}
				return *steps, nil

			} else if variable.MatchString(t.Children[0].Value) && t.Children[0].Value == t.Children[1].Value {
				// left + right = same variable
				err := sameVariable(t, steps)
				if err != nil {
					return []string{""}, err
				}
				return *steps, nil
			} else if (variable.MatchString(t.Children[0].Value) && operation.IsOperation(t.Children[1].Value)) ||
				(variable.MatchString(t.Children[1].Value) && operation.IsOperation(t.Children[0].Value)) {
				// variable and operator (subtree)
				// which means, subtree has a variable as child, too.
				// Otherwise there would be no subtree.
				var v string
				var subtree *tree.Tree
				if variable.MatchString(t.Children[0].Value) {
					v = t.Children[0].Value
					subtree = &t.Children[1]
				} else {
					v = t.Children[1].Value
					subtree = &t.Children[0]
				}

				err := variableAndOperation(t, subtree, steps, v)
				if err != nil {
					return []string{""}, err
				}
				return *steps, nil
			} else if variable.MatchString(t.Children[0].Value) || variable.MatchString(t.Children[1].Value) {
				// left or right = different variable
				// nothing to simplyfy
				return *steps, nil
			} else if (operation.IsOperation(t.Children[0].Value) && number.MatchString(t.Children[1].Value)) ||
				(operation.IsOperation(t.Children[1].Value) && number.MatchString(t.Children[0].Value)) {
				// number and operation
				fmt.Printf("Number and Operation: %v\n", t)
				var v string
				var subtree *tree.Tree
				if number.MatchString(t.Children[0].Value) {
					v = t.Children[0].Value
					subtree = &t.Children[1]
				} else if number.MatchString(t.Children[1].Value) {
					v = t.Children[1].Value
					subtree = &t.Children[0]
				}

				n, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return []string{""}, err
				}

				err = numberAndOperation(t, subtree, steps, n)
				if err != nil {
					return []string{""}, err
				}
			} else {
				// number and number
				err := numberAndNumber(t, steps)
				if err != nil {
					return []string{""}, err
				}
			}
			// functions with any amount of operands
		} else if operation.IsFunction(t.Value) {
			t, *steps, err = evalFunction(t, steps)
		} // IsFunction
	} // IsOperation
	return *steps, nil
}

func evalFunction(t *tree.Tree, steps *[]string) (*tree.Tree, []string, error) {
	var result float64

	switch t.Value {
	case "max":
		compare, err := t.ChildrenToFloat()
		if err != nil {
			return &tree.Tree{}, []string{""}, err
		}
		result = math2.Max(compare)
	case "min":
		compare, err := t.ChildrenToFloat()
		if err != nil {
			return &tree.Tree{}, []string{""}, err
		}
		result = math2.Min(compare)

	case "sqrt":
		fmt.Println("SQRT!")
		n, err := strconv.ParseFloat(t.Children[0].Value, 64)
		if err != nil {
			return &tree.Tree{}, []string{""}, err
		}
		result = math.Sqrt(n)
	default:
		return &tree.Tree{}, []string{""}, fmt.Errorf("Function %v not yet implemented.", t.Value)
	} // Switch

	// append step
	childStr := t.ChildrenToString()
	*steps = append(*steps, fmt.Sprintf("%v(%v) = %v", t.Value, childStr, result))

	// Set Result in Root node.
	t.Value = strconv.FormatFloat(result, 'G', -1, 64)
	// Remove children.
	t.Children = nil
	return t, *steps, nil
}

func numberAndNumber(t *tree.Tree, steps *[]string) error {
	//	 	*
	//	 / \
	// 	2		5
	var result float64
	right, err := strconv.ParseFloat(t.Children[0].Value, 64)
	if err != nil {
		return err
	}
	left, err := strconv.ParseFloat(t.Children[1].Value, 64)
	if err != nil {
		return err
	}

	switch t.Value {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "/", ":":
		result = left / right
	case "*":
		result = left * right
	case "^":
		result = math.Pow(left, right)
	} // Switch

	// Add step
	*steps = append(*steps, fmt.Sprintf("%v %v %v = %v", left, t.Value, right, result))
	// Save Result in Root node.
	t.Value = strconv.FormatFloat(result, 'G', -1, 64)
	// Remove children.
	t.Children = nil
	return nil
}

func sameVariable(t *tree.Tree, steps *[]string) error {
	fmt.Printf("SAME VARIABLE %v\n", t)
	oldTree := t.InOrder()

	// left + right = same variable
	switch t.Value {
	case "+":
		// a + a = 2 * a
		t.Value = "*"
		t.Children[1].Value = "2"
	case "-":
		// a-a = 0
		t.Value = "0"
		t.Children = nil
	case "/", ":":
		// a/a = 1
		t.Value = "1"
		t.Children = nil
	case "*":
		// a*a = a^2
		t.Value = "^"
		t.Children[0].Value = "2"
		// case "^":
		// nothing to simplify
	}
	fmt.Printf("SAME VARIABLE RESULT: %v\n", t)

	// append step if the tree changed
	if oldTree != t.InOrder() {
		*steps = append(*steps, fmt.Sprintf("%v = %v", oldTree, t.InOrder()))
	}

	return nil
}
