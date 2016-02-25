package eval

import (
	"fmt"
	"strconv"

	"github.com/sj14/calc/tree"
)

func variableAndOperation(t, subtree *tree.Tree, steps *[]string, v string) error {
	fmt.Printf("Variable and Operation: %v\n", subtree)
	//	 	*
	//	 / \
	// 	a		:
	//		/  \
	// 		b		a
	for i := range subtree.Children {
		err := variableAndOperation(t, &subtree.Children[i], steps, v)
		if err != nil {
			return err
		}
	}
	op := t.Value
	oldTree := t.InOrder()
	fmt.Printf("op: %v, subtree: %v\n", op, subtree.Value)
	switch op {
	case "+":
		fmt.Println("VAR+OP: PLUS")
		switch subtree.Value {
		case "*":
			fmt.Println("VAR+OP: PLUS MULTI")
			err := varOpIncSubtree(t, subtree, v)
			if err != nil {
				return err
			}
		case "+":
			fmt.Println("VAR+OP: PLUS PLUS")
			varOp2MultiTree(t, subtree, v)
		case "-":
			fmt.Println("VAR+OP: PLUS MINUS")
			varOp2MultiTree(t, subtree, v)
		case "/", ":":
			// nothing to do
			// b+4/b keeps
		case "^":
			// nothing to do
			// b+4^b keeps
		} // Switch subtree
	case "-":
		fmt.Printf("VAR+OP: MINUS subtree: %v\n", subtree)
		switch subtree.Value {
		case "-":
			fmt.Println("VAR+OP: MINUS MINUS")
			// a-5-a = 0-5 	§1
			// -a-a = -2*a 	§2
			// TODO -a-5-a = -2*a-5

			// § 1
			i := -1
			if subtree.Children[0].Value == v && subtree.Children[1].Value != "0" {
				i = 0
			} else if subtree.Children[1].Value == v && subtree.Children[0].Value != "0" {
				i = 1
			}

			if i != -1 {
				subtree.Children[i].Value = "0"
				replaceRoot(t, v)
			} else {
				// § 2
				j := -1
				if subtree.Children[0].Value == v && subtree.Children[1].Value == "0" {
					j = 1
				} else if subtree.Children[1].Value == v && subtree.Children[0].Value == "0" {
					j = 0
				}

				if j != -1 {
					subtree.Value = "*"
					subtree.Children[j].Value = "-2"
					replaceRoot(t, v)
				}
			}

		case "*":
			fmt.Println("VAR+OP: MINUS MULTI")

			// 5*a-a = 4*a
			// a-5*a = -4*a 	§1

			i := -1
			if subtree.Children[0].Value == v {
				i = 1
			} else if subtree.Children[1].Value == v {
				i = 0
			}

			if i != -1 {
				num, err := strconv.ParseFloat(subtree.Children[i].Value, 64)
				if err != nil {
					return err
				}
				num--
				if num == 1 {
					// e.g. 2*a-a = a not 1*a
					subtree.Children = nil
					subtree.Value = v
				} else {

					if i == 1 {
						// §1
						num *= -1
					}
					subtree.Children[i].Value = strconv.FormatFloat(num, 'G', -1, 64)
				}
				replaceRoot(t, v)
			}
		case "/", ":":
			fmt.Println("VAR+OP: MINUS DIV")
			// Nothing to simplify
			// 4/a-a keeps the same
		case "+":
			fmt.Println("VAR+OP: MINUS PLUS")
			i := -1
			if subtree.Children[0].Value == v {
				i = 1
			} else if subtree.Children[1].Value == v {
				i = 0
			}

			if i != -1 {
				*subtree = subtree.Children[i]
				replaceRoot(t, v)
			}
		case "^":
			fmt.Println("VAR+OP: MINUS EXP")
			// Nothing to simplify
			// 4^a-a keeps the same
		} // switch subtree
	case "*":
		fmt.Println("VAR+OP: MULTI")
		switch subtree.Value {
		case "^":
			fmt.Println("VAR+OP: MULTI EXP")

			if subtree.Children[1].Value == v && subtree.Children[0].Value != v {
				// check only right child because of §3
				// and not a*a^a
				err := varOpIncSubtree(t, subtree, v)
				if err != nil {
					return err
				}
			}

		case "-":
			fmt.Println("VAR+OP: MULTI MINUS")
		case "+":
			fmt.Println("VAR+OP: MULTI PLUS")
		case "*":
			fmt.Println("VAR+OP: MULTI MULTI")
			//	 	* 	(gets deleted) §1
			//	 / \
			// 	a		*
			//		/  \
			// 		b		a*
			//===================
			//		*
			//	 / \
			//	b		^ 	(former a*) §2
			//		 / \
			//    2  a (appended to former a*)
			//
			// e.g. a*b*a = a^2*b
			//			a*5*a = a^2*5

			varOp2ExpTree(t, subtree, v)
		case "/", ":":
			fmt.Println("VAR+OP: MULTI DIV")
			//	 	* 	(gets deleted) §1
			//	 / \
			// 	a		:
			//		/  \
			// 		b		a*
			//===================
			//		:
			//	 / \
			//	b		^ 	(former a*) §2
			//		 / \
			//    2  a (appended to former a*)
			//
			// e.g. a/b*a = a^2/b
			//			a/5*a = a^2/5

			varOp2ExpTree(t, subtree, v)
		} // switch subtree
	case "/", ":":
		fmt.Println("VAR+OP: DIVIDE")
		switch subtree.Value {
		case "/", ":":
			fmt.Println("VAR+OP: DIVIDE DIVIDE")
			//	 	: 	(gets deleted) §1
			//	 / \
			// 	a		:
			//		/  \
			// 		b		a*
			//===================
			//		:
			//	 / \
			//	b		1 	(former a*) §2
			//
			// e.g. a/b/a = 1/b
			//			a/5/a = 1/5
			i := -1
			if subtree.Children[0].Value == v {
				i = 0
			} else if subtree.Children[1].Value == v {
				i = 1
			}

			if i != -1 {
				subtree.Children[i].Value = "1"
				replaceRoot(t, v)
			}
		case "^":
			fmt.Println("VAR+OP: DIVIDE EXP")
			//	 	: 	(gets deleted) §1
			//	 / \
			// 	a		^
			//		/  \
			// 		2		a
			//===================
			//
			//		a 	(former ^) §2 (a^1 gets to a)
			//
			////////////////////////////
			// Example 2
			////////////////////////////
			//	 	: 	(gets deleted) §1
			//	 / \
			// 	a		^
			//		/  \
			// 		3		a
			//===================
			//	 	:
			//	 / \
			// 	a		^
			//		/  \
			// 		2		a (decrement 3 to 2) §4
			//
			// e.g. a^3/a = a^2/a
			//
			// but	2^a/a still 2^a/a (§3)

			i := -1
			if subtree.Children[1].Value == v && subtree.Children[0].Value != v {
				// check only right child because of §3
				// and not a/a^a
				i = 0
			}

			if i != -1 {
				num, err := strconv.ParseFloat(subtree.Children[i].Value, 64)
				if err != nil {
					return err
				}
				num--
				if num == 1 {
					// §2: a^2/a = a not a^1
					subtree.Children = nil
					subtree.Value = v
				} else {
					// §4
					subtree.Children[i].Value = strconv.FormatFloat(num, 'G', -1, 64)
				}
				replaceRoot(t, v)
			}
		case "+":
			fmt.Println("VAR+OP: DIVIDE PLUS")
			// Nothing to simplify
			// 2+3/a keeps the same
		case "-":
			fmt.Println("VAR+OP: DIVIDE MINUS")
			// Nothing to simplify
			// 2-3/a keeps the same
		case "*":
			fmt.Println("VAR+OP: DIV MULTI")
			//	 	: 	(gets deleted) §1
			//	 / \
			// 	a		*
			//		/  \
			// 		b		a
			//===================
			//		b (former *) §2
			//
			// e.g. a*b/a = b
			//			a*5/a = 5

			i := -1
			if subtree.Children[0].Value == v {
				i = 1
			} else if subtree.Children[1].Value == v {
				i = 0
			}

			if i != -1 {
				*subtree = subtree.Children[i] // §2
				replaceRoot(t, v)
			}
		} // switch subtree
	} // Switch op

	// append step if the tree changed
	if oldTree != t.InOrder() {
		*steps = append(*steps, fmt.Sprintf("%v = %v", oldTree, t.InOrder()))
	}

	return nil
}

func varOpIncSubtree(t, subtree *tree.Tree, v string) error {
	//	 	+
	//	 / \
	// 	a		*
	//		/  \
	// 		2		a
	//===================
	//	 	+
	//	 / \
	// 	a		*
	//		/  \
	// 		3		a

	i := -1
	if subtree.Children[0].Value == v {
		i = 1
	} else if subtree.Children[1].Value == v {
		// subtree.Children[1].Value = subtree.Children[0].Value
		// subtree.Children[0].Value = v
		i = 0
	}

	if i != -1 {
		num, err := strconv.ParseFloat(subtree.Children[i].Value, 64)
		if err != nil {
			return err
		}
		num++
		subtree.Children[i].Value = strconv.FormatFloat(num, 'G', -1, 64)
		replaceRoot(t, v)
	}
	return nil
}

func varOp2MultiTree(t, subtree *tree.Tree, v string) {
	//	 	- 	(gets deleted) §1
	//	 / \
	// 	a		-
	//		/  \
	// 		b		a*
	//===================
	//	 	-
	//	 / \
	// 	b		* (former a*) §2
	//		/  \
	// 		2		a
	//
	// e.g. a-b-a = b-2*a
	//			a+5+a = 5+2*a

	i := -1
	if subtree.Children[0].Value == v {
		i = 0
	} else if subtree.Children[1].Value == v {
		i = 1
	}

	if i != -1 {
		subtree.Children[i].Value = "*" // §2
		multi := tree.Tree{Value: "2"}
		vari := tree.Tree{Value: v}
		subtree.Children[i].Children = append(subtree.Children[i].Children, vari)
		subtree.Children[i].Children = append(subtree.Children[i].Children, multi)
		replaceRoot(t, v)
	}
}

func varOp2ExpTree(t, subtree *tree.Tree, v string) {
	//	 	* 	(gets deleted) §1
	//	 / \
	// 	a		:*
	//		/  \
	// 		b		a*
	//===================
	//		:*
	//	 / \
	//	b		^ 	(former a*) §2
	//		 / \
	//    2  a (appended to former a*)
	//
	// e.g. a/b*a = a^2/b
	//			a/5*a = a^2/5
	// e.g. a*b*a = a^2*b
	//			a*5*a = a^2*5
	i := -1
	if subtree.Children[0].Value == v {
		i = 0
	} else if subtree.Children[1].Value == v {
		i = 1
	}

	if i != -1 {
		subtree.Children[i].Value = "^" // §2
		subtree.Children[i].Children = append(subtree.Children[i].Children, tree.Tree{Value: "2"})
		subtree.Children[i].Children = append(subtree.Children[i].Children, tree.Tree{Value: v})
		replaceRoot(t, v) // §1
	}
}

func replaceRoot(t *tree.Tree, v string) {
	if t.Children[0].Value == v {
		*t = t.Children[1]
	} else {
		*t = t.Children[0]
	}
}
