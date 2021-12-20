// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a float64, b ...float64) float64 {
	for _, v := range b {
		a += v
	}
	return a
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a float64, b ...float64) float64 {
	for _, v := range b {
		a -= v
	}
	return a
}

// Subtract takes two numbers and returns the result of multiplying them.
func Multiply(a float64, b ...float64) float64 {
	for _, v := range b {
		a *= v
	}
	return a
}

func Divide(a float64, b ...float64) (float64, error) {
	for _, v := range b {
		if v == 0 {
			return 0, fmt.Errorf("bad input %b: division by 0 not possible", b)
		}
		a /= v
	}

	return a, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input %f: negative numbers do not have square roots", a)
	}
	return math.Sqrt(a), nil
}
