package operation

import "fmt"

// Operation of mathematical kind
// e.g. operators or functions
type Operation struct {
	Precedence       uint8
	RightAssociative bool
	Operands         uint8 // zero means any number of operands/inputs
}

var operators = map[string]Operation{
	"+": Operation{1, false, 2},
	"-": Operation{1, false, 2},
	"*": Operation{2, false, 2},
	"/": Operation{2, false, 2},
	"^": Operation{3, true, 2},
}

var functions = map[string]Operation{
	"sqrt": Operation{3, true, 1},
	"max":  Operation{3, true, 0},
	"min":  Operation{3, true, 0},
}

// Get returns the Operator or Function.
func Get(s string) (Operation, error) {
	if o, err := GetOperator(s); err == nil {
		return o, nil
	}
	if o, err := GetFunction(s); err == nil {
		return o, nil
	}
	return Operation{0, false, 0}, fmt.Errorf("Could not determine %v", s)
}

// GetOperator returns the Operator matching input s
func GetOperator(s string) (Operation, error) {
	o := operators[s]
	if o.Precedence == 0 {
		return Operation{}, fmt.Errorf("Operation %v not found.", s)
	}
	return o, nil
}

// GetFunction returns the Operator matching input s
func GetFunction(s string) (Operation, error) {
	f := functions[s]
	if f.Precedence == 0 {
		return Operation{}, fmt.Errorf("Function %v not found.", s)
	}
	return f, nil
}

// IsOperation returns true if input is an Operator or a Function.
func IsOperation(s string) bool {
	_, err := Get(s)
	if err == nil {
		return true
	}
	return false
}

// IsOperator returns true if input is an operator.
func IsOperator(s string) bool {
	_, err := GetOperator(s)
	if err == nil {
		return true
	}
	return false
}

// IsFunction returns true if input is a function.
func IsFunction(s string) bool {
	_, err := GetFunction(s)
	if err == nil {
		return true
	}
	return false
}
