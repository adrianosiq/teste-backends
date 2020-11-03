package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProposalLoanValueValidation(t *testing.T) {
	t.Run("Should return a error if value is less than minimum", func(t *testing.T) {
		value := 29999.99
		out := ProposalLoanValueValidation(value)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("proposal loan value is less than minimum of %g", minLoanValue), out.Err)
	})

	t.Run("Should return a error if value is greater than maximum", func(t *testing.T) {
		value := 3000000.01
		out := ProposalLoanValueValidation(value)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("proposal loan value is greater than maximum of %g", maxLoanValue), out.Err)
	})

	t.Run("Should return error nil if validation succeeds", func(t *testing.T) {
		value := 600000.0
		out := ProposalLoanValueValidation(value)
		var a = assert.New(t)
		a.Nil(out.Err)
	})
}
