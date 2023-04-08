package main

import (
    "fmt"
    "github.com/jerichosy/CSADPRG-MP/calc"
)

func main() {
	fmt.Println("--- PH Tax Calculator ---")

	var monthlyIncome float64
	fmt.Print("Enter monthly income: ")
	fmt.Scan(&monthlyIncome)

	// Get taxable income (minus SSS/GSIS, PhilHealth, Pag-IBIG)
	fmt.Println(calc.CalcSSS(monthlyIncome))
}