package validate

import (
	"fmt"

	"github.com/adrianosiq/teste-backends/golang/infra"
)

//WarrantySumTwiceLoanAmountValidation is responsible for validation sum warranties is greater or equal than twice loan amount
func WarrantySumTwiceLoanAmountValidation(sumWarranty float64, loanAmount float64) infra.Validation {
	loanAmount *= 2
	if sumWarranty < loanAmount {
		return infra.Validation{Err: fmt.Errorf("sum of warranties is less than double the loan amount")}
	}
	return infra.Validation{Err: nil}
}
