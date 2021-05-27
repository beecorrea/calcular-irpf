package tax

import (
	"testing"
)


func TestINSSPerc(t *testing.T){
	salary := 4800.00
	perc := GetINSSPerc(salary);

	if perc != 0.14 {
		t.Fatalf("Percentual is %f, expected %f", perc, 0.14);
	}
}

func TestINSSMaxContrib(t *testing.T) {
	salary := 4800.00
	maxContrib := GetINSSMaxContrib(salary);

	if maxContrib != 0.00 {
		t.Fatalf("Max contribution is %f, expected %f", maxContrib, 0.00);
	}
}

func TestINSSDeduc(t *testing.T) {
	salary := 4800.00
	deduction := GetINSSDeduc(salary);

	if deduction != 148.72 {
		t.Fatalf("Deduction is %f, expected %f", deduction, 148.72);
	}
}
func TestINSS(t *testing.T) {
	salary := 4800.00
	tax := GetINSS(salary);

	if tax != 523.28{
		t.Fatalf("INSS Tax is %f, expected %f", tax, 523.27);
	}
}

func TestIRRFPerc(t *testing.T){
	baseSalary := 4276.82
	tax := GetIRRFPerc(baseSalary);

	if tax != 0.225{
		t.Fatalf("Percentual is %f, expected %f", tax, 0.225);
	}
}

func TestIRFFDeduc(t *testing.T) {
	baseSalary := 4276.82
	tax := GetIRRFDeduc(baseSalary)

	if tax != 636.13 {
		t.Fatalf("Deduction is %f, expected %f", tax, 636.13);
	}
}

func TestIRRF(t *testing.T) {
	baseSalary := 4276.82
	totalTax := GetIRRF(baseSalary)
	if totalTax != 326.15 {
		t.Fatalf("Tax is %f, expected %f", totalTax, 326.15);
	}
}
