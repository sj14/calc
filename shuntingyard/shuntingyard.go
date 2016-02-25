package shuntingyard

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"text/scanner"

	"github.com/sj14/calc/aslice"
	"github.com/sj14/calc/operation"
	"github.com/sj14/calc/tree"
)

var number = regexp.MustCompile(`[0-9]+`)
var variable = regexp.MustCompile(`[a-z]`)

// InfixToAST converts an infix notation to an abstract syntax tree
func InfixToAST(input string) (tree.Tree, error) {
	var (
		output    aslice.Aslice // contains trees/nodes
		operators aslice.Aslice // contains trees/nodes
	)

	var scan scanner.Scanner
	scan.Init(strings.NewReader(input))
	var tok rune
	var str string

	for tok != scanner.EOF {
		tok = scan.Scan()
		str = scan.TokenText()
		fmt.Printf("ShuntingYard token: %v\n", str)
		switch {
		case number.MatchString(str):
			output.Push(tree.Tree{Value: str, Children: nil})
			fmt.Printf("Push %v to output\n", str)

		case operation.IsOperator(str):
			err := operatorLogic(str, &operators, &output, &scan)
			if err != nil {
				return tree.Tree{}, err
			}
			fmt.Printf("Push %v to operators\n", str)
			t := tree.Tree{Value: str, Children: nil}
			operators.Push(t)

		case str == "(":
			fmt.Printf("Push %v to operators\n", str)
			t := tree.Tree{Value: str, Children: nil}
			operators.Push(t)

		case str == ")":
			fmt.Printf("found %v\n", str)
			// Create nodes from stack and output until
			// we reach the opening parenthese at the stack.
			for stop := false; stop != true; {
				opLast, err := operators.Last()
				if err != nil {
					return tree.Tree{}, err
				}
				if opLast.Value != "(" {
					fmt.Printf("opLast: %v and not (\n", opLast)
					err := operatorToOutput(&operators, &output, &scan)
					if err != nil {
						return tree.Tree{}, err
					}
				} else {
					fmt.Printf("opLast: (\n")
					stop = true
				}
			} // for loop

			// pop the left parenthese
			_, err := operators.Pop()
			if err != nil {
				return tree.Tree{}, err
			}

		case operation.IsFunction(str): //str == "max":
			fmt.Println("FUNCTION")

			// Get substring which is surrounded by parentheses.
			substring, posTo, err := parenthesisSubstring(input, scan.Pos().Offset)
			if err != nil {
				return tree.Tree{}, err
			}

			fmt.Printf("SUBSTRING: %v\n", substring)

			// Recursion to create a seperate tree of the substring calculation.
			subtree, err := InfixToAST(substring)
			if err != nil {
				return tree.Tree{}, err
			}
			fmt.Printf("SUBTREE sy: %v\n", subtree)

			if subtree.Value == "" {
				// InfixToAST returned a new root node with children
				// but doesn't know the operation/value at the root.
				subtree.Value = str
				output.Push(subtree)
			} else {
				// Behavior on Functions with 1 Argument, e.g. sqrt(x)
				// Because Subtree just returned a single node with x.
				root := tree.Tree{Value: str}
				root.Children = append(root.Children, subtree)
				output.Push(root)
			}
			// forward scanner to skip already handled substring
			input = input[posTo+1:]
			scan.Init(strings.NewReader(input))

		case variable.MatchString(str):
			output.Push(tree.Tree{Value: str, Children: nil})
			fmt.Printf("Push %v to output\n", str)

		case str == ",":
			// Function Seperator, this means we are inside a Function-Substring,
			// build at "case: isFunction".
			opLast, err := operators.Last()
			if err != nil {
				return tree.Tree{}, err
			}
			if opLast.Value != "(" {
				// e.g. max(4+2, 3)
				// where 4+2 gets first on the output stack
				// 		Max
				// 		/ \
				// 	 +  3
				//  / \
				// 	4 2
				err = operatorToOutput(&operators, &output, &scan)
				if err != nil {
					return tree.Tree{}, err
				}
			}
		} // switch
	} // scanner
	for operators.Len() > 0 {
		// TODO check error, which is not nil because can't find '(' as operation.
		operatorToOutput(&operators, &output, &scan)
	}

	if output.Len() > 1 {
		var outTree tree.Tree

		for i := range output.Slice {
			//fmt.Println(o)
			// outTree.Children = append(outTree.Children, output.Slice[i].(tree.Tree))
			outTree.Children = append(outTree.Children, output.Slice[i])
		}
		return outTree, nil
	}

	outTree, err := output.Pop()
	if err != nil {
		return tree.Tree{}, err
	}
	fmt.Println(outTree)
	return outTree, nil
}

// Pushes operator with children (popped from output) to output.
func operatorToOutput(operators, output *aslice.Aslice, scan *scanner.Scanner) error {
	// Get Last Operator from Stack
	temp, err := operators.Pop()
	if err != nil {
		return err
	}
	// Get corresponding Type Operator
	op, err := operation.Get(temp.Value)
	if err != nil {
		return err
	}
	// Create root
	node := tree.Tree{Value: temp.Value}

	// If the operation has a fixed number of operands,
	// pop numbers from output as children to operation
	// and push operation with childs to output.
	if op.Operands != 0 {
		if output.Len() == 1 && temp.Value == "-" {
			// if there is a leading minus (e.g. -4-3) add 0 to the tree: 0-4-3
			child, err := output.Pop()
			if err != nil {
				return err
			}
			node.Children = append(node.Children, child)
			zero := tree.Tree{Value: "0"}
			node.Children = append(node.Children, zero)
		} else if output.Len() == 1 {
			if scan.TokenText() == "-" {
				//2*-3 gets to 0-2*3 = -6
				child1, err := output.Pop()
				if err != nil {
					return err
				}

				_ = scan.Scan()
				minusval := scan.TokenText()
				fmt.Printf("Minusval: %v\n", minusval)

				child0 := tree.Tree{Value: minusval}
				// Order of append is importend,
				// otherwise divide operands would have been switched.
				node.Children = append(node.Children, child0) // value after minus symbol (from input)
				node.Children = append(node.Children, child1) // last value from the output stack
			}
		} else {
			// no leading minus
			for i := uint8(0); i < op.Operands; i++ {
				child, err := output.Pop()
				if err != nil {
					return err
				}
				node.Children = append(node.Children, child)
			}
		}
	} else {
		// The operation has a variable number of inputs (probably function)
		// Check if last root node value from output is a number.
		// If it is a number, pop from output and append as child to new root ('node').
		// Repeat until last element is not a number.
		for {
			// Output is not empty
			if output.Len() <= 0 {
				break
			}

			outNode, err := output.Last()
			if err != nil {
				return err
			}

			if !number.MatchString(outNode.Value) {
				break
			}

			// last element from output is a number.
			// pop it from the output and add as child.
			fmt.Println("IS NUMBER")
			child, err := output.Pop()
			if err != nil {
				return err
			}
			node.Children = append(node.Children, child)
		} // for loop
	}
	fmt.Printf("Push %v to output with children %#v\n",
		node.Value, node.Children)
	node.PostOrder()

	output.Push(node)
	return nil
}

// Pushes operators to output if current prescendence is lower-equal
// to last operator on stack and last op not right associative
// OR prescendence is lower and last op is right associative.
// Repeat until stack empty or last op doesn't match one of the conditions.
func operatorLogic(str string, operators, output *aslice.Aslice, scan *scanner.Scanner) error {
	for {
		if operators.Len() <= 0 {
			// output is empty
			return nil
		}

		opLast, err := operators.Last()
		if err != nil {
			return err
		}

		if !operation.IsOperation(opLast.Value) {
			// is no Operation
			return nil
		}

		op1, err := operation.Get(str)
		if err != nil {
			return err
		}

		op2, err := operation.Get(opLast.Value)
		if err != nil {
			return err
		}

		if (op2.RightAssociative == false && op1.Precedence <= op2.Precedence) ||
			(op2.RightAssociative == true && op1.Precedence < op2.Precedence) {
			err := operatorToOutput(operators, output, scan)
			if err != nil {
				return err
			}
		} else {
			return nil
		}
	} // for loop
}

// Find substring inside parenthesis.
// If there are n opening parenthesis
// it searches for n closing parenthesis.
func parenthesisSubstring(s string, offset int) (string, int, error) {
	count := 0
	posTo := 0
	for i, o := range s[offset:] {
		if o == '(' {
			fmt.Println("FOUND (")
			count++
		}
		if o == ')' {
			fmt.Println("FOUND )")
			count--
		}
		if count == 0 {
			fmt.Printf("AMOUNT ( = ), POS: %v\n", i)
			posTo = i
			break
		}
	}
	if posTo == 0 {
		return "", 0, errors.New("Missing closing parenthesis.")
	}
	posTo += offset
	substring := s[offset:posTo]
	return substring, posTo, nil
}
