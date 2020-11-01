package infra

import (
	"github.com/adrianosiq/teste-backends/golang/domain"
)

//ConvertStringToJSON is responsible for methods public the string to json
type ConvertStringToJSON interface {
	StringToJSON(inputs []string) ([][]byte, error)
}

//ConvertEvent is responsible for methods public the event
type ConvertEvent interface {
	CreateEvent(data [][]byte) []domain.Event
}

//ValidateEvent is responsible for methods public the validate
type ValidateEvent interface {
	Validation(events []domain.Event) []domain.ProposalApproved
}
