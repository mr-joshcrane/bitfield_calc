package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		description string
		a, b, want  float64
	}{
		{"adding small integers", 1, 1, 2},
		{"small floats", 1.4, 1.2, 2.6},
		{"negative integers", -1, -10, -11},
		{"adding zero to integer returns original", 134, 0, 134},
		{"negative integer plus positive integer", -12, 12, 0},
		{"large integers", 4566584654, 654665465132, 659232049786},
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
	cases := []struct {
		description string
		a, b, want  float64
	}{
		{"subtracting small integers", 1, 1, 0},
		{"small floats", 1.4, 1.2, 0.2},
		{"negative integers", -1, -10, 9},
		{"subtracting zero to integer returns original", 134, 0, 134},
		{"negative integer minus positive integer", -12, 12, -24},
		{"large integers", 4566584654, 654665465132, -650098880478},
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
	cases := []struct {
		description string
		a, b, want  float64
	}{
		{"identity function", 1, 1, 1},
		{"small floats", 1.4, 1.2, 1.68},
		{"multiplying negative integers is positive", -1, -10, 10},
		{"multiplying by zero is 0", 134, 0, 0},
		{"negative integer minus positive integer", -12, 12, -144},
		{"large integers", 4566584654, 654665465132, 2989585266575563292672},
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
	t.Parallel()
	cases := []struct {
		description string
		a, b, want  float64
		errExpected bool
	}{
		{"dividing by 1 returns a", 1, 1, 1, false},
		{"dividing small floats", 1.4, 1.2, 1.167, false},
		{"dividing by fraction is the same as multiplying by it's denominator", 1, 0.5, 2, false},
		{"dividing two negative numbers returns a positive number", -12, -6, 2, false},
		{"division by zero is illegal", 1, 0, -1, true},
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
	t.Parallel()
	cases := []struct {
		description string
		a, want     float64
		errExpected bool
	}{
		{"square root of 1 is 1", 1, 1, false},
		{"small nondecimal floats", 4, 2, false},
		{"larger non decimal floats", 1024, 32, false},
		{"nondecimal floats with non int sqrts", 33, 5.745, false},
		{"decimal floats with non int sqrts", 33.3, 5.771, false},
		{"square root of 0 is 0", 0, 0, false},
		{"cannot square root a negative number", -12, 0, true},
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
