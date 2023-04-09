// --------------------
// Group: S11 - Group 7
// Members:
//  - LADA, MARTIN JOSE MERCHAN
//  - LIM, JUSTIN NATHANIEL UY
//  - MERTO, ALLEN NEO MASANGKAY
//  - SY, MATTHEW JERICHO GO
// Language: Go
// Paradigm: Procedural
// --------------------

package calc

import "testing"

type IncomeTaxTest struct {
	input    float64
	expected float64
}

var incomeTaxTests = []IncomeTaxTest{
	{0, 0.0},
	{10000, 0.0},
	{20833, 0.0},
	{30000, 1833.4},
	{33332, 2499.8},
	{33333, 2500.0},
	{50000, 6666.75},
	{66666, 10833.25},
	{100000, 20833.23},
	{166666, 40833.03},
	{500000, 147499.89},
	{666666, 200833.01},
	{1000000, 317499.88},
}

func TestIncomeTax(t *testing.T) {
	for _, test := range incomeTaxTests {
		got := CalcIncomeTax(test.input)
		if got != test.expected {
			t.Errorf("CalcIncomeTax(%f) = %f; want %f", test.input, got, test.expected)
		}
	}
}
