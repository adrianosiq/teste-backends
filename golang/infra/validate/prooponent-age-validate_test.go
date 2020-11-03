package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProponentAgeValidation(t *testing.T) {
	t.Run("Should return a error if the age is less than minimum", func(t *testing.T) {
		var age int16 = 17
		out := ProponentAgeValidation(age)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("proponent age is less than minimum of %d", minProponentAge), out.Err)
	})

	t.Run("Should return error nil if validation succeeds", func(t *testing.T) {
		var age int16 = 18
		out := ProponentAgeValidation(age)
		var a = assert.New(t)
		a.Nil(out.Err)
	})
}
