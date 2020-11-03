package factories

import (
	"github.com/adrianosiq/teste-backends/golang/cmd/presentation/controllers"
	"github.com/adrianosiq/teste-backends/golang/domain"
	"github.com/adrianosiq/teste-backends/golang/domain/messages"
	"github.com/adrianosiq/teste-backends/golang/infra/convert"
	"github.com/adrianosiq/teste-backends/golang/infra/validate"
)

//MessagesFactory is responsible for factory the messages
func MessagesFactory(inputs []string) ([]domain.ProposalApproved, error) {
	newString := convert.NewString()
	newEvent := convert.NewEvent()
	newValidate := validate.NewValidate()
	newParse := messages.NewParse(newString, newEvent)
	newConsume := messages.NewConsume(newValidate)
	newMessages := controllers.NewMessages(newParse, newConsume)
	return newMessages.Handle(inputs)
}
