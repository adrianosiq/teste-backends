package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

//MockConvertString is responsible for mocks the convert string to JSON
type MockConvertString struct {
	mock.Mock
}

func (m MockConvertString) StringToJSON(inputs []string) ([][]byte, error) {
	args := m.Called(inputs)
	return args.Get(0).([][]byte), args.Error(1)
}

type MockConvertEvent struct {
	mock.Mock
}

func (m MockConvertEvent) CreateEvent(data [][]byte) []Event {
	args := m.Called(data)
	return args.Get(0).([]Event)
}

type MockValidateEvent struct {
	mock.Mock
}

func (m MockValidateEvent) Validation(events []Event) []ProposalApproved {
	args := m.Called(events)
	return args.Get(0).([]ProposalApproved)
}

//MockEvent is responsible for mock events return
var MockEvent = []Event{
	{
		EventID:                             uuid.MustParse("c283491e-c546-4468-a264-ceba753c9de0"),
		EventSchema:                         "proposal",
		EventAction:                         "created",
		EventTimestamp:                      time.Now(),
		ProposalID:                          uuid.MustParse("dad6498a-78ad-4a4f-b1e1-34e1f240000c"),
		ProposalLoanValue:                   50000.0,
		ProposalNumberOfMonthlyInstallments: 180,
	},
	{
		EventID:          uuid.MustParse("36cedc9f-f636-47bc-b27e-ef3a86b50616"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("dad6498a-78ad-4a4f-b1e1-34e1f240000c"),
		WarrantyID:       uuid.MustParse("aa795eff-ecbb-4de6-915e-92bb91beebf5"),
		WarrantyValue:    5679086.17,
		WarrantyProvince: "BA",
	},
	{
		EventID:          uuid.MustParse("c3768df0-6a82-4350-a193-20b6296afa47"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("dad6498a-78ad-4a4f-b1e1-34e1f240000c"),
		WarrantyID:       uuid.MustParse("87e1a079-3d20-4255-adfa-aae8b3ba0a06"),
		WarrantyValue:    6139723.75,
		WarrantyProvince: "ES",
	},
	{
		EventID:                uuid.MustParse("f27566eb-4df6-4879-be4f-cc8b8fc5f1ce"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("dad6498a-78ad-4a4f-b1e1-34e1f240000c"),
		ProponentID:            uuid.MustParse("2f384cfb-5b25-44f2-8989-0321ab687cd7"),
		ProponentName:          "Ms. Daryl Pfannerstill",
		ProponentAge:           18,
		ProponentMonthlyIncome: 1120.0,
		ProponentIsMain:        true,
	},
	{
		EventID:                uuid.MustParse("700e70c5-8ff1-4902-a3b9-0c546f7d7a0b"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("dad6498a-78ad-4a4f-b1e1-34e1f240000c"),
		ProponentID:            uuid.MustParse("e1042bdc-28f4-47f5-92fe-dcc55a12e88e"),
		ProponentName:          "Kerry Collier",
		ProponentAge:           19,
		ProponentMonthlyIncome: 61400.0,
		ProponentIsMain:        false,
	},
}

//MockProposalApproved is reponsible for testing return Messages
var MockProposalApproved = []ProposalApproved{
	{
		ProposalID: uuid.MustParse("bd6abe95-7c44-41a4-92d0-edf4978c9f4e"),
	},
	{
		ProposalID: uuid.MustParse("af6e600b-2622-40d1-89ad-d3e5b6cc2fdf"),
	},
}
