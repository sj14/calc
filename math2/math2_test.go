package math2

import "testing"

func TestGCD(t *testing.T) {
	solveGCD(12, 18, 6, t)
	solveGCD(12, 19, 1, t)
}

func TestLCM(t *testing.T) {
	solveLCM(21, 6, 42, t)
	solveLCM(12, 19, 228, t)
	solveLCM(6, 2, 6, t)

}

func solveGCD(a, b, want float64, t *testing.T) {
	result, err := GCD(a, b)
	if err != nil {
		t.Errorf("GCD(%v, %v): Error not nil: %v", a, b, err)
	}
	if result != want {
		t.Errorf("Failed GCD(%v, %v) want: %v, got: %v", a, b, want, result)
	}
}

func solveLCM(a, b, want float64, t *testing.T) {
	result, err := LCM(a, b)
	if err != nil {
		t.Errorf("LCM(%v, %v): Error not nil: %v", a, b, err)
	}
	if result != want {
		t.Errorf("Failed GCD(%v, %v) want: %v, got: %v", a, b, want, result)
	}
}
