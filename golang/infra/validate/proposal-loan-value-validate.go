package validate

import (
	"fmt"

	"github.com/adrianosiq/teste-backends/golang/infra"
)

const (
	minLoanValue float64 = 30000.0
	maxLoanValue float64 = 3000000.0
)

//ProposalLoanValueValidation is responsible for validating whether the loan amount is between the minimum and maximum configured
func ProposalLoanValueValidation(value float64) infra.Validation {
	if value < minLoanValue {
		return infra.Validation{Err: fmt.Errorf("proposal loan value is less than minimum of %g", minLoanValue)}
	}
	if value > maxLoanValue {
		return infra.Validation{Err: fmt.Errorf("proposal loan value is greater than maximum of %g", maxLoanValue)}
	}
	return infra.Validation{Err: nil}
}
