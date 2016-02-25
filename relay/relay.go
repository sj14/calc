package relay

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/sj14/calc/eval"
	"github.com/sj14/calc/shuntingyard"
)

// ^ begin; $ end
// + at least once
// * none to any
var bin = regexp.MustCompile(`^[0-1]+$`)
var oct = regexp.MustCompile(`^0[0-7]+$`)
var hex = regexp.MustCompile(`^0[xX][0-9a-fA-F]+$`)

//var calculation = regexp.MustCompile(`[0-9]`)

// Relay decides what to calculate
func Relay(input string) (string, []string, error) {
	switch {
	case bin.MatchString(input):
		fmt.Println("binary")
		result, err := strconv.ParseUint(input, 2, 64)
		if err != nil {
			return "", []string{""}, err
		}
		output := strconv.FormatUint(result, 10)
		return output, []string{""}, nil

	case oct.MatchString(input):
		fmt.Println("octal")
		result, err := strconv.ParseUint(input, 0, 64)
		if err != nil {
			return "", []string{""}, err
		}
		output := strconv.FormatUint(result, 10)

		return output, []string{""}, nil

	case hex.MatchString(input):
		fmt.Println("hexa")
		result, err := strconv.ParseUint(input, 0, 64)
		if err != nil {
			return "", []string{""}, err
		}
		output := strconv.FormatUint(result, 10)

		return output, []string{""}, nil
	default:
		t, err := shuntingyard.InfixToAST(input)
		if err != nil {
			return "", []string{""}, err
		}
		result, steps, err := eval.AST(&t)
		if err != nil {
			return "", []string{""}, err
		}
		fmt.Println(steps)
		return result, steps, nil
	}
}
