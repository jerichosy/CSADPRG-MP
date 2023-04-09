package main

import (
	"fmt"

	"github.com/jerichosy/CSADPRG-MP/tree/main/calc"
)

func main() {
	fmt.Println("--- PH Tax Calculator ---")

	var monthlyIncome float64

	var sssContribution float64
	var philhealthContribution float64
	var pagibigContribution float64

	var totalContributions float64

	var incomeTax float64

	var totalDeductions float64

	fmt.Print("Enter monthly income: ")
	fmt.Scan(&monthlyIncome)

	// Calculate SSS, PhilHealth, and Pag-IBIG from monthly income
	sssContribution = calc.CalcSSS(monthlyIncome)
	philhealthContribution = calc.CalcPhilHealth(monthlyIncome)
	// pagibigContribution = calc.CalcPagIBIG(monthlyIncome)

	totalContributions = sssContribution + philhealthContribution + pagibigContribution

	// Calculate income tax from taxable income (monthly income minus SSS, PhilHealth, and Pag-IBIG)
	incomeTax = calc.CalcIncomeTax(monthlyIncome - totalContributions)

	// Calculate total deductions (SSS, PhilHealth, Pag-IBIG, and income tax)
	totalDeductions = totalContributions + incomeTax

	// Display SSS, PhilHealth, Pag-IBIG contributions
	fmt.Println("\n- Monthly Contributions - ")
	fmt.Printf("SSS: %.2f\n", sssContribution)
	fmt.Printf("PhilHealth: %.2f\n", philhealthContribution)
	fmt.Printf("Pag-IBIG: %.2f\n", pagibigContribution)
	fmt.Printf("TOTAL CONTRIBUTIONS: %.2f\n", totalContributions)

	// Display income tax
	fmt.Println("\n- Tax Computation - ")
	fmt.Printf("Income Tax: %.2f\n", incomeTax)
	fmt.Printf("NET PAY AFTER TAX: %.2f\n", monthlyIncome-incomeTax)

	// Display total deductions and actual net pay
	fmt.Println("\n-------------------------")
	fmt.Printf("TOTAL DEDUCTIONS: %.2f\n", totalDeductions)
	fmt.Printf("NET PAY AFTER DEDUCTIONS: %.2f\n", monthlyIncome-totalDeductions)
}
