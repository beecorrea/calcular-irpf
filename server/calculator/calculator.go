package calculator

import (
	"fmt"
	"math"
	"txc/m/tax"
)

func LiquidSalary(baseSalary float64, otherDiscounts float64, dependents uint64) float64 {
	inss := baseSalary - tax.GetINSS(baseSalary)
	irrf := inss - (tax.GetIRRF(inss) - tax.GetIRRFDependents(dependents))
	discounts := irrf - otherDiscounts
	liquid := math.Floor(discounts*100)/100
	fmt.Printf("Liquid salary is %f", liquid)
	return liquid
}