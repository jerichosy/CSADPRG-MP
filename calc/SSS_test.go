package calc

import "testing"

type sssTest struct {
	input float64
	expected  float64
}

var sssTests = []sssTest{
	sssTest{1, 180},
	sssTest{4249.99, 180},
	sssTest{4250, 202.5},
	sssTest{7000, 315},
	sssTest{8249.99, 360},
	sssTest{29750, 1350},
	sssTest{500000, 1350},
}

func TestSSS(t *testing.T) {
	for _, test := range sssTests {
		got := CalcSSS(test.input)
		if got != test.expected {
			t.Errorf("CalcSSS(%f) = %f; want %f", test.input, got, test.expected)
		}
	}
}