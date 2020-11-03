package validate

import (
	"github.com/adrianosiq/teste-backends/golang/domain"
	"github.com/google/uuid"
)

func removeEvent(checkEvent domain.Event, events []domain.Event) bool {
	for _, event := range events {
		if event.EventSchema == "proposal" || event.EventAction != "removed" {
			continue
		}
		if checkEvent.WarrantyID == event.WarrantyID || event.ProponentID == event.ProponentID {
			return true
		}
	}
	return false
}

func updateEvent(updateEvent domain.Event, events []domain.Event) domain.Event {
	for _, event := range events {
		if event.EventAction != "updated" || event.EventSchema != updateEvent.EventSchema {
			continue
		}
		if event.ProposalID == updateEvent.ProposalID && event.EventTimestamp.After(updateEvent.EventTimestamp) {
			updateEvent = event
		}
	}
	return updateEvent
}

func filterEvents(events []domain.Event, schema string) []domain.Event {
	filterEvents := make(map[uuid.UUID]domain.Event)
	for _, event := range events {
		if event.EventSchema == schema {
			_, ok := filterEvents[event.EventID]
			if !ok && (event.EventAction == "created" || event.EventAction == "added") {
				filterEvents[event.EventID] = event
			}
		}
	}
	selected := []domain.Event{}
	for _, event := range filterEvents {
		if !removeEvent(event, events) {
			event = updateEvent(event, events)
			selected = append(selected, event)
		}
	}
	return selected
}

type validateInterface interface {
	Group(events []domain.Event) map[uuid.UUID][]domain.Event
	ProposalValidations(items []domain.Event) (domain.Event, bool)
	ProponentValidations(items []domain.Event, proposal domain.Event) bool
	WarrantyValidations(items []domain.Event, proposal domain.Event) bool
}

type validateDependencies struct{}

func (d validateDependencies) Group(events []domain.Event) map[uuid.UUID][]domain.Event {
	group := make(map[uuid.UUID][]domain.Event)
	for _, event := range events {
		group[event.ProposalID] = append(group[event.ProposalID], event)
	}
	return group
}

func (d validateDependencies) ProposalValidations(items []domain.Event) (domain.Event, bool) {
	if ok := ProposalDeletedValidation(items); !ok {
		return domain.Event{}, false
	}
	proposal := filterEvents(items, "proposal")[0]
	if err := ProposalLoanValueValidation(proposal.ProposalLoanValue).Err; err != nil {
		return domain.Event{}, false
	}
	if err := ProposalNumberOfMonthlyInstallmentsValidation(proposal.ProposalNumberOfMonthlyInstallments).Err; err != nil {
		return domain.Event{}, false
	}
	return proposal, true
}

func (d validateDependencies) ProponentValidations(items []domain.Event, proposal domain.Event) bool {
	proponents := filterEvents(items, "proponent")
	if err := ProponentMinNumberValidation(len(proponents)).Err; err != nil {
		return false
	}
	countProponentMain := 0
	proponetMain := domain.Event{}
	for _, proponent := range proponents {
		if err := ProponentAgeValidation(proponent.ProponentAge).Err; err != nil {
			return false
		}
		if proponent.ProponentIsMain {
			proponetMain = proponent
			countProponentMain++
		}
	}
	if err := ProponentExactlyNumberMainValidation(countProponentMain).Err; err != nil {
		return false
	}
	mainIncomeAgeRange := MainIncomeAgeRange{
		MainAge:             proponetMain.ProponentAge,
		MonthlyIncome:       proponetMain.ProponentMonthlyIncome,
		LoanValue:           proposal.ProposalLoanValue,
		MonthlyInstallments: proposal.ProposalNumberOfMonthlyInstallments,
	}
	if err := ProponentMainIncomeAgeRangeValidation(mainIncomeAgeRange).Err; err != nil {
		return false
	}
	return true
}

func (d validateDependencies) WarrantyValidations(items []domain.Event, proposal domain.Event) bool {
	warranties := filterEvents(items, "warranty")
	if err := WarrantyMinValidation(len(warranties)).Err; err != nil {
		return false
	}
	var warrantySumVlues float64
	for _, warranty := range warranties {
		if err := WarrantyProvinceValidation(warranty.WarrantyProvince).Err; err != nil {
			return false
		}
		warrantySumVlues += warranty.WarrantyValue
	}
	if err := WarrantySumTwiceLoanAmountValidation(warrantySumVlues, proposal.ProposalLoanValue).Err; err != nil {
		return false
	}
	return true
}

//Validate load is responsible for load depedencies
type Validate struct {
	deps validateInterface
}

//NewValidate create new validate
func NewValidate() Validate {
	return Validate{
		deps: validateDependencies{},
	}
}

//Validation is responsible for return proposal approved
func (validate Validate) Validation(events []domain.Event) []domain.ProposalApproved {
	var approvals []domain.ProposalApproved
	group := validate.deps.Group(events)
	for proposalID, items := range group {
		proposal, isProposalValid := validate.deps.ProposalValidations(items)
		if !isProposalValid {
			continue
		}
		if isProponentValid := validate.deps.ProponentValidations(items, proposal); !isProponentValid {
			continue
		}
		if warrantyIsValid := validate.deps.WarrantyValidations(items, proposal); !warrantyIsValid {
			continue
		}
		approvals = append(approvals, domain.ProposalApproved{ProposalID: proposalID})
	}
	return approvals
}
