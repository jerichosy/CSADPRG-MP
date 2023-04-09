package calc

// https://www.pagibigfund.gov.ph/document/pdf/circulars/provident/HDMF%20Circular%20No.%20274%20-%20Revised%20Guidelines%20on%20Pag-IBIG%20Fund%20Membership.pdf
// https://filipiknow.net/pag-ibig-contribution-table/#1-pag-ibig-contribution-table-in-2023-for-employees-and-employers
func CalcPagIBIG(monthlyCompensation float64) float64 {
	var result float64

	// "The maximum monthly compensation to be used in computing the employee and
	// employer contributions shall not be more than P5,000."
	if monthlyCompensation > 5000 {
		monthlyCompensation = 5000
	}

	// These are only employee contributions. Employer contributions are 2% regardless of bracket.
	if monthlyCompensation <= 1500 {
		result = monthlyCompensation * 0.01
	} else if monthlyCompensation > 1500 {
		result = monthlyCompensation * 0.02
	}

	return result
}
