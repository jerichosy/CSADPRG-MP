package calc

import "testing"

type sssTest struct {
	input float64
	expected  float64
}

var sssTests = []sssTest{
	sssTest{999.99, 0},
	sssTest{1000, 135},
	sssTest{4249.99, 180},
	sssTest{4250, 202.5},
	sssTest{6250, 292.5},
	sssTest{7000, 315},
	sssTest{8249.99, 360},
	sssTest{24749.99, 1102.5},
	sssTest{24750, 1125},
	sssTest{500000, 1125},
}

func TestSSS(t *testing.T) {
	for _, test := range sssTests {
		got := CalcSSS(test.input)
		if got != test.expected {
			t.Errorf("CalcSSS(%f) = %f; want %f", test.input, got, test.expected)
		}
	}
}