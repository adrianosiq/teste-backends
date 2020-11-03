package validate

import (
	"fmt"

	"github.com/adrianosiq/teste-backends/golang/infra"
)

const (
	minProponentNumber = 2
)

//ProponentMinNumberValidation is responsible for validating the minimum number of bidders
func ProponentMinNumberValidation(number int) infra.Validation {
	if number < minProponentNumber {
		return infra.Validation{Err: fmt.Errorf("proponents number is less than minimum of %d", minProponentNumber)}
	}
	return infra.Validation{Err: nil}
}
