package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWarrantyMinValidation(t *testing.T) {
	t.Run("Should return a error if warranties number is less than minimum", func(t *testing.T) {
		number := 0
		out := WarrantyMinValidation(number)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("warranties number is less than minimum of %d", minWarrantyNumber), out.Err)
	})

	t.Run("Should return error nil if validation succeeds", func(t *testing.T) {
		number := 1
		out := WarrantyMinValidation(number)
		var a = assert.New(t)
		a.Nil(out.Err)
	})
}
