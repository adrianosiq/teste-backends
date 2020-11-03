package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProponentExactlyNumberMainValidation(t *testing.T) {
	t.Run("Should return a error if the number of the main bidder exactly equal", func(t *testing.T) {
		number := 0
		out := ProponentExactlyNumberMainValidation(number)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("proponents number main is not equal a %d", proponentMainNumber), out.Err)
	})

	t.Run("Should return error nil if validation succeeds", func(t *testing.T) {
		number := 1
		out := ProponentExactlyNumberMainValidation(number)
		var a = assert.New(t)
		a.Nil(out.Err)
	})
}
