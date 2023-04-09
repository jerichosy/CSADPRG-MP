package calc

import "testing"

type sssTest struct {
	input    float64
	expected float64
}

var sssTests = []sssTest{
	{999.99, 0},
	{1000, 135},
	{4249.99, 180},
	{4250, 202.5},
	{6250, 292.5},
	{7000, 315},
	{8249.99, 360},
	{24749.99, 1102.5},
	{24750, 1125},
	{500000, 1125},
}

func TestSSS(t *testing.T) {
	for _, test := range sssTests {
		got := CalcSSS(test.input)
		if got != test.expected {
			t.Errorf("CalcSSS(%f) = %f; want %f", test.input, got, test.expected)
		}
	}
}
