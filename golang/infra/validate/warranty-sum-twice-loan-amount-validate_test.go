package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWarrantySumTwiceLoanAmountValidation(t *testing.T) {
	t.Run("Should return a error if sum of warranties is less than twice the loan amount", func(t *testing.T) {
		sumWarranty := 1200000.0
		loanAmount := 600000.01
		out := WarrantySumTwiceLoanAmountValidation(sumWarranty, loanAmount)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("sum of warranties is less than double the loan amount"), out.Err)
	})

	t.Run("Should return error nil if validation succeeds", func(t *testing.T) {
		sumWarranty := 1200000.0
		loanAmount := 600000.0
		out := WarrantySumTwiceLoanAmountValidation(sumWarranty, loanAmount)
		var a = assert.New(t)
		a.Nil(out.Err)
	})
}
