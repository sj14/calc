package eval

import (
	"fmt"
	"math"
	"strconv"

	"github.com/sj14/calc/tree"
)

func numberAndOperation(t, subtree *tree.Tree, steps *[]string, n float64) error {

	fmt.Printf("Number and Operation: %v\n", subtree)
	//	 	*
	//	 / \
	// 	5		*
	//		/  \
	// 		a		3
	for i := range subtree.Children {
		err := numberAndOperation(t, &subtree.Children[i], steps, n)
		if err != nil {
			return err
		}
	}
	op := t.Value
	oldTree := t.InOrder()
	fmt.Printf("op: %v, subtree: %v\n", op, subtree.Value)
	switch op {
	case "+":
		fmt.Println("NUM+OP: PLUS")
		switch subtree.Value {
		case "*":
			fmt.Println("NUM+OP: PLUS MULTI")
			// nothing to do
			// b*4+5 still the same
		case "/":
			fmt.Println("NUM+OP: PLUS DIV")
			// nothing to do
			// b/4+5 still the same
		case "^":
			fmt.Println("NUM+OP: PLUS EXP")
			// nothing to do
			// b^4+5 still the same
		case "+":
			fmt.Println("NUM+OP: PLUS PLUS")
			i := -1
			if number.MatchString(subtree.Children[0].Value) {
				i = 0
			} else if number.MatchString(subtree.Children[1].Value) {
				i = 1
			}

			if i != -1 {
				val, err := strconv.ParseFloat(subtree.Children[i].Value, 64)
				if err != nil {
					return err
				}
				val += n
				str := strconv.FormatFloat(val, 'f', -1, 64)
				subtree.Children[i].Value = str

				*t = *subtree
			}
		case "-":
			fmt.Println("NUM+OP: PLUS MINUS")
			//	 	+
			//	 / \
			// 	5		-
			//		/  \
			// 	 4	 a

			i := -1
			if number.MatchString(subtree.Children[0].Value) {
				i = 0
			} else if number.MatchString(subtree.Children[1].Value) {
				i = 1
			}

			if i != -1 {
				val, err := strconv.ParseFloat(subtree.Children[i].Value, 64)
				if err != nil {
					return err
				}
				if i == 0 {
					// a-5+4 = a-1
					// a-4+5 = a+1 §1
					val = n - val // second num - first num
					if val > 0 {
						// §1
						subtree.Value = "+"
					}
					val = math.Abs(val)
				} else {
					// 5-a+2 = 7-a
					val += n
				}
				str := strconv.FormatFloat(val, 'f', -1, 64)
				subtree.Children[i].Value = str
				*t = *subtree
			}
		} // Switch subtree
	case "-":
		fmt.Printf("NUM+OP: MINUS subtree: %v\n", subtree)
		switch subtree.Value {
		case "-":
			fmt.Println("NUM+OP: MINUS MINUS")

			i := -1
			if number.MatchString(subtree.Children[0].Value) {
				i = 0
			} else if number.MatchString(subtree.Children[1].Value) {
				i = 1
			}

			if i != -1 {
				val, err := strconv.ParseFloat(subtree.Children[i].Value, 64)
				if err != nil {
					return err
				}

				if i == 1 {
					// 5-a-2 = 3-a
					val -= n // first num - second num
				} else {
					// a-4-3 = a-7
					val += n
				}
				str := strconv.FormatFloat(val, 'f', -1, 64)
				subtree.Children[i].Value = str
				*t = *subtree
			}
		case "*":
			fmt.Println("NUM+OP: MINUS MULTI")
			// nothing to simplify
			// b*4-3 still the same
		case "/":
			fmt.Println("NUM+OP: MINUS DIV")
			// nothing to simplify
			// b/4-3 still the same
		case "^":
			fmt.Println("NUM+OP: MINUS EXP")
			// nothing to simplify
			// b^4-3 still the same
		case "+":
			fmt.Println("NUM+OP: MINUS PLUS")
			// a+4-5 = a-1
			// 4+a-5 = -1+a 	§1

			i := -1
			if number.MatchString(subtree.Children[0].Value) {
				i = 0
			} else if number.MatchString(subtree.Children[1].Value) {
				i = 1
			}

			if i != -1 {
				val, err := strconv.ParseFloat(subtree.Children[i].Value, 64)
				if err != nil {
					return err
				}
				val -= n
				if i == 0 {
					// e.g. b+4-5 = b-1 not b+-1
					// and §1
					if val < 0 {
						val = math.Abs(val)
						subtree.Value = "-"
					}
				}
				str := strconv.FormatFloat(val, 'f', -1, 64)
				subtree.Children[i].Value = str
				*t = *subtree
			}
		} // switch subtree
	case "*":
		fmt.Println("NUM+OP: MULTI")
		switch subtree.Value {
		case "^":
			fmt.Println("NUM+OP: MULTI EXP")
			// nothing to simplify
			// b^4*3 still the same
		case "-":
			fmt.Println("NUM+OP: MULTI MINUS")
			// not possible?
			// b-4*3 gets directly to b-12
		case "+":
			fmt.Println("NUM+OP: MULTI PLUS")
			// same as multi plus?
		case "*":
			fmt.Println("NUM+OP: MULTI MULTI")
			i := -1
			if number.MatchString(subtree.Children[0].Value) {
				i = 0
			} else if number.MatchString(subtree.Children[1].Value) {
				i = 1
			}

			if i != -1 {
				val, err := strconv.ParseFloat(subtree.Children[i].Value, 64)
				if err != nil {
					return err
				}
				val *= n
				str := strconv.FormatFloat(val, 'f', -1, 64)
				subtree.Children[i].Value = str
				*t = *subtree
			}
		case "/", ":":
			fmt.Println("NUM+OP: MULTI DIV")
			i := -1
			if number.MatchString(subtree.Children[0].Value) {
				i = 0
			} else if number.MatchString(subtree.Children[1].Value) {
				i = 1
			}

			if i != -1 {
				val, err := strconv.ParseFloat(subtree.Children[i].Value, 64)
				if err != nil {
					return err
				}
				if i == 1 {
					val *= n
				} else {
					val /= n
				}
				str := strconv.FormatFloat(val, 'f', -1, 64)
				subtree.Children[i].Value = str
				*t = *subtree
			}
		} // switch subtree
	case "/", ":":
		fmt.Println("NUM+OP: DIVIDE")
		switch subtree.Value {
		case "+":
			fmt.Println("NUM+OP: DIV PLUS")
			// nothing to simplify
			// b+4/2 = b+2 (solved earlier)
		case "-":
			fmt.Println("NUM+OP: DIV MINUS")
			// nothing to simplify
			// b-4/2 = b-2 (solved earlier)
		case "/", ":":
			// 10/a/2 = 5/a (divide §1)
			//
			//	 	:
			//	 / \
			// 	2	  :
			//		/  \
			// 	 a	 10
			//
			//===================
			// a/10/2 = a/20 (multiply §2)
			//
			//	 	:
			//	 / \
			// 	2	  :
			//		/  \
			// 	 10	 a
			//
			//===================
			// TODO
			// 10/(a/2) = 20/a (multiply §3)
			//
			//	 	     :
			//	     /  \
			// 	    :		10
			//		/  \
			// 	 2	 a

			fmt.Println("NUM+OP: DIVIDE DIVIDE")
			i := -1
			if number.MatchString(subtree.Children[0].Value) {
				i = 0
			} else if number.MatchString(subtree.Children[1].Value) {
				i = 1
			}

			if i != -1 {
				val, err := strconv.ParseFloat(subtree.Children[i].Value, 64)
				if err != nil {
					return err
				}

				if i == 1 {
					// §1
					val /= n
				} else {
					// §2
					val *= n

					if t.Children[0].Value == "/" {
						subtree.Children[0].Value = subtree.Children[1].Value
						i = 1
					}
				}
				str := strconv.FormatFloat(val, 'f', -1, 64)
				subtree.Children[i].Value = str
				*t = *subtree
			}
		case "*":
			fmt.Println("NUM+OP: DIVIDE MULTI")
			i := -1
			if number.MatchString(subtree.Children[0].Value) {
				i = 0
			} else if number.MatchString(subtree.Children[1].Value) {
				i = 1
			}

			if i != -1 {
				val, err := strconv.ParseFloat(subtree.Children[i].Value, 64)
				if err != nil {
					return err
				}
				val /= n
				str := strconv.FormatFloat(val, 'f', -1, 64)
				subtree.Children[i].Value = str
				*t = *subtree
			}

		case "^":
			fmt.Println("NUM+OP: DIVIDE EXP")
			// nothin to simplify
			// b^4/3 still the same
		} // switch subtree
	default:
		return fmt.Errorf("Operation %v not yet implemented.", op)
	} // Switch op

	// append step if the tree changed
	if oldTree != t.InOrder() {
		*steps = append(*steps, fmt.Sprintf("%v = %v", oldTree, t.InOrder()))
	}

	return nil
}
