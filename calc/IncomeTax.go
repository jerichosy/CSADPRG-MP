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

// https://www.bir.gov.ph/images/bir_files/internal_communications_2/RMCs/2018/WT%20table.pdf
func CalcIncomeTax(taxableIncome float64) float64 {
	switch {
	case taxableIncome <= 20833:
		return 0
	case taxableIncome <= 33332:
		return (taxableIncome - 20833) * 0.2
	case taxableIncome <= 66666:
		return (taxableIncome-33333)*0.25 + 2500
	case taxableIncome <= 166666:
		return (taxableIncome-66667)*0.3 + 10833.33
	case taxableIncome <= 666666:
		return (taxableIncome-166667)*0.32 + 40833.33
	default:
		return (taxableIncome-666667)*0.35 + 200833.33
	}
}
