package calc

// https://www.philhealth.gov.ph/news/2019/new_contri.php#gsc.tab=0
// https://www.philhealth.gov.ph/advisories/2022/adv2022-0010.pdf
func CalcPhilHealth(income float64) float64 {
	var result float64 = 0

	if income <= 10000 {
		result = 200.00
	} else if income > 10000 && income < 80000 {
		result = income * 0.02
	} else if income >= 80000 {
		result = 3200.00
	}

	return result
}
