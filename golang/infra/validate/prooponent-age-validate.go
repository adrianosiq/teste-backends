package validate

import (
	"fmt"

	"github.com/adrianosiq/teste-backends/golang/infra"
)

const (
	minProponentAge = 18
)

//ProponentAgeValidation is responsible for validating if preponent age is less the minumum
func ProponentAgeValidation(age int16) infra.Validation {
	if age < minProponentAge {
		return infra.Validation{Err: fmt.Errorf("proponent age is less than minimum of %d", minProponentAge)}
	}
	return infra.Validation{Err: nil}
}
