package controllers

import (
	"errors"
	"testing"

	"github.com/adrianosiq/teste-backends/golang/domain"

	"github.com/adrianosiq/teste-backends/golang/cmd/presentation"
	"github.com/stretchr/testify/assert"
)

func TestMessages(t *testing.T) {
	t.Run("Should Handle throws if StringToEvent throws", func(t *testing.T) {
		mockMessagesParse := &presentation.MockMessagesParse{}
		mockMessagesParse.On("StringToEvent", []string{}).Return([]domain.Event{}, errors.New("StringToEvent throws")).Once()
		mockMessagesConsume := &presentation.MockMessagesConsume{}
		messages := NewMessages(mockMessagesParse, mockMessagesConsume)
		out, err := messages.Handle([]string{})
		var a = assert.New(t)
		a.Nil(out)
		a.Error(err)
	})

	t.Run("Should Handle return ProposalApproved on success", func(t *testing.T) {
		mockMessagesParse := &presentation.MockMessagesParse{}
		mockMessagesParse.On("StringToEvent", []string{}).Return([]domain.Event{}, nil).Once()
		mockMessagesConsume := &presentation.MockMessagesConsume{}
		mockMessagesConsume.On("Process", []domain.Event{}).Return(presentation.MockProposalApproved).Once()
		messages := NewMessages(mockMessagesParse, mockMessagesConsume)
		out, err := messages.Handle([]string{})
		var a = assert.New(t)
		a.Nil(err)
		a.Equal(presentation.MockProposalApproved, out)
	})
}
