package messages

import (
	"errors"
	"testing"

	"github.com/adrianosiq/teste-backends/golang/domain"
	"github.com/stretchr/testify/assert"
)

func TestStringToEvent(t *testing.T) {
	t.Run("Should StringToEvent throws if StringToJSON throws", func(t *testing.T) {
		mockConvertStringToJSON := &domain.MockConvertStringToJSON{}
		mockConvertStringToJSON.On("StringToJSON", []string{}).Return([][]byte{}, errors.New("StringToJSON throws")).Once()
		mockConvertEvent := &domain.MockConvertEvent{}
		parse := NewParse(mockConvertStringToJSON, mockConvertEvent)
		out, err := parse.StringToEvent([]string{})
		var a = assert.New(t)
		a.Nil(out)
		a.Error(err)
	})

	t.Run("Should StringToEvent return Event on success", func(t *testing.T) {
		mockConvertStringToJSON := &domain.MockConvertStringToJSON{}
		mockConvertStringToJSON.On("StringToJSON", []string{}).Return([][]byte{}, nil).Once()
		mockConvertEvent := &domain.MockConvertEvent{}
		mockConvertEvent.On("CreateEvent", [][]byte{}).Return(domain.MockEvent).Once()
		parse := NewParse(mockConvertStringToJSON, mockConvertEvent)
		out, err := parse.StringToEvent([]string{})
		var a = assert.New(t)
		a.Nil(err)
		a.Equal(domain.MockEvent, out)
	})
}
