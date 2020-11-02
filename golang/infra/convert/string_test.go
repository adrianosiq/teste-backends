package convert

import (
	"errors"
	"testing"

	"github.com/adrianosiq/teste-backends/golang/infra"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Run("Should StringToJSON throws if SelectType throws", func(t *testing.T) {
		mockString := &infra.MockString{}
		mockString.On("Fields").Return(infra.MockFieldsExp).Once()
		mockString.On("Split", infra.MockStringToJSON[0]).Return(infra.MockStringSplit).Once()
		mockString.On("SelectType", "proposal", "created", infra.MockFieldsExp).Return([]string{}, errors.New("SelectType throws"))
		str := String{deps: mockString}
		out, err := str.StringToJSON(infra.MockStringToJSON)
		var a = assert.New(t)
		a.Error(err)
		a.Equal([][]byte{}, out)
	})

	t.Run("Should StringToJSON throws if CreateJSON throws", func(t *testing.T) {
		mockString := &infra.MockString{}
		mockString.On("Fields").Return(infra.MockFieldsExp).Once()
		mockString.On("Split", infra.MockStringToJSON[0]).Return(infra.MockStringSplit).Once()
		mockString.On("SelectType", "proposal", "created", infra.MockFieldsExp).Return(infra.MockSelectedFields, nil)
		mockString.On("CreateJSON", infra.MockSelectedFields, infra.MockStringSplit).Return([]byte{}, errors.New("CreateJSON throws"))
		str := String{deps: mockString}
		out, err := str.StringToJSON(infra.MockStringToJSON)
		var a = assert.New(t)
		a.Error(err)
		a.Equal([][]byte{}, out)
	})

	t.Run("Should StringToJSON return JSON on success", func(t *testing.T) {
		str := NewString()
		out, err := str.StringToJSON(infra.MockStringToJSON)
		var a = assert.New(t)
		a.Nil(err)
		a.NotNil(out)
	})
}
