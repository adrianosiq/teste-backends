package validate

import (
	"testing"

	"github.com/adrianosiq/teste-backends/golang/infra"
	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	t.Run("Should Validation return proposal approved on success", func(t *testing.T) {
		validate := NewValidate()
		out := validate.Validation(infra.MockEvents)
		var a = assert.New(t)
		a.Equal(infra.MockApprovedExp, out)
	})
}
