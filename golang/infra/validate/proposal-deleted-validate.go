package validate

import "github.com/adrianosiq/teste-backends/golang/domain"

//ProposalDeletedValidation is responsible for validation if proposal than deleted
func ProposalDeletedValidation(events []domain.Event) bool {
	for _, event := range events {
		if event.EventSchema == "proposal" && event.EventAction == "deleted" {
			return false
		}
	}
	return true
}
