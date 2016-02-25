package math2

import "fmt"

// Max returns the maximum value from the input.
func Max(input []float64) float64 {
	var max float64

	if len(input) > 0 {
		max = input[0]
	}

	for _, o := range input {
		if o > max {
			max = o
		}
	}
	return max
}

// Min returns the minimum value from the input.
func Min(input []float64) float64 {
	var min float64

	if len(input) > 0 {
		min = input[0]
	}

	for _, o := range input {
		if o < min {
			min = o
		}
	}
	return min
}

// LCM calculates the Least Common Multiple
func LCM(a, b float64) (float64, error) {
	gcd, err := GCD(a, b)
	if err != nil {
		return 0, err
	}

	lcm := (a * b) / gcd
	return lcm, nil
}

// GCD calculates the Greatest Common Divisor
func GCD(a, b float64) (float64, error) {
	fmt.Printf("a: %v b: %v\n", a, b)
	if a == 0 || a == b {
		return a, nil
	}
	if b == 0 {
		return b, nil
	}

	if a < 0 || b < 0 {
		return 0, fmt.Errorf("Could not find a gcd for %v and %v", a, b)
	}

	if a > b {
		return GCD(a-b, b)
	}
	return GCD(a, b-a)

}
