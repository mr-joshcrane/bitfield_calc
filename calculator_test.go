package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	type Case struct {
		a           float64
		b           float64
		want        float64
		description string
	}
	cases := []Case{
		{1, 1, 2, "adding small integers"},
		{1.4, 1.2, 2.6, "small floats"},
		{-1, -10, -11, "negative integers"},
		{134, 0, 134, "adding zero to integer returns original"},
		{-12, 12, 0, "negative integer plus positive integer"},
		{4566584654, 654665465132, 659232049786, "large integers"},
	}
	t.Parallel()
	for _, tc := range cases {
		got := round(calculator.Add(tc.a, tc.b), 3)
		if tc.want != got {
			t.Fatalf("Test case: %s: want %f, got %f", tc.description, tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	type Case struct {
		a           float64
		b           float64
		want        float64
		description string
	}
	cases := []Case{
		{1, 1, 0, "subtracting small integers"},
		{1.4, 1.2, 0.2, "small floats"},
		{-1, -10, 9, "negative integers"},
		{134, 0, 134, "subtracting zero to integer returns original"},
		{-12, 12, -24, "negative integer minus positive integer"},
		{4566584654, 654665465132, -650098880478, "large integers"},
	}
	t.Parallel()
	for _, tc := range cases {
		got := round(calculator.Subtract(tc.a, tc.b), 3)
		if tc.want != got {
			t.Fatalf("Test case: %s: want %f, got %f", tc.description, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	type Case struct {
		a           float64
		b           float64
		want        float64
		description string
	}
	cases := []Case{
		{1, 1, 1, "identity function"},
		{1.4, 1.2, 1.68, "small floats"},
		{-1, -10, 10, "multiplying negative integers is positive"},
		{134, 0, 0, "multiplying by zero is 0"},
		{-12, 12, -144, "negative integer minus positive integer"},
		{4566584654, 654665465132, 2989585266575563292672, "large integers"},
	}
	t.Parallel()
	for _, tc := range cases {
		got := round(calculator.Multiply(tc.a, tc.b), 3)
		if tc.want != got {
			t.Fatalf("Test case: %s: want %f, got %f", tc.description, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	type Case struct {
		a           float64
		b           float64
		want        float64
		errExpected bool
		description string
	}
	t.Parallel()
	cases := []Case{
		{1, 1, 1, false, "dividing by 1 returns a"},
		{1.4, 1.2, 1.167, false, "dividing small floats"},
		{1, 0.5, 2, false, "dividing by fraction is the same as multiplying by it's denominator"},
		{-12, -6, 2, false, "dividing two negative numbers returns a positive number"},
		{1, 0, -1, true, "division by zero is illegal"},
	}
	for _, tc := range cases {
		i, err := calculator.Divide(tc.a, tc.b)
		if tc.errExpected {
			if err == nil {
				t.Fatalf("expected divide to return an error, but did not!")
			}
		} else {
			got := round(i, 3)
			if tc.want != got {
				t.Fatalf("Test case: %s: want %f, got %f", tc.description, tc.want, got)
			}
		}
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
	type Case struct {
		a           float64
		want        float64
		errExpected bool
		description string
	}
	t.Parallel()
	cases := []Case{
		{1, 1, false, "square root of 1 is 1"},
		{4, 2, false, "small nondecimal floats"},
		{1024, 32, false, "larger non decimal floats"},
		{33, 5.745, false, "nondecimal floats with non int sqrts"},
		{33.3, 5.771, false, "decimal floats with non int sqrts"},
		{0, 0, false, "square root of 0 is 0"},
		{-12, 0, true, "cannot square root a negative number"},
	}
	for _, tc := range cases {
		i, err := calculator.Sqrt(tc.a)
		if tc.errExpected {
			if err == nil {
				t.Fatalf("expected sqrt to return an error, but did not!")
			}
		} else {
			got := round(i, 3)
			if tc.want != got {
				t.Fatalf("Test case: %s: want %f, got %f", tc.description, tc.want, got)
			}
		}
	}
}

func round(i float64, precision int) float64 {
	sigFigures := float64(math.Pow(10, float64(precision)))
	rounded := math.Round(i*sigFigures) / sigFigures
	return rounded
}
