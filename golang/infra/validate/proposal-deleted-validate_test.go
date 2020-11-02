package validate

import (
	"testing"
	"time"

	"github.com/adrianosiq/teste-backends/golang/domain"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func TestProposalDeletedValidation(t *testing.T) {
	t.Run("Should return a false is proposal is deleted", func(t *testing.T) {
		proposals := []domain.Event{
			{
				EventID:                             uuid.MustParse("e058782a-d520-470b-ab71-7aaf138d57ce"),
				EventSchema:                         "proposal",
				EventAction:                         "created",
				EventTimestamp:                      time.Now(),
				ProposalID:                          uuid.MustParse("924166e0-1d9b-4cf4-8d68-f59a80060844"),
				ProposalLoanValue:                   50000.0,
				ProposalNumberOfMonthlyInstallments: 180,
			},
			{
				EventID:        uuid.MustParse("e058782a-d520-470b-ab71-7aaf138d57ce"),
				EventSchema:    "proposal",
				EventAction:    "deleted",
				EventTimestamp: time.Now(),
				ProposalID:     uuid.MustParse("924166e0-1d9b-4cf4-8d68-f59a80060844"),
			},
		}
		out := ProposalDeletedValidation(proposals)
		var a = assert.New(t)
		a.False(out)
	})

	t.Run("Should return a true is proposal is not deleted", func(t *testing.T) {
		proposals := []domain.Event{
			{
				EventID:                             uuid.MustParse("e058782a-d520-470b-ab71-7aaf138d57ce"),
				EventSchema:                         "proposal",
				EventAction:                         "created",
				EventTimestamp:                      time.Now(),
				ProposalID:                          uuid.MustParse("924166e0-1d9b-4cf4-8d68-f59a80060844"),
				ProposalLoanValue:                   50000.0,
				ProposalNumberOfMonthlyInstallments: 180,
			},
		}
		out := ProposalDeletedValidation(proposals)
		var a = assert.New(t)
		a.True(out)
	})
}
