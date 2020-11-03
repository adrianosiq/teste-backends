package presentation

import (
	"github.com/adrianosiq/teste-backends/golang/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

//MockMessagesParse is responsible for mocks the messages parse
type MockMessagesParse struct {
	mock.Mock
}

func (m MockMessagesParse) StringToEvent(inputs []string) ([]domain.Event, error) {
	args := m.Called(inputs)
	return args.Get(0).([]domain.Event), args.Error(1)
}

//MockMessagesConsume is responsible for mocks the messages consume
type MockMessagesConsume struct {
	mock.Mock
}

func (m MockMessagesConsume) Process(events []domain.Event) []domain.ProposalApproved {
	args := m.Called(events)
	return args.Get(0).([]domain.ProposalApproved)
}

//MockProposalApproved is reponsible for testing return Messages
var MockProposalApproved = []domain.ProposalApproved{
	{
		ProposalID: uuid.MustParse("bd6abe95-7c44-41a4-92d0-edf4978c9f4e"),
	},
	{
		ProposalID: uuid.MustParse("af6e600b-2622-40d1-89ad-d3e5b6cc2fdf"),
	},
}
