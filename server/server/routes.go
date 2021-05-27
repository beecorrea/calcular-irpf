package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"txc/m/calculator"
)

type CalcDTO struct {
	Salary float64;
	Dependents int;
	Others float64;
}
// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
func CalcSalaryHandler(res http.ResponseWriter, req *http.Request) {
	var payload CalcDTO
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// 
	salary, dependents, others := payload.Salary, payload.Dependents, payload.Others
	newSalary := calculator.LiquidSalary(salary, others, uint64(dependents))
	// Convert to byte array to be able to send response.
	salaryString := fmt.Sprintf("%f", newSalary);
	salaryBytes := []byte(salaryString);

	res.Write(salaryBytes);
}

	

