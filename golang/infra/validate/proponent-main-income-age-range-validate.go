package validate

import (
	"fmt"

	"github.com/adrianosiq/teste-backends/golang/infra"
)

func getAgeRange(age int16) (int, error) {
	switch condition := age; {
	case condition >= 18 && condition < 24:
		return 4, nil
	case condition >= 24 && condition <= 50:
		return 3, nil
	case condition > 50:
		return 2, nil
	}
	return 0, fmt.Errorf("age range not found for %d", age)
}

//MainIncomeAgeRange is responsible for params the ProponentMainIncomeAgeRangeValidation
type MainIncomeAgeRange struct {
	MainAge             int16
	MonthlyIncome       float64
	LoanValue           float64
	MonthlyInstallments int32
}

//ProponentMainIncomeAgeRangeValidation is responsible for validating the proponent's income range
func ProponentMainIncomeAgeRangeValidation(mainIncomeAgeRange MainIncomeAgeRange) infra.Validation {
	multiplier, err := getAgeRange(mainIncomeAgeRange.MainAge)
	if err != nil {
		return infra.Validation{Err: err}
	}
	installmentValue := mainIncomeAgeRange.LoanValue / float64(mainIncomeAgeRange.MonthlyInstallments)
	if mainIncomeAgeRange.MonthlyIncome < (installmentValue * float64(multiplier)) {
		return infra.Validation{Err: fmt.Errorf("monthly income is less than %d times the installment of the loan", multiplier)}
	}
	return infra.Validation{Err: nil}
}
