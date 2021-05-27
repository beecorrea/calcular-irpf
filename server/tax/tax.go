package tax

// Source for values and ranges: https://www.dicionariofinanceiro.com/calculadora-de-salario-liquido/
import (
	"math"
)

func GetINSSPerc(salary float64) float64 {
	switch {
	case salary <= 1100.00:
		return 0.075
	case salary <= 2203.48:
		return 0.09
	case salary <= 3305.22:
		return 0.12
	case salary <= 6433.57:
		return 0.14
	default:
		return 0.00
	}
}
func GetINSSMaxContrib(salary float64) float64 {
	switch {
	case salary == 1100.00:
		return 82.50
	case salary == 2203.48:
		return 181.81
	case salary == 3305.22:
		return 314.01
	case salary >= 6433.57:
		return 751.97
	default:
		return 0.00
	}
}

func GetINSSDeduc(salary float64) float64 {
	switch {
	case salary <= 1100.00:
		return 0.00
	case salary <= 2203.48:
		return 16.50
	case salary <= 3305.22:
		return 82.61
	case salary >= 3305.23:
		return 148.72
	default:
		return 0.00
	}
}

func GetINSS(salary float64) float64 {
		// Floor to 2 decimal places
		var percentage float64 = GetINSSPerc(salary)
		var deduction float64 = GetINSSDeduc(salary)

		var totalDeduction float64
		if maxContrib := GetINSSMaxContrib(salary); maxContrib != 0 {
			// If salary is at the range ceiling, use the max contribution (flat value).
			totalDeduction = maxContrib - deduction
		} else {
			// Else, use percentage.
			totalDeduction = (salary * percentage) - deduction
		}
		return math.Floor(totalDeduction*100)/100
}

func GetIRRFPerc(salary float64) float64 {
	switch {
	case salary < 1903.98:
		return 0
	case salary <= 2826.65:
		return 0.075
	case salary <= 3751.05:
		return 0.15
	case salary <= 4664.68:
		return 0.225
	case salary > 4664.68:
		return 0.275
	default:
		return 0.00
	}
}

func GetIRRFDeduc(salary float64) float64 {
	switch {
	case salary < 1903.98:
		return 0
	case salary <= 2826.65:
		return 142.80
	case salary <= 3751.05:
		return 354.80
	case salary <= 4664.68:
		return 636.13
	case salary > 4664.68:
		return 869.36
	default:
		return 0.00
	}
}

func GetIRRF(salary float64) float64 {
	var percentual float64 = GetIRRFPerc(salary)
	var deducted float64 = GetIRRFDeduc(salary)
	var totalDeduction = (salary * percentual) - deducted
	// Floor to 2 decimal places
	return math.Floor(totalDeduction*100)/100
}

func GetIRRFDependents(dependents uint64) float64 {
	return 189.59 * float64(dependents);
}