package validate

import (
	"github.com/adrianosiq/teste-backends/golang/domain"
	"github.com/google/uuid"
)

func getEvents(events map[uuid.UUID]domain.Event) []domain.Event {
	var result []domain.Event
	for _, event := range events {
		result = append(result, event)
	}
	return result
}

func checkProposalDelete(events map[uuid.UUID]domain.Event) bool {
	for _, event := range events {
		if event.EventAction == "deleted" {
			return true
		}
	}
	return false
}

func getProposalEvent(events map[uuid.UUID]domain.Event) domain.Event {
	var proposalID uuid.UUID
	proposal := make(map[uuid.UUID]domain.Event)
	for _, event := range events {
		_, ok := proposal[event.ProposalID]
		if !ok || (event.EventAction == "update" && event.EventTimestamp.After(proposal[event.ProposalID].EventTimestamp)) {
			proposal[event.ProposalID] = event
			proposalID = event.ProposalID
		}
	}
	return proposal[proposalID]
}

func getProponentsEvents(events map[uuid.UUID]domain.Event) []domain.Event {
	proponents := make(map[uuid.UUID]domain.Event)
	for _, event := range events {
		_, ok := proponents[event.ProponentID]
		if !ok || (event.EventAction == "update" && event.EventTimestamp.After(proponents[event.ProponentID].EventTimestamp)) {
			proponents[event.ProponentID] = event
		}
	}
	return getEvents(proponents)
}

func getWarrantiesEvents(events map[uuid.UUID]domain.Event) []domain.Event {
	warranties := make(map[uuid.UUID]domain.Event)
	for _, event := range events {
		_, ok := warranties[event.WarrantyID]
		if !ok || (event.EventAction == "update" && event.EventTimestamp.After(warranties[event.WarrantyID].EventTimestamp)) {
			warranties[event.WarrantyID] = event
		}
	}
	return getEvents(warranties)
}

type validateInterface interface {
	Group(events []domain.Event) map[uuid.UUID][]domain.Event
	FilterEventType(events []domain.Event, schema string) map[uuid.UUID]domain.Event
	ProposalValidations(events []domain.Event) (domain.Event, bool)
	ProponentValidations(events []domain.Event, proposal domain.Event) bool
	WarrantyValidations(events []domain.Event, proposal domain.Event) bool
}

type validateDependencies struct{}

func (d validateDependencies) Group(events []domain.Event) map[uuid.UUID][]domain.Event {
	group := make(map[uuid.UUID][]domain.Event)
	for _, event := range events {
		group[event.ProposalID] = append(group[event.ProposalID], event)
	}
	return group
}

func (d validateDependencies) FilterEventType(events []domain.Event, schema string) map[uuid.UUID]domain.Event {
	filter := make(map[uuid.UUID]domain.Event)
	for _, event := range events {
		_, ok := filter[event.EventID]
		if event.EventSchema == schema && !ok {
			filter[event.EventID] = event
		}
	}
	return filter
}

func (d validateDependencies) ProposalValidations(events []domain.Event) (domain.Event, bool) {
	filterEvents := d.FilterEventType(events, "proposal")
	if ok := checkProposalDelete(filterEvents); ok {
		return domain.Event{}, false
	}
	proposal := getProposalEvent(filterEvents)
	if err := ProposalLoanValueValidation(proposal.ProposalLoanValue).Err; err != nil {
		return domain.Event{}, false
	}
	if err := ProposalNumberOfMonthlyInstallmentsValidation(proposal.ProposalNumberOfMonthlyInstallments).Err; err != nil {
		return domain.Event{}, false
	}
	return proposal, true
}

func (d validateDependencies) ProponentValidations(events []domain.Event, proposal domain.Event) bool {
	filterEvents := d.FilterEventType(events, "proponent")
	proponents := getProponentsEvents(filterEvents)
	if err := ProponentMinNumberValidation(len(proponents)).Err; err != nil {
		return false
	}
	countMain := 0
	var proponentMain domain.Event
	for _, proponent := range proponents {
		if err := ProponentAgeValidation(proponent.ProponentAge).Err; err != nil {
			return false
		}
		if proponent.ProponentIsMain {
			countMain++
			proponentMain = proponent
		}
	}
	if err := ProponentExactlyNumberMainValidation(countMain).Err; err != nil {
		return false
	}
	mainIncomeAgeRange := MainIncomeAgeRange{
		MainAge:             proponentMain.ProponentAge,
		MonthlyIncome:       proponentMain.ProponentMonthlyIncome,
		LoanValue:           proposal.ProposalLoanValue,
		MonthlyInstallments: proposal.ProposalNumberOfMonthlyInstallments,
	}
	if err := ProponentMainIncomeAgeRangeValidation(mainIncomeAgeRange).Err; err != nil {
		return false
	}
	return true
}

func (d validateDependencies) WarrantyValidations(events []domain.Event, proposal domain.Event) bool {
	filterEvents := d.FilterEventType(events, "warranty")
	warranties := getWarrantiesEvents(filterEvents)
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
func (validate Validate) Validation(data []domain.Event) []domain.ProposalApproved {
	var approvals []domain.ProposalApproved
	group := validate.deps.Group(data)
	for proposalID, events := range group {
		proposal, isProposalValid := validate.deps.ProposalValidations(events)
		if !isProposalValid {
			continue
		}
		if isProponentValid := validate.deps.ProponentValidations(events, proposal); !isProponentValid {
			continue
		}
		if isWarrantyValid := validate.deps.WarrantyValidations(events, proposal); !isWarrantyValid {
			continue
		}
		approvals = append(approvals, domain.ProposalApproved{ProposalID: proposalID})
	}
	return approvals
}
