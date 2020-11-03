package messages

import (
	"github.com/adrianosiq/teste-backends/golang/domain"
	"github.com/adrianosiq/teste-backends/golang/infra"
)

//Consume is responsible for load services
type Consume struct {
	validate infra.ValidateEvent
}

//NewConsume create a new parse
func NewConsume(validate infra.ValidateEvent) Consume {
	return Consume{
		validate: validate,
	}
}

//Process is responsible for process events and return proposal approved
func (consume Consume) Process(events []domain.Event) []domain.ProposalApproved {
	approvals := consume.validate.Validation(events)
	return approvals
}
