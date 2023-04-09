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

type PagIBIGTest struct {
	input    float64
	expected float64
}

var pagibigTests = []PagIBIGTest{
	{0, 0},
	{10, 0.1},
	{100, 1},
	{1000, 10},
	{1450, 14.5},
	{1499, 14.99},
	{1499.99, 14.9999},
	{1500, 15},
	{1500.01, 30.0002},
	{1501, 30.02},
	{1550, 31},
	{3000, 60},
	{4999.99, 99.9998},
	{5000, 100},
	{5000.01, 100},
	{5001, 100},
	{500000, 100},
}

func TestPagIBIG(t *testing.T) {
	for _, test := range pagibigTests {
		got := CalcPagIBIG(test.input)
		if got != test.expected {
			t.Errorf("CalcPagIBIG(%f) = %f; want %f", test.input, got, test.expected)
		}
	}
}
