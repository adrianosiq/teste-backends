package validate

import (
	"fmt"

	"github.com/adrianosiq/teste-backends/golang/infra"
)

const (
	minNumberMonthlyInstallments = 24
	maxNumberMonthlyInstallments = 180
)

//ProposalNumberOfMonthlyInstallmentsValidation is responsible for validating whether the number of monthly installments is between the minimum and maximum configured
func ProposalNumberOfMonthlyInstallmentsValidation(number int32) infra.Validation {
	if number < minNumberMonthlyInstallments {
		return infra.Validation{Err: fmt.Errorf("proposal number of monthly installments is less than minimum of %d", minNumberMonthlyInstallments)}
	}
	if number > maxNumberMonthlyInstallments {
		return infra.Validation{Err: fmt.Errorf("proposal number of monthly installments is greater than maximum of %d", maxNumberMonthlyInstallments)}
	}
	return infra.Validation{Err: nil}
}
