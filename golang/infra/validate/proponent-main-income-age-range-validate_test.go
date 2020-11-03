package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProponentMainIncomeAgeRangeValidation(t *testing.T) {
	t.Run("Should return an error if age range not found", func(t *testing.T) {
		mainIncomeAgeRange := MainIncomeAgeRange{
			MainAge:             17,
			MonthlyIncome:       10000.0,
			LoanValue:           30000.0,
			MonthlyInstallments: 72,
		}
		out := ProponentMainIncomeAgeRangeValidation(mainIncomeAgeRange)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("age range not found for %d", mainIncomeAgeRange.MainAge), out.Err)
	})

	t.Run("Should return an error if monthly income is less than 4 times the installment of the loan", func(t *testing.T) {
		mainIncomeAgeRange := MainIncomeAgeRange{
			MainAge:             18,
			MonthlyIncome:       1000.0,
			LoanValue:           30000.0,
			MonthlyInstallments: 48,
		}
		out := ProponentMainIncomeAgeRangeValidation(mainIncomeAgeRange)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("monthly income is less than 4 times the installment of the loan"), out.Err)
	})

	t.Run("Should return an error if monthly income is less than 3 times the installment of the loan", func(t *testing.T) {
		mainIncomeAgeRange := MainIncomeAgeRange{
			MainAge:             24,
			MonthlyIncome:       1000.0,
			LoanValue:           30000.0,
			MonthlyInstallments: 48,
		}
		out := ProponentMainIncomeAgeRangeValidation(mainIncomeAgeRange)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("monthly income is less than 3 times the installment of the loan"), out.Err)
	})

	t.Run("Should return an error if monthly income is less than 2 times the installment of the loan", func(t *testing.T) {
		mainIncomeAgeRange := MainIncomeAgeRange{
			MainAge:             51,
			MonthlyIncome:       1000.0,
			LoanValue:           30000.0,
			MonthlyInstallments: 48,
		}
		out := ProponentMainIncomeAgeRangeValidation(mainIncomeAgeRange)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("monthly income is less than 2 times the installment of the loan"), out.Err)
	})

	t.Run("Should return error nil if validation succeeds", func(t *testing.T) {
		mainIncomeAgeRange := MainIncomeAgeRange{
			MainAge:             51,
			MonthlyIncome:       2000.0,
			LoanValue:           30000.0,
			MonthlyInstallments: 48,
		}
		out := ProponentMainIncomeAgeRangeValidation(mainIncomeAgeRange)
		var a = assert.New(t)
		a.Nil(out.Err)
	})
}
