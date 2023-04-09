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

// 2022 PhilHealth Premium Rate
// https://www.philhealth.gov.ph/news/2019/new_contri.php#gsc.tab=0
// https://www.philhealth.gov.ph/advisories/2022/adv2022-0010.pdf
func CalcPhilHealth(income float64) float64 {
	var result float64 = 0

	if income <= 10000 {
		result = 400.00
	} else if income > 10000 && income < 80000 {
		result = income * 0.04
	} else if income >= 80000 {
		result = 3200.00
	}

	// Split between employer and employee
	result = result / 2

	return result
}
