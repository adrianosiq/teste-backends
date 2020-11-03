package controllers

import (
	"github.com/adrianosiq/teste-backends/golang/domain"
)

//Messages is responsible for the services of the Messages
type Messages struct {
	parse   domain.MessagesParse
	consume domain.MessagesConsume
}

//NewMessages create new Messages
func NewMessages(parse domain.MessagesParse, consume domain.MessagesConsume) Messages {
	return Messages{
		parse:   parse,
		consume: consume,
	}
}

//Handle is reponsible for process inputs and return proposals approved
func (messages Messages) Handle(inputs []string) ([]domain.ProposalApproved, error) {
	events, err := messages.parse.StringToEvent(inputs)
	if err != nil {
		return nil, err
	}
	approved := messages.consume.Process(events)
	return approved, nil
}
