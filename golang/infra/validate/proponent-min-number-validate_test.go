package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProponentMinNumberValidation(t *testing.T) {
	t.Run("Should return a error if proponents number is less than minimum", func(t *testing.T) {
		number := 1
		out := ProponentMinNumberValidation(number)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("proponents number is less than minimum of %d", minProponentNumber), out.Err)
	})

	t.Run("Should return error nil if validation succeeds", func(t *testing.T) {
		number := 2
		out := ProponentMinNumberValidation(number)
		var a = assert.New(t)
		a.Nil(out.Err)
	})
}
