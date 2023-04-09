// --------------------
// Group: S11 - Group 7
// Members:
//  - LADA, MARTIN JOSE MERCHAN
//  - LIM, JUSTIN NATHANIEL UY
//  - MERTO, ALLEN NEO MASANGKAY
//  - SY, MATTHEW JERICHO GO
// Language: Go
// --------------------

package calc

import "testing"

type PhilHealthTest struct {
	input    float64
	expected float64
}

var philHealthTests = []PhilHealthTest{
	{30000, 600},
	{1000, 200},
	{500000, 1600},
	{1337, 200},
	{4250, 200},
	{1984, 200},
	{33332, 666.64},
	{33333, 666.66},
	{9999.99, 200},
	{10000, 200},
	{10000.01, 200.0002},
	{79999.99, 1599.9998},
	{80000, 1600},
	{80000.01, 1600},
}

func TestPhilHealth(t *testing.T) {
	for _, test := range philHealthTests {
		got := CalcPhilHealth(test.input)
		if got != test.expected {
			t.Errorf("CalcPhilHealth(%f) = %f; want %f", test.input, got, test.expected)
		}
	}
}
