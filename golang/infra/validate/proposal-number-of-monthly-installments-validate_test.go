package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProposalNumberOfMonthlyInstallmentsValidation(t *testing.T) {
	t.Run("Should return a error if number is less than minimum", func(t *testing.T) {
		var number int32 = 23
		out := ProposalNumberOfMonthlyInstallmentsValidation(number)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("proposal number of monthly installments is less than minimum of %d", minNumberMonthlyInstallments), out.Err)
	})

	t.Run("Should return a error if number is greater than maximum", func(t *testing.T) {
		var number int32 = 181
		out := ProposalNumberOfMonthlyInstallmentsValidation(number)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("proposal number of monthly installments is greater than maximum of %d", maxNumberMonthlyInstallments), out.Err)
	})

	t.Run("Should return error nil if validation succeeds", func(t *testing.T) {
		var number int32 = 80
		out := ProposalNumberOfMonthlyInstallmentsValidation(number)
		var a = assert.New(t)
		a.Nil(out.Err)
	})
}
