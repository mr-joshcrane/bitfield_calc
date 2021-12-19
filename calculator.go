// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a float64, b ...float64) float64 {
	for _, i := range b {
		a += i
	}
	return a
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a float64, b ...float64) float64 {
	for _, i := range b {
		a -= i
	}
	return a
}

// Subtract takes two numbers and returns the result of multiplying them.
func Multiply(a float64, b ...float64) float64 {
	for _, i := range b {
		a *= i
	}
	return a
}

func Divide(a float64, b ...float64) (float64, error) {
	for _, i := range b {
		if i == 0 {
			return math.NaN(), errors.New("division by 0 not possible")
		}
		a /= i
	}

	return a, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return math.NaN(), errors.New("negative numbers do not have square roots")
	}
	return math.Sqrt(a), nil
}
