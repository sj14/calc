package eval

import (
	"testing"

	"github.com/sj14/calc/shuntingyard"
)

func TestAST(t *testing.T) {

	/////////////////////////////////////////////////////////////////////////////
	// Number & Number Simple
	/////////////////////////////////////////////////////////////////////////////

	// Simple Addition
	solve("3+4", "7", t)

	// Simple Substraction
	solve("5-2", "3", t)

	// Simple Substraction To Minus
	solve("5-8", "-3", t)

	// Addition Substraction
	solve("3+5-7", "1", t)

	// Simple Multiplication
	solve("2*3", "6", t)

	// Simple Multiplication To Minus
	solve("-2*3", "-6", t)

	// Simple Multiplication To Minus 2
	solve("2*-3", "-6", t)

	// Simple Division
	solve("6/3", "2", t)

	// Simple Division To Minus
	solve("-6/3", "-2", t)

	// Simple Division To Minus 2
	solve("6/-3", "-2", t)

	// Simple Exponentiate
	solve("2^3", "8", t)

	/////////////////////////////////////////////////////////////////////////////
	// Variable Variable
	/////////////////////////////////////////////////////////////////////////////

	// Simple Addition
	solve("a+a", "2*a", t)

	// Simple Substraction
	solve("a-a", "0", t)

	// Simple Substraction To Minus
	solve("-a-a", "-2*a", t)

	// Simple Multiplication
	solve("a*a", "a^2", t)

	// Multi to Minus
	// TODO remove 0
	solve("a*-a", "0-a^2", t)

	// Simple Division
	solve("a/a", "1", t)

	/////////////////////////////////////////////////////////////////////////////
	// Variable Operation
	/////////////////////////////////////////////////////////////////////////////

	// Addition Addition
	solve("a+5+a", "2*a+5", t)

	// Addition Addition
	solve("a+a+a", "3*a", t)

	// Addition Substraction
	solve("a+5-a", "5", t)

	// Addition Substraction 1
	solve("a+a-a", "a", t)

	// Addition Substraction 2
	solve("a-4+5", "a+1", t)

	// Addition Substraction 3
	solve("a-5+4", "a-1", t)

	// Addition Multiplication
	solve("a+5*a", "6*a", t)

	// Addition Multiplication
	solve("a+a*a", "a+a^2", t)

	// Addition Division
	// nothing to simplify
	solve("a+5/a", "a+5/a", t)

	// Addition Division
	solve("a+a/a", "a+1", t)

	// Addition Exponentiate
	// nothing to simplify
	solve("a+5^a", "a+5^a", t)

	// Addition Exponentiate
	// nothing to simplify
	solve("a+a^a", "a+a^a", t)

	// Substraction Substraction
	// TODO remove 0
	solve("a-5-a", "0-5", t)

	// Substraction Substraction
	// TODO remove 0
	solve("a-a-a", "0-a", t)

	// Substraction Addition
	solve("a-5+a", "2*a-5", t)

	// Substraction Addition
	// TODO remove 0
	solve("a-a+a", "0+a", t)

	// Substraction Multiplication
	solve("a-5*a", "-4*a", t)

	// Substraction Multiplication
	solve("a-a*a", "a-a^2", t)

	// Substraction Division
	// nothing to simplify
	solve("a-5/a", "a-5/a", t)

	// Substraction Division
	solve("a-a/a", "a-1", t)

	// Substraction Exponentiate
	// nothing to simplify
	solve("a-5^a", "a-5^a", t)

	// Substraction Exponentiate
	// nothing to simplify
	solve("a-a^a", "a-a^a", t)

	// Multiplication Addition
	solve("a*5+a", "a*6", t)

	// Multiplication Addition Multiplication
	// TODO remove 0
	solve("a*5+a*3", "a*8+0", t)

	// Multiplication Addition
	solve("a*a+a", "a^2+a", t)

	// Multiplication Substraction
	solve("a*5-a", "a*4", t)

	// Multiplication Substraction Multiplication
	// TODO remove 0
	solve("a*5-a*3", "a*2-0", t)
	solve("a*5-b*3", "a*5-b*3", t)

	// TODO
	//solve("a*b+a*b", "2*a*b", t)

	// Multiplication Substraction
	solve("a*a-a", "a^2-a", t)

	// Multiplication Multiplication
	solve("a*5*a", "a^2*5", t)

	// Multiplication Multiplication
	solve("a*a*a", "a^3", t)

	// Multiplication Division
	solve("a*5/a", "5", t)

	// Multiplication Division
	solve("a*a/a", "a", t)

	// Multiplication Exponentiate
	// nothing to simplify
	solve("a*5^a", "a*5^a", t)

	// Multiplication Exponentiate
	solve("a*a^a", "a*a^a", t)

	// Division Addition
	solve("a/5+a", "a/5+a", t)

	// Division Addition
	solve("a/a+a", "1+a", t)

	// Division Substraction
	// TODO or -(4*a)/5
	solve("a/5-a", "a/5-a", t)

	// Division Substraction
	solve("a/a-a", "1-a", t)

	// Division Multiplication
	solve("a/5*a", "a^2/5", t)

	// Division Multiplication
	// TODO remove 1*
	solve("a/a*a", "1*a", t)

	// Division Division
	solve("a/5/a", "1/5", t)

	// Division Division
	solve("a/a/a", "1/a", t)

	// Division Exponentiate
	// TODO or 5^(-a)*a
	solve("a/5^a", "a/5^a", t)

	// Division Exponentiate
	// TODO or a^(1-a)
	solve("a/a^a", "a/a^a", t)

	solve("a-b-a+a-a", "0-b", t)

	/////////////////////////////////////////////////////////////////////////////
	// Number Operation
	/////////////////////////////////////////////////////////////////////////////

	// Addition Addition
	solve("5+a+5", "10+a", t)

	// Addition Substraction 1
	// TODO remove 0
	solve("5+a-5", "0+a", t)

	// Addition Substraction 2
	solve("5+a-4", "1+a", t)

	// Addition Substraction 3
	solve("a+4-5", "a-1", t)

	// Addition Substraction 3,5
	solve("a+5-4", "a+1", t)

	// Addition Substraction 3
	solve("4+a-5", "-1+a", t)

	// Addition Multiplication
	// nothing to simplify
	solve("5+a*2", "5+a*2", t)

	// Addition Division
	// nothing to simplify
	solve("5+a/2", "5+a/2", t)

	// Addition Exponentiate
	// nothing to simplify
	solve("5+a^2", "5+a^2", t)

	// Substraction Addition
	solve("5-a+2", "7-a", t)

	solve("5-a-2", "3-a", t)

	// Substraction Substraction 1
	solve("a-4-3", "a-7", t)

	// Substraction Substraction 2
	solve("a-4-5", "a-9", t)

	// Substraction Multiplication
	// nothing to simplify
	solve("5-a*2", "5-a*2", t)

	// Substraction Division
	// nothing to simplify
	solve("5-a/2", "5-a/2", t)

	// Substraction Exponentiate
	// nothing to simplify
	solve("5+a^2", "5+a^2", t)

	// Multiplication Addition
	// nothing to simplify
	solve("5*a+2", "5*a+2", t)

	// Multiplication Substraction
	// nothing to simplify
	solve("5*a-2", "5*a-2", t)

	// Multiplication Multiplication
	solve("5*a*2", "10*a", t)

	// Multiplication Division
	solve("5*a/2", "2.5*a", t)

	// Multiplication Exponentiate
	// nothing to simplify
	solve("5*a^2", "5*a^2", t)

	// Division Addition
	// nothing to simplify
	solve("5/a+2", "5/a+2", t)

	// Division Substraction
	// nothing to simplify
	solve("5/a-2", "5/a-2", t)

	// Division Multiplication 1
	solve("5/a*2", "10/a", t)

	// Division Multiplication 2
	solve("4/(a*2)", "a*0.5", t)

	// Division Multiplication 3
	solve("a/6*2", "a/3", t)

	// Division Division 1
	solve("10/a/2", "5/a", t)

	// Division Division 1
	solve("a/10/2", "a/20", t)

	// Division Division 2
	solve("5/(a/2)", "10/a", t)
	solve("5/(a/10)", "50/a", t)

	// Division Exponentiate
	// nothing to simplify
	solve("5/a^2", "5/a^2", t)

	/////////////////////////////////////////////////////////////////////////////
	// Operation And Operation
	/////////////////////////////////////////////////////////////////////////////

	//	PLUS/MINUS
	//	  /		|
	// MULTI	 MULTI
	solve("a*5+c*3+a*2", "a*7+c*3+0", t)
	solve("a*5-c*3+a*2", "a*7-c*3+0", t)
	solve("a*5+c*3-a*2", "a*3+c*3-0", t)
	solve("a*5-c*3-a*2", "a*3-c*3-0", t)
	solve("a*5+c*3-a*10", "a*-5+c*3-0", t)

	//	PLUS/MINUS
	//	  /		|
	// DIV	 DIV
	// var = numerator
	solve("a/6-a/2", "-1*a/3", t)
	solve("a/6+a/2", "1*a/1.5", t)
	solve("a/2-a/6", "1*a/3", t)
	solve("a/2+a/6", "1*a/1.5", t)
	solve("a/19-a/12", "-7*a/228", t)
	solve("a/19+a/12", "31*a/228", t)
	solve("a/12-a/19", "7*a/228", t)
	solve("a/12+a/19", "31*a/228", t)

	// var = denominator
	solve("6/a-2/a", "4/a", t)
	solve("6/a+2/a", "8/a", t)

	//	PLUS/MINUS
	//	  /		|
	//  DIV/Multi
	solve("a*2-a/6", "11*a/6", t)
	solve("a*12-a/19", "227*a/19", t)

	// TODO
	// solve("a*-12-a/19", "-229*a/19", t)
	// solve("-6/a-2/a", "-8/a", t)

	solve("a*12+a/19", "229*a/19", t)
	solve("a/19-a*12", "-227*a/19", t)
	solve("a/19+a*12", "229*a/19", t)

	/////////////////////////////////////////////////////////////////////////////
	// Functions
	/////////////////////////////////////////////////////////////////////////////

	// Sqrt
	solve("sqrt(9)", "3", t)

	// Min
	solve("min(5,0,-5)", "-5", t)

	// Max
	solve("max(-5,4,0)", "4", t)
}

func solve(solve, want string, t *testing.T) {
	tree, err := shuntingyard.InfixToAST(solve)
	result, _, err := AST(&tree)
	if err != nil {
		t.Errorf("%v: Error not nil: %v", solve, err)
	}
	if result != want {
		t.Errorf("Failed %v. want: %v, got: %v", solve, want, result)
	}
}
