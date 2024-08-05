package e

import (
	"math"
	"testing"
)

var tcs = []struct {
	a         float64
	b         float64
	addExpect float64
	subExpect float64
	mulExpect float64
	divExpect float64
}{
	{0, 0, 0, 0, 0, math.NaN()},
	{1, 0, 1, 1, 0, math.Inf(1)},
	{0, 1, 1, -1, 0, 0},
	{2, 5, 7, -3, 10, 0.4},
	{5, 2, 7, 3, 10, 2.5},
}

func compare(a, b float64) bool {
	aNaN := math.IsNaN(a)
	bNaN := math.IsNaN(b)

	if aNaN && bNaN {
		return true
	} else if aNaN || bNaN {
		return false
	}

	if math.IsInf(a, -1) && math.IsInf(b, -1) {
		return true
	}
	if math.IsInf(a, 1) && math.IsInf(b, 1) {
		return true
	}

	const epsilon = 0.00001
	res := math.Abs(a - b)
	return res < epsilon
}

func TestAdd(t *testing.T) {
	for _, tc := range tcs {
		res := Add(tc.a, tc.b)
		expect := tc.addExpect

		if !compare(res, expect) {
			t.Errorf("Add(%f, %f) != %f", tc.a, tc.b, expect)
		}
	}
}

func TestSubtract(t *testing.T) {
	for _, tc := range tcs {
		res := Subtract(tc.a, tc.b)
		expect := tc.subExpect

		if !compare(res, expect) {
			t.Errorf("Subtract(%f, %f) != %f", tc.a, tc.b, expect)
		}
	}
}

func TestMultiply(t *testing.T) {
	for _, tc := range tcs {
		res := Multiply(tc.a, tc.b)
		expect := tc.mulExpect

		if !compare(res, expect) {
			t.Errorf("Multiply(%f, %f) != %f", tc.a, tc.b, expect)
		}
	}
}

func TestDivide(t *testing.T) {
	for _, tc := range tcs {
		res := Divide(tc.a, tc.b)
		expect := tc.divExpect

		if !compare(res, expect) {
			t.Errorf("Divide(%f, %f) != %f", tc.a, tc.b, expect)
		}
	}
}
