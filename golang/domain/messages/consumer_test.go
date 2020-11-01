package messages

import (
	"testing"

	"github.com/adrianosiq/teste-backends/golang/domain"
	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	t.Run("Should Process return ProposalApproved on success", func(t *testing.T) {
		mockValidateEvent := &domain.MockValidateEvent{}
		mockValidateEvent.On("Validation", domain.MockEvent).Return(domain.MockProposalApproved).Once()
		consume := NewConsume(mockValidateEvent)
		out := consume.Process(domain.MockEvent)
		var a = assert.New(t)
		a.Equal(domain.MockProposalApproved, out)
	})
}
