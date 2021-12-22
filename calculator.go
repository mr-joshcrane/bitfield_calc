// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64, c ...float64) float64 {
	result := a + b
	for _, v := range c {
		result += v
	}
	return result
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64, c ...float64) float64 {
	result := a - b
	for _, v := range c {
		result -= v
	}
	return result
}

// Subtract takes two numbers and returns the result of multiplying them.
func Multiply(a, b float64, c ...float64) float64 {
	result := a * b
	for _, v := range c {
		result *= v
	}
	return result
}

func Divide(a, b float64, c ...float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad input %b: division by 0 not possible", b)
	}
	result := a / b
	for _, v := range c {
		if v == 0 {
			return 0, fmt.Errorf("bad input %b: division by 0 not possible", c)
		}
		result /= v
	}
	return result, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input %f: negative numbers do not have square roots", a)
	}
	return math.Sqrt(a), nil
}
