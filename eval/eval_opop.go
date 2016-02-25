package eval

import (
	"fmt"
	"strconv"

	"github.com/sj14/calc/math2"
	"github.com/sj14/calc/tree"
)

func operationAndOperation(fulltree, subtree0, subtree1 *tree.Tree, steps *[]string) error {
	for i := range subtree0.Children {
		err := operationAndOperation(fulltree, &subtree0.Children[i], subtree1, steps)
		if err != nil {
			return err
		}
	}
	for i := range subtree1.Children {
		err := operationAndOperation(fulltree, subtree0, &subtree1.Children[i], steps)
		if err != nil {
			return err
		}
	}
	oldTree := fulltree.InOrder()
	if (fulltree.Value == "+" || fulltree.Value == "-") && subtree0.Value == "*" && subtree1.Value == "*" {
		fmt.Printf("Full tree: %v\n", fulltree)
		//	PLUS/MINUS
		//	  /		|
		// 	 *		*
		i := -1
		if variable.MatchString(subtree0.Children[0].Value) && subtree0.Children[0].Value == subtree1.Children[0].Value {
			if number.MatchString(subtree0.Children[1].Value) && number.MatchString(subtree1.Children[1].Value) {
				i = 1
			}
		} else if variable.MatchString(subtree0.Children[1].Value) && subtree0.Children[1].Value == subtree1.Children[1].Value {
			if number.MatchString(subtree0.Children[0].Value) && number.MatchString(subtree1.Children[0].Value) {
				i = 0
			}
		}
		if i != -1 {
			// same variable in booth subtrees
			// merge booth sides by adding or substracting values
			num0, err := strconv.ParseFloat(subtree0.Children[i].Value, 64)
			if err != nil {
				return err
			}
			num1, err := strconv.ParseFloat(subtree1.Children[i].Value, 64)
			if err != nil {
				return err
			}
			var result float64
			if fulltree.Value == "+" {
				result = num1 + num0
			} else if fulltree.Value == "-" {
				result = num1 - num0 // order is important. difference between 5-3 and 3-5
			}
			resStr := strconv.FormatFloat(result, 'f', -1, 64)
			fmt.Printf("Subtree0: %v, Subtree1: %v\n", subtree0, subtree1)
			subtree1.Children[i].Value = resStr
			subtree0.Value = "0"
			subtree0.Children = nil
		}
	} else if (fulltree.Value == "+") && subtree0.Value == "/" && subtree1.Value == "/" {
		//		 PLUS
		//	  /		|
		// 	 DIV	DIV
		if variable.MatchString(subtree0.Children[0].Value) && subtree0.Children[0].Value == subtree1.Children[0].Value {
			if number.MatchString(subtree0.Children[1].Value) && number.MatchString(subtree1.Children[1].Value) {
				// variable is the denominator (dt: Nenner)
				// 6/a+2/a = 8/a
				num0, err := strconv.ParseFloat(subtree0.Children[1].Value, 64)
				if err != nil {
					return err
				}
				num1, err := strconv.ParseFloat(subtree1.Children[1].Value, 64)
				if err != nil {
					return err
				}
				result := num1 + num0 // order is important. difference between 5-3 and 3-5
				resStr := strconv.FormatFloat(result, 'f', -1, 64)
				subtree0.Children[1].Value = resStr
				*fulltree = *subtree0
			}
		} else if variable.MatchString(subtree0.Children[1].Value) && subtree0.Children[1].Value == subtree1.Children[1].Value {
			if number.MatchString(subtree0.Children[0].Value) && number.MatchString(subtree1.Children[0].Value) {
				// variable is the numerator (dt: Zähler)
				// a/6+a/2 = 1*a/1.5
				// a/12+a/19 = 31*a/228

				vari := subtree0.Children[1].Value
				num0, err := strconv.ParseFloat(subtree0.Children[0].Value, 64)
				if err != nil {
					return err
				}
				num1, err := strconv.ParseFloat(subtree1.Children[0].Value, 64)
				if err != nil {
					return err
				}

				lcm, err := math2.LCM(num0, num1)
				if err != nil {
					return err
				}
				var numerator, denominator float64
				if lcm == num0 {
					numerator = lcm / num1         // 6/2 = 3
					numerator++                    // = 4
					denominator = num0 / numerator // 6/4 = 1,5
					numerator = 1                  // a/6+a/2 = 1*a/1.5
				} else if lcm == num1 {
					numerator = lcm / num0
					denominator = num1
					numerator++
					denominator = num1 / numerator
					numerator = 1 // a/2+a/6 = 2*a/3 == 1*a/1.5
				} else {
					denominator = lcm
					numerator = num0 + num1
				}
				lcmStr := strconv.FormatFloat(denominator, 'f', -1, 64)
				resStr := strconv.FormatFloat(numerator, 'f', -1, 64)

				subtree0.Children = nil
				childValue := tree.Tree{Value: resStr}
				childVari := tree.Tree{Value: vari}
				childMulti := tree.Tree{Value: "*", Children: []tree.Tree{childVari, childValue}}
				childLCM := tree.Tree{Value: lcmStr}
				subtree0.Children = append(subtree0.Children, childLCM, childMulti)
				*fulltree = *subtree0
			}
		}
	} else if (fulltree.Value == "-") && subtree0.Value == "/" && subtree1.Value == "/" {
		//		MINUS
		//	  /		|
		// 	 DIV	DIV
		if variable.MatchString(subtree0.Children[0].Value) && subtree0.Children[0].Value == subtree1.Children[0].Value {
			if number.MatchString(subtree0.Children[1].Value) && number.MatchString(subtree1.Children[1].Value) {
				// variable is the denominator (dt: Nenner)
				// 6/a-2/a = 4/a
				num0, err := strconv.ParseFloat(subtree0.Children[1].Value, 64)
				if err != nil {
					return err
				}
				num1, err := strconv.ParseFloat(subtree1.Children[1].Value, 64)
				if err != nil {
					return err
				}
				result := num1 - num0 // order is important. difference between 5-3 and 3-5
				resStr := strconv.FormatFloat(result, 'f', -1, 64)
				subtree0.Children[1].Value = resStr
				*fulltree = *subtree0
			}
		} else if variable.MatchString(subtree0.Children[1].Value) && subtree0.Children[1].Value == subtree1.Children[1].Value {
			if number.MatchString(subtree0.Children[0].Value) && number.MatchString(subtree1.Children[0].Value) {
				// variable is the numerator (dt: Zähler)
				// a/6-a/2 = -1*a/3
				// a/3-a/6 = a/6
				// a/19-a/12 = -7*a/228

				vari := subtree0.Children[1].Value
				num0, err := strconv.ParseFloat(subtree0.Children[0].Value, 64)
				if err != nil {
					return err
				}
				num1, err := strconv.ParseFloat(subtree1.Children[0].Value, 64)
				if err != nil {
					return err
				}

				lcm, err := math2.LCM(num0, num1)
				if err != nil {
					return err
				}
				var numerator, denominator float64
				if lcm == num0 {
					numerator = lcm / num1         // 6/2 = 3
					numerator--                    // = 2
					denominator = num0 / numerator // 6/2 = 3
					numerator = 1                  // a/6-a/2 = 1*a/3
				} else if lcm == num1 {
					numerator = lcm / num0
					denominator = num1
					numerator--
					denominator = num1 / numerator
					numerator = -1 // a/2-a/6 = -1*a/3
				} else {
					denominator = lcm
					numerator = num0 - num1 // order is important. difference between 5-3 and 3-5
				}
				lcmStr := strconv.FormatFloat(denominator, 'f', -1, 64)
				resStr := strconv.FormatFloat(numerator, 'f', -1, 64)

				subtree0.Children = nil
				childValue := tree.Tree{Value: resStr}
				childVari := tree.Tree{Value: vari}
				childMulti := tree.Tree{Value: "*", Children: []tree.Tree{childVari, childValue}}
				childLCM := tree.Tree{Value: lcmStr}
				subtree0.Children = append(subtree0.Children, childLCM, childMulti)
				*fulltree = *subtree0
			}
		}
	} else if (fulltree.Value == "-" || fulltree.Value == "+") && ((subtree0.Value == "/" || subtree0.Value == "*") && (subtree1.Value == "/" || subtree1.Value == "*")) {
		fmt.Println("op - or + and / * or * /")
		//	PLUS/MINUS
		//	  /		|
		// 	DIV/MULTI
		//
		// a*6+a/2 = 13*a/2 <-> a/2+6*a = 13*a/2
		// a*6-a/2 = 11*a/2 <-> a/2-6*a = 11*a/2
		//
		// a*6+2/a <-> 6*a+2/a (same)
		// a*6-2/a <-> 6*a-2/a (same)

		if number.MatchString(subtree0.Children[0].Value) && number.MatchString(subtree1.Children[0].Value) &&
			variable.MatchString(subtree0.Children[1].Value) && subtree0.Children[1].Value == subtree1.Children[1].Value {
			fmt.Println("Variable is numerator")
			// a/19 + a*12 = 229*a/19
			// a = vari
			// 19 = denominator
			// 12 = multiplicator
			// 229 = lcm (+1)

			// a/6+a*2 = 13*a/6

			vari := subtree0.Children[1].Value

			var denominatorStr, multiplicatorStr string

			if subtree0.Value == "/" {
				denominatorStr = subtree0.Children[0].Value
				multiplicatorStr = subtree1.Children[0].Value
			} else {
				denominatorStr = subtree1.Children[0].Value
				multiplicatorStr = subtree0.Children[0].Value
			}

			denominator, err := strconv.ParseFloat(denominatorStr, 64)
			if err != nil {
				return err
			}
			multiplicator, err := strconv.ParseFloat(multiplicatorStr, 64)
			if err != nil {
				return err
			}

			lcm, err := math2.LCM(denominator, multiplicator)
			if err != nil {
				return err
			}

			var numerator float64
			if lcm == denominator {
				numerator = lcm * multiplicator
			} else if lcm == multiplicator {
				numerator = lcm * denominator
			} else {
				numerator = lcm
			}

			if fulltree.Value == "+" {
				numerator++
			} else {
				numerator--

				if subtree1.Value == "/" {
					// 	a/19 - a*12 = -227*a/19 §1
					// 	a*12 - a/19 =  227*a/19
					numerator *= -1 // §1
				}
			}
			numeratorStr := strconv.FormatFloat(numerator, 'f', -1, 64)
			denominatorStr = strconv.FormatFloat(denominator, 'f', -1, 64)

			// Create new tree nodes
			childNumerator := tree.Tree{Value: numeratorStr}
			childdenominator := tree.Tree{Value: denominatorStr}
			childVari := tree.Tree{Value: vari}
			childMulti := tree.Tree{Value: "*", Children: []tree.Tree{childVari, childNumerator}}
			divTree := tree.Tree{Value: "/", Children: []tree.Tree{childdenominator, childMulti}}
			*fulltree = divTree
		}
	}
	// append step if the tree changed
	if oldTree != fulltree.InOrder() {
		*steps = append(*steps, fmt.Sprintf("%v = %v", oldTree, fulltree.InOrder()))
	}

	return nil
}
