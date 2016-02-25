package relay

import (
	"testing"
)

func TestRelay(t *testing.T) {
	// Binary
	resultBin1, _, err := Relay("101010")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultBin1 != "42" {
		t.Errorf("Failed. want: 42, got: %v", resultBin1)
	}

	resultBin2, _, err := Relay("0101010")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultBin2 != "42" {
		t.Errorf("Failed. want: 42, got: %v", resultBin2)
	}

	// Octal
	resultOct, _, err := Relay("052")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultOct != "42" {
		t.Errorf("Failed. want: 42, got: %v", resultOct)
	}

	// Hex
	resultHex, _, err := Relay("0xBEEF")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultHex != "48879" {
		t.Errorf("Failed. want: 48879, got: %v", resultHex)
	}

	// Calculate
	resultCalc1, _, err := Relay("2+3*4")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultCalc1 != "14" {
		t.Errorf("Failed. want: 14, got: %v", resultCalc1)
	}

	resultCalc2, _, err := Relay("3+4*2.5/(1-5)^2^3")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultCalc2 != "3.000152587890625" {
		t.Errorf("Failed. want: 3.000152587890625, got: %v", resultCalc2)
	}

	// Calculate with Functions
	resultFunc1, _, err := Relay("6+max(max(7,10),11)+3^2")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultFunc1 != "26" {
		t.Errorf("Failed. want: 26, got: %v", resultFunc1)
	}

	resultFunc2, _, err := Relay("sqrt(9)")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultFunc2 != "3" {
		t.Errorf("Failed. want: 3, got: %v", resultFunc2)
	}

	// Variables
	resultVar1, _, err := Relay("a+b+a+b+a+b")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar1 != "3*a+3*b" {
		t.Errorf("Failed. want: 3*a+3*b, got: %v", resultVar1)
	}

	resultVar2, _, err := Relay("a*a+b*a*a")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar2 != "a^2+b*a^2" {
		t.Errorf("Failed. want: a^2+b*a^2, got: %v", resultVar2)
	}

	resultVar3, _, err := Relay("a*a/a")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar3 != "a" {
		t.Errorf("Failed. want: a, got: %v", resultVar3)
	}

	resultVar4, _, err := Relay("b+4-5")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar4 != "b-1" {
		t.Errorf("Failed. want: b-1, got: %v", resultVar4)
	}

	resultVar5, _, err := Relay("2*a-a")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar5 != "a" {
		t.Errorf("Failed. want: a, got: %v", resultVar5)
	}

	// Div Div
	resultVar6, _, err := Relay("a/5/a")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar6 != "1/5" {
		t.Errorf("Failed. want: 1/5, got: %v", resultVar6)
	}

	// Multi Div
	resultVar7, _, err := Relay("a/b*a")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar7 != "a^2/b" {
		t.Errorf("Failed. want: a^2/b, got: %v", resultVar7)
	}

	// Multi Multi
	resultVar8, _, err := Relay("a*b*a")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar8 != "a^2*b" {
		t.Errorf("Failed. want: a^2*b, got: %v", resultVar8)
	}

	// Div Multi
	resultVar9, _, err := Relay("a*b/a")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar9 != "b" {
		t.Errorf("Failed. want: b, got: %v", resultVar9)
	}

	// Div Exp 1
	resultVar10, _, err := Relay("a^2/a")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar10 != "a" {
		t.Errorf("Failed. want: a, got: %v", resultVar10)
	}

	// Div Exp 2
	resultVar11, _, err := Relay("2^a/a")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar11 != "2^a/a" {
		t.Errorf("Failed. want: 2^a/a, got: %v", resultVar11)
	}

	// Div Exp 3
	resultVar12, _, err := Relay("a^3/a")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar12 != "a^2" {
		t.Errorf("Failed. want: a^2, got: %v", resultVar12)
	}

	// Plus Minus
	varOpPlusMinus, _, err := Relay("b+4-10+b")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if varOpPlusMinus != "2*b-6" {
		t.Errorf("Failed. want: 2*b-6, got: %v", varOpPlusMinus)
	}

	////////////////////////
	// Number And Operation
	////////////////////////

	// Plus Minus 1
	resultVar13, _, err := Relay("b+4-5+5")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar13 != "b+4" {
		t.Errorf("Failed. want: b+4, got: %v", resultVar13)
	}

	// Plus Minus 2
	resultVar14, _, err := Relay("b+4-10+5")
	if err != nil {
		t.Errorf("Failed. err not nil: %v", err)
	}
	if resultVar14 != "b-1" {
		t.Errorf("Failed. want: b-1, got: %v", resultVar14)
	}

}
