package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWarrantyProvinceValidation(t *testing.T) {
	t.Run("Should return an error if warranty province equal a PR", func(t *testing.T) {
		province := "PR"
		out := WarrantyProvinceValidation(province)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("warranty province PR not supported"), out.Err)
	})

	t.Run("Should return an error if warranty province equal a SC", func(t *testing.T) {
		province := "SC"
		out := WarrantyProvinceValidation(province)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("warranty province SC not supported"), out.Err)
	})

	t.Run("Should return an error if warranty province equal a RS", func(t *testing.T) {
		province := "RS"
		out := WarrantyProvinceValidation(province)
		var a = assert.New(t)
		a.NotNil(out.Err)
		a.Equal(fmt.Errorf("warranty province RS not supported"), out.Err)
	})

	t.Run("Should return error nil if validation succeeds", func(t *testing.T) {
		province := "SP"
		out := WarrantyProvinceValidation(province)
		var a = assert.New(t)
		a.Nil(out.Err)
	})
}
