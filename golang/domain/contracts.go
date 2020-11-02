package domain

import (
	"time"

	"github.com/google/uuid"
)

//Event is responsible for grouping the event data
type Event struct {
	EventID                             uuid.UUID `json:"event_id,omitempty"`
	EventSchema                         string    `json:"event_schema,omitempty"`
	EventAction                         string    `json:"event_action,omitempty"`
	EventTimestamp                      time.Time `json:"event_timestamp,omitempty"`
	ProposalID                          uuid.UUID `json:"proposal_id,omitempty"`
	ProposalLoanValue                   float64   `json:"proposal_loan_value,string"`
	ProposalNumberOfMonthlyInstallments int32     `json:"proposal_number_of_monthly_installments,string"`
	WarrantyID                          uuid.UUID `json:"warranty_id"`
	WarrantyValue                       float64   `json:"warranty_value,string"`
	WarrantyProvince                    string    `json:"warranty_province"`
	ProponentID                         uuid.UUID `json:"proponent_id"`
	ProponentName                       string    `json:"proponent_name"`
	ProponentAge                        int16     `json:"proponent_age,string"`
	ProponentMonthlyIncome              float64   `json:"proponent_monthly_income,string"`
	ProponentIsMain                     bool      `json:"proponent_is_main"`
}

//ProposalApproved is responsible for grouping data from approved proposal
type ProposalApproved struct {
	ProposalID uuid.UUID `json:"proposal_id"`
}

//MessagesParse is responsible for methods public the parse
type MessagesParse interface {
	StringToEvent(inputs []string) ([]Event, error)
}

//MessagesConsume is responsible for methods public the consume
type MessagesConsume interface {
	Process(events []Event) []ProposalApproved
}
