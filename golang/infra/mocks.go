package infra

import (
	"time"

	"github.com/adrianosiq/teste-backends/golang/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockString struct {
	mock.Mock
}

func (m MockString) Fields() []string {
	args := m.Called()
	return args.Get(0).([]string)
}

func (m MockString) Split(input string) []string {
	args := m.Called(input)
	return args.Get(0).([]string)
}

func (m MockString) SelectType(schema string, action string, fields []string) ([]string, error) {
	args := m.Called(schema, action, fields)
	return args.Get(0).([]string), args.Error(1)
}

func (m MockString) CreateJSON(fields []string, values []string) ([]byte, error) {
	args := m.Called(fields, values)
	return args.Get(0).([]byte), args.Error(1)
}

var MockMatchedBy = func(x interface{}) bool { return true }

var MockFieldsExp = []string{
	"event_id", "event_schema", "event_action", "event_timestamp", "proposal_id",
	"proposal_loan_value", "proposal_number_of_monthly_installments",
	"warranty_id", "warranty_value", "warranty_province",
	"proponent_id", "proponent_name", "proponent_age", "proponent_monthly_income", "proponent_is_main",
}

var MockSelectedFields = []string{
	"event_id", "event_schema", "event_action", "event_timestamp", "proposal_id",
	"proposal_loan_value", "proposal_number_of_monthly_installments",
}

var MockStringToJSON = []string{
	"c2d06c4f-e1dc-4b2a-af61-ba15bc6d8610,proposal,created,2019-11-11T13:26:04Z,bd6abe95-7c44-41a4-92d0-edf4978c9f4e,684397.0,72",
}

var MockStringSplit = []string{
	"c2d06c4f-e1dc-4b2a-af61-ba15bc6d8610",
	"proposal",
	"created",
	"2019-11-11T13:26:04Z",
	"bd6abe95-7c44-41a4-92d0-edf4978c9f4e",
	"684397.0",
	"72",
}

var parseEventTimestamp, _ = time.Parse(time.RFC3339, "2019-11-11T13:26:04Z")

var MockEventExp = []domain.Event{
	{
		EventID:                             uuid.MustParse("c2d06c4f-e1dc-4b2a-af61-ba15bc6d8610"),
		EventSchema:                         "proposal",
		EventAction:                         "created",
		EventTimestamp:                      parseEventTimestamp,
		ProposalID:                          uuid.MustParse("bd6abe95-7c44-41a4-92d0-edf4978c9f4e"),
		ProposalLoanValue:                   684397.0,
		ProposalNumberOfMonthlyInstallments: 72,
	},
}
