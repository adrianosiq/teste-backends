package validate

import (
	"fmt"

	"github.com/adrianosiq/teste-backends/golang/infra"
)

const (
	proponentMainNumber = 1
)

//ProponentExactlyNumberMainValidation is responsible for validating if preponents number is exactly proponent main number
func ProponentExactlyNumberMainValidation(number int) infra.Validation {
	if number != proponentMainNumber {
		return infra.Validation{Err: fmt.Errorf("proponents number main is not equal a %d", proponentMainNumber)}
	}
	return infra.Validation{Err: nil}
}
