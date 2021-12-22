package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	cases := []struct {
		description string
		a, b, want  float64
	}{
		{description: "adding small integers", a: 1, b: 1, want: 2},
		{description: "small floats", a: 1.4, b: 1.2, want: 2.6},
		{description: "negative integers", a: -1, b: -10, want: -11},
		{description: "adding zero to integer returns original", a: 134, b: 0, want: 134},
		{description: "negative integer plus positive integer", a: -12, b: 12, want: 0},
		{description: "large integers", a: 4566584654, b: 654665465132, want: 659232049786},
	}
	for _, tc := range cases {
		got := calculator.Add(tc.a, tc.b)
		if !closeEnough(tc.want, got, 0.000001) {
			t.Fatalf("Test case: %s: want %f, got %f", tc.description, tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	cases := []struct {
		description string
		a, b, want  float64
	}{
		{description: "subtracting small integers", a: 1, b: 1, want: 0},
		{description: "small floats", a: 1.4, b: 1.2, want: 0.2},
		{description: "negative integers", a: -1, b: -10, want: 9},
		{description: "subtracting zero to integer returns original", a: 134, b: 0, want: 134},
		{description: "negative integer minus positive integer", a: -12, b: 12, want: -24},
		{description: "large integers", a: 4566584654, b: 654665465132, want: -650098880478},
	}
	for _, tc := range cases {
		got := calculator.Subtract(tc.a, tc.b)
		if !closeEnough(tc.want, got, 0.000001) {
			t.Fatalf("Test case: %s: want %f, got %f", tc.description, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	cases := []struct {
		description string
		a, b, want  float64
	}{
		{description: "identity function", a: 1, b: 1, want: 1},
		{description: "small floats", a: 1.4, b: 1.2, want: 1.68},
		{description: "multiplying negative integers is positive", a: -1, b: -10, want: 10},
		{description: "multiplying by zero is 0", a: 134, b: 0, want: 0},
		{description: "negative integer minus positive integer", a: -12, b: 12, want: -144},
		{description: "large integers", a: 4566584654, b: 654665465132, want: 2989585266575563292672},
	}
	for _, tc := range cases {
		got := calculator.Multiply(tc.a, tc.b)
		if !closeEnough(tc.want, got, 0.000001) {
			t.Fatalf("Test case: %s: want %f, got %f", tc.description, tc.want, got)
		}
	}
}

func TestDivideValidInput(t *testing.T) {
	t.Parallel()
	cases := []struct {
		description string
		a, b, want  float64
	}{
		{description: "dividing by 1 returns a", a: 1, b: 1, want: 1},
		{description: "dividing small floats", a: 1.4, b: 1.2, want: 1.166667},
		{description: "dividing by fraction is the same as multiplying by it's denominator", a: 1, b: 0.5, want: 2},
		{description: "dividing two negative numbers returns a positive number", a: -12, b: -6, want: 2},
	}
	for _, tc := range cases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("divide threw on valid input %f with error %t!", tc.a, err)
		}
		if !closeEnough(tc.want, got, 0.000001) {
			t.Fatalf("Test case: %s: want %f, got %f", tc.description, tc.want, got)

		}
	}
}

func TestDivideByZeroIsInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Fatalf("expected divide to return an error, but did not!")
	}
}

func TestMultiInput(t *testing.T) {
	t.Parallel()
	add := calculator.Add(0, 1, 2, 3, 4, 5)
	if add != 15 {
		t.Fatalf("multi input for add broken, got %f", add)
	}
	subtract := calculator.Subtract(0, 1, 2, 3, 4, 5)
	if subtract != -15 {
		t.Fatalf("multi input for subtract broken")
	}
	multiply := calculator.Multiply(1, 1, 2, 3, 4, 5)
	if multiply != 120 {
		t.Fatalf("multi input for multiply broken")
	}
	divide, err := calculator.Divide(240, 1, 2, 3, 4, 5)
	if err != nil {
		t.Fatalf("should not error")
	}
	if divide != 2 {
		t.Fatalf("multi input for division broken")
	}
	_, err = calculator.Divide(240, 1, 2, 3, 4, 0)
	if err == nil {
		t.Fatalf("should have errored, division by 0 illegal")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	cases := []struct {
		description string
		a, want     float64
		errExpected bool
	}{
		{description: "square root of 1 is 1", a: 1, want: 1, errExpected: false},
		{description: "small nondecimal floats", a: 4, want: 2, errExpected: false},
		{description: "larger non decimal floats", a: 1024, want: 32, errExpected: false},
		{description: "nondecimal floats with non int sqrts", a: 33, want: 5.744563, errExpected: false},
		{description: "decimal floats with non int sqrts", a: 33.3, want: 5.770615, errExpected: false},
		{description: "square root of 0 is 0", a: 0, want: 0, errExpected: false},
		{description: "cannot square root a negative number", a: -12, want: 0, errExpected: true},
	}
	for _, tc := range cases {
		got, err := calculator.Sqrt(tc.a)
		if tc.errExpected {
			if err == nil {
				t.Fatalf("expected sqrt to return an error, but did not!")
			}
		} else {
			if !closeEnough(tc.want, got, 0.000001) {
				t.Fatalf("Test case: %s: want %f, got %f", tc.description, tc.want, got)
			}
		}
	}
}

func closeEnough(a, b, epsilon float64) bool {
	diff := math.Abs(a - b)
	return diff < epsilon
}
